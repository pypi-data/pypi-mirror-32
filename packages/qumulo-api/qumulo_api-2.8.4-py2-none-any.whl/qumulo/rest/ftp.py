# Copyright (c) 2017 Qumulo, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License"); you may not
# use this file except in compliance with the License. You may obtain a copy of
# the License at http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
# WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
# License for the specific language governing permissions and limitations under
# the License.

import qumulo.lib.request as request

@request.request
def get_status(conninfo, credentials):
    return request.rest_request(conninfo, credentials, "GET", "/v0/ftp/status")

@request.request
def get_settings(conninfo, credentials):
    return request.rest_request(
        conninfo, credentials, "GET", "/v0/ftp/settings")

@request.request
def modify_settings(
    conninfo,
    credentials,
    enabled=None,
    check_remote_host=None,
    log_operations=None):

    request_body = {}
    if enabled is not None:
        request_body['enabled'] = enabled
    if check_remote_host is not None:
        request_body['check_remote_host'] = check_remote_host
    if log_operations is not None:
        request_body['log_operations'] = log_operations

    return request.rest_request(
        conninfo, credentials, "PATCH", "/v0/ftp/settings", body=request_body)

