# Copyright 2018 Capital One Services, LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import importlib
import os
import logging
from azure.cli.core.cloud import AZURE_PUBLIC_CLOUD
from azure.cli.core._profile import Profile
from azure.common.credentials import ServicePrincipalCredentials, BasicTokenAuthentication
from c7n_azure.utils import ResourceIdParser


class Session(object):

    def __init__(self, subscription_id=None):
        """
        Creates a session using available authentication type.

        Auth priority:
        1. Token Auth
        2. Tenant Auth
        3. Azure CLI Auth

        :param subscription_id: If provided, overrides environment variables.
        """

        self.log = logging.getLogger('custodian.azure.session')
        self._provider_cache = {}

        tenant_auth_variables = [
            'AZURE_TENANT_ID', 'AZURE_SUBSCRIPTION_ID',
            'AZURE_CLIENT_ID', 'AZURE_CLIENT_SECRET'
        ]
        token_auth_variables = ['AZURE_ACCESS_TOKEN', 'AZURE_SUBSCRIPTION_ID']

        if all(k in os.environ for k in token_auth_variables):
            # Token authentication
            self.credentials = BasicTokenAuthentication(
                token={
                    'access_token': os.environ['AZURE_ACCESS_TOKEN']
                })
            self.subscription_id = os.environ['AZURE_SUBSCRIPTION_ID']
            self.log.info("Creating session with Token Authentication")

        elif all(k in os.environ for k in tenant_auth_variables):
            # Tenant (service principal) authentication
            self.credentials = ServicePrincipalCredentials(
                client_id=os.environ['AZURE_CLIENT_ID'],
                secret=os.environ['AZURE_CLIENT_SECRET'],
                tenant=os.environ['AZURE_TENANT_ID']
            )
            self.subscription_id = os.environ['AZURE_SUBSCRIPTION_ID']
            self.tenant_id = os.environ['AZURE_TENANT_ID']
            self.log.info("Creating session with Service Principal Authentication")

        else:
            # Azure CLI authentication
            (self.credentials,
             self.subscription_id,
             self.tenant_id) = Profile().get_login_credentials(
                resource=AZURE_PUBLIC_CLOUD.endpoints.active_directory_resource_id)
            self.log.info("Creating session with Azure CLI Authentication")

        # Let provided id parameter override everything else
        if subscription_id is not None:
            self.subscription_id = subscription_id

        self.log.info("Session using Subscription ID: %s" % self.subscription_id)

        if self.credentials is None:
            self.log.error('Unable to locate credentials for Azure session.')

    def client(self, client):
        service_name, client_name = client.rsplit('.', 1)
        svc_module = importlib.import_module(service_name)
        klass = getattr(svc_module, client_name)
        return klass(self.credentials, self.subscription_id)

    def resource_api_version(self, resource_id):
        """ latest non-preview api version for resource """

        namespace = ResourceIdParser.get_namespace(resource_id)
        resource_type = ResourceIdParser.get_resource_type(resource_id)

        if resource_type in self._provider_cache:
            return self._provider_cache[resource_type]

        resource_client = self.client('azure.mgmt.resource.ResourceManagementClient')
        provider = resource_client.providers.get(namespace)

        rt = next((t for t in provider.resource_types
            if t.resource_type == str(resource_type).split('/')[-1]), None)
        if rt and rt.api_versions:
            versions = [v for v in rt.api_versions if 'preview' not in v.lower()]
            api_version = versions[0] if versions else rt.api_versions[0]
            self._provider_cache[resource_type] = api_version
            return api_version
