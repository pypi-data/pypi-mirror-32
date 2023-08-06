#!/usr/bin/env python

# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.
# Copyright (c) 2017 Mozilla Corporation
# Contributors: Guillaume Destuynder <kang@mozilla.com>

import http.client
import json
import logging
import time


class DotDict(dict):
    """return a dict.item notation for dict()'s"""
    __getattr__ = dict.__getitem__
    __setattr__ = dict.__setitem__
    __delattr__ = dict.__delitem__

    def __init__(self, dct):
        for key, value in dct.items():
            if hasattr(value, 'keys'):
                value = DotDict(value)
            self[key] = value


class AuthZeroRule(object):
    """Lightweight Rule Object"""
    def __init__(self):
        self.id = None
        self.enabled = False
        self.script = None
        self.name = None
        self.order = 0
        self.stage = 'login_success'

    def validate(self):
        if self.script is None:
            raise Exception('RuleValidationError', ('script cannot be None'))
        if self.name is None:
            raise Exception('RuleValidationError', ('name cannot be None'))
        if self.order <= 0:
            raise Exception('RuleValidationError', ('order must be greater than 0'))
        return True

    def json(self):
        tmp = {'id': self.id, 'enabled': self.enabled, 'script': self.script,
               'name': self.name, 'order': self.order, 'stage': self.stage}
        # Remove id if we don't have one. It means it's a new rule.
        if tmp.get('id') is None:
            tmp.pop('id')
        return json.dumps(tmp)

    def __str__(self):
        return self.json().__str__()


class AuthZero(object):
    def __init__(self, config):
        config = DotDict(config)
        self.default_headers = {
            'content-type': "application/json"
        }
        self.uri = config.uri
        self.client_id = config.client_id
        self.client_secret = config.client_secret
        self.access_token = None
        self.access_token_scope = None
        self.access_token_valid_until = 0
        self.access_token_auto_renew = False
        self.access_token_auto_renew_leeway = 60
        self.conn = http.client.HTTPSConnection(config.uri)
        self.rules = []

        self.logger = logging.getLogger('AuthZero')

    def __del__(self):
        self.client_secret = None
        self.conn.close()

    def get_rules(self):
        return self._request("/api/v2/rules")

    def delete_rule(self, rule_id):
        """
        rule_id: string
        rule: AuthZeroRule object

        Deletes an Auth0 rule
        Auth0 API doc: https://auth0.com/docs/api/management/v2#!/rules
        Auth0 API endpoint: PATH /api/v2/rules/{id}
        Auth0 API parameters: id (rule_id, required)
        """
        return self._request("/api/v2/rules/{}".format(rule_id), "DELETE")

    def create_rule(self, rule):
        """
        rule_id: string
        rule: AuthZeroRule object

        Creates an Auth0 rule
        Auth0 API doc: https://auth0.com/docs/api/management/v2#!/rules
        Auth0 API endpoint: PATH /api/v2/rules
        Auth0 API parameters: body (required)
        """
        payload = DotDict(dict())
        payload.script = rule.script
        payload.name = rule.name
        payload.order = rule.order
        payload.enabled = rule.enabled
        payload_json = json.dumps(payload)
        return self._request("/api/v2/rules", "POST", payload_json)

    def update_rule(self, rule_id, rule):
        """
        rule_id: string
        rule: AuthZeroRule object

        Updates an Auth0 rule
        Auth0 API doc: https://auth0.com/docs/api/management/v2#!/rules
        Auth0 API endpoint: PATH /api/v2/rules/{id}
        Auth0 API parameters: id (rule_id, required), body (required)
        """

        payload = DotDict(dict())
        payload.script = rule.script
        payload.name = rule.name
        payload.order = rule.order
        payload.enabled = rule.enabled
        payload_json = json.dumps(payload)
        return self._request("/api/v2/rules/{}".format(rule_id), "PATCH", payload_json)

    def get_clients(self):
        """
        Get list of Auth0 clients
        Auth0 API doc: https://auth0.com/docs/api/management/v2#!/clients
        Auth0 API endpoint: PATH /api/v2/clients
        Auth0 API parameters: fields (optional), ...
        """

        payload = DotDict(dict())
        payload_json = json.dumps(payload)
        page = 0
        per_page = 100
        totals = 0
        done = -1
        clients = []
        while totals > done:
            ret = self._request("/api/v2/clients?"
                                "&per_page={per_page}"
                                "&page={page}&include_totals=true"
                                "".format(page=page, per_page=per_page),
                                "GET",
                                payload_json)
            clients += ret['clients']
            done = done + per_page
            page = page + 1
            totals = ret['total']
            logging.debug("Got {} clients out of {} - current page {}".format(done, totals, page))
        return clients

    def create_client(self, client):
        """
        client: client object (dict)

        Create an Auth0 client
        Auth0 API doc: https://auth0.com/docs/api/management/v2#!/clients
        Auth0 API endpoint: PATH /api/v2/clients
        Auth0 API parameters: body (required)
        """
        payload_json = json.dumps(client)
        return self._request("/api/v2/clients", "POST", payload_json)

    def update_client(self, client_id, client):
        """
        client_id: client id (string)
        client: client object (dict)

        Update an Auth0 client
        Auth0 API doc: https://auth0.com/docs/api/management/v2#!/clients
        Auth0 API endpoint: PATH /api/v2/clients
        Auth0 API parameters: id (required), body (required)
        """
        # We accept clients ready by Auth0's API (GET) but updates (PATCH) use a DIFFERENT format.
        # This means we need to CONVERT GET'd clients so that they will be accepted for PATCH'ing.
        # That's quite annoying, by the way. Why?! :-(
        # See also https://auth0.com/docs/api/management/v2#!/Clients/patch_clients_by_id

        patch_api_schema = """
        {
          "name": "",
          "description": "",
          "client_secret": "",
          "logo_uri": "",
          "callbacks": [
            ""
          ],
          "allowed_origins": [
            ""
          ],
          "web_origins": [
            ""
          ],
          "grant_types": [
            ""
          ],
          "client_aliases": [
            ""
          ],
          "allowed_clients": [
            ""
          ],
          "allowed_logout_urls": [
            ""
          ],
          "jwt_configuration": {
            "lifetime_in_seconds": 0,
            "scopes": {},
            "alg": ""
          },
          "encryption_key": {
            "pub": "",
            "cert": "",
            "subject": ""
          },
          "sso": false,
          "cross_origin_auth": false,
          "cross_origin_loc": "",
          "sso_disabled": false,
          "custom_login_page_on": false,
          "token_endpoint_auth_method": "",
          "app_type": "",
          "oidc_conformant": false,
          "custom_login_page": "",
          "custom_login_page_preview": "",
          "form_template": "",
          "addons": {
            "aws": {},
            "azure_blob": {},
            "azure_sb": {},
            "rms": {},
            "mscrm": {},
            "slack": {},
            "sentry": {},
            "box": {},
            "cloudbees": {},
            "concur": {},
            "dropbox": {},
            "echosign": {},
            "egnyte": {},
            "firebase": {},
            "newrelic": {},
            "office365": {},
            "salesforce": {},
            "salesforce_api": {},
            "salesforce_sandbox_api": {},
            "samlp": {},
            "layer": {},
            "sap_api": {},
            "sharepoint": {},
            "springcm": {},
            "wams": {},
            "wsfed": {},
            "zendesk": {},
            "zoom": {}
          },
          "client_metadata": {},
          "mobile": {
            "android": {
              "app_package_name": "",
              "sha256_cert_fingerprints": [
              ]
            },
            "ios": {
              "team_id": "",
              "app_bundle_identifier": ""
            }
          }
        }
        """

        patch_api_dict = json.loads(patch_api_schema)
        payload = {}
        for i in patch_api_dict:
            val = client.get(i)
            if (val is not None):
                # Sub structure! we just check one level deep
                if isinstance(val, dict) and len(patch_api_dict[i]) > 0:
                    for y in patch_api_dict[i]:
                        payload[i] = {}
                        try:
                            payload[i][y] = val[y]
                        except KeyError:
                            pass #don't have this, then ignore it
                else:
                    payload[i] = val

        payload_json = json.dumps(payload)
        return self._request("/api/v2/clients/{}".format(client_id),
                             "PATCH",
                             payload_json)

    def delete_client(self, client_id):
        """
        client_id: client id (string)

        Delete an Auth0 client
        Auth0 API doc: https://auth0.com/docs/api/management/v2#!/clients
        Auth0 API endpoint: PATH /api/v2/clients
        Auth0 API parameters: id (required), body (required)
        """
        return self._request("/api/v2/clients/{}".format(client_id), "DELETE")

    def get_users(self, fields="username,user_id,name,email,identities,groups", query_filter=""):
        """
        Returns a list of users from the Auth0 API.
        query_filter: string
        returns: JSON dict of the list of users
        """
        payload = DotDict(dict())
        payload_json = json.dumps(payload)
        page = 0
        per_page = 100
        totals = 0
        done = -1
        users = []
        while totals > done:
            ret = self._request("/api/v2/users?fields={fields}&"
                                "search_engine=v2&q={query_filter}&per_page={per_page}"
                                "&page={page}&include_totals=true"
                                "".format(fields=fields, query_filter=query_filter, page=page, per_page=per_page),
                                "GET",
                                payload_json)
            users += ret['users']
            done = done + per_page
            page = page + 1
            totals = ret['total']
            logging.debug("Got {} users out of {} - current page {}".format(done, totals, page))
        return users

    def get_logs(self):
        return self._request("/api/v2/logs")

    def get_user(self, user_id):
        """Return user from the Auth0 API.
        user_id: string
        returns: JSON dict of the user profile
        """

        payload = DotDict(dict())
        payload_json = json.dumps(payload)
        return self._request("/api/v2/users/{}".format(user_id),
                             "GET",
                             payload_json)

    def update_user(self, user_id, new_profile):
        """
        user_id: string
        new_profile: dict (can be a JSON string loaded with json.loads(str) for example)

        Update a user in auth0 and return it as a dict to the caller.
        Auth0 API doc: https://auth0.com/docs/api/management/v2
        Auth0 API endpoint: PATCH /api/v2/users/{id}
        Auth0 API parameters: id (user_id, required), body (required)
        """

        payload = DotDict(dict())
        assert type(new_profile) is dict
        # Auth0 does not allow passing the user_id attribute
        # as part of the payload (it's in the PATCH query already)
        if 'user_id' in new_profile.keys():
            del new_profile['user_id']
        payload.app_metadata = new_profile
        # This validates the JSON as well
        payload_json = json.dumps(payload)

        return self._request("/api/v2/users/{}".format(user_id),
                             "PATCH",
                             payload_json)

    def get_access_token(self):
        """
        Returns a JSON object containing an OAuth access_token.
        This is also stored in this class other functions to use.
        """
        payload = DotDict(dict())
        payload.client_id = self.client_id
        payload.client_secret = self.client_secret
        payload.audience = "https://{}/api/v2/".format(self.uri)
        payload.grant_type = "client_credentials"
        payload_json = json.dumps(payload)

        ret = self._request("/oauth/token", "POST", payload_json, authorize=False)

        access_token = DotDict(ret)
        # Validation
        if ('access_token' not in access_token.keys()):
            raise Exception('InvalidAccessToken', access_token)
        self.access_token = access_token.access_token
        self.access_token_valid_until = time.time() + access_token.expires_in
        self.access_token_scope = access_token.scope
        return access_token

    def _request(self, rpath, rtype="GET", payload_json={}, authorize=True):
        self.logger.debug('Sending Auth0 request {} {}'.format(rtype, rpath))
        if authorize:
            self.conn.request(rtype, rpath, payload_json, self._authorize(self.default_headers))
        else:
            # Public req w/o oauth header
            self.conn.request(rtype, rpath, payload_json, self.default_headers)
        return self._handle_response()

    def _handle_response(self):
        res = self.conn.getresponse()
        self._check_http_response(res)
        ret = json.loads(res.read().decode('utf-8'))
        return ret

    def _authorize(self, headers):
        if not self.access_token:
            raise Exception('InvalidAccessToken')
        if self.access_token_auto_renew:
            if self.access_token_valid_until < time.time() + self.access_token_auto_renew_leeway:
                self.access_token = self.get_access_token()
        elif self.access_token_valid_until < time.time():
            raise Exception('InvalidAccessToken', 'The access token has expired')

        local_headers = {}
        local_headers.update(headers)
        local_headers['Authorization'] = 'Bearer {}'.format(self.access_token)

        return local_headers

    def _check_http_response(self, response):
        """Check that we got a 2XX response from the server, else bail out"""
        if (response.status >= 300) or (response.status < 200):
            self.logger.debug("_check_http_response() HTTP communication failed: {} {}"
                              .format(response.status, response.reason, response.read().decode('utf-8')))
            raise Exception('HTTPCommunicationFailed', (response.status, response.reason))
