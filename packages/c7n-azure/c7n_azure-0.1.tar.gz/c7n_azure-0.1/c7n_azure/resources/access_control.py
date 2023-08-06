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

from c7n_azure.provider import resources
from c7n_azure.query import QueryResourceManager
from c7n_azure.session import Session

from c7n.filters import ValueFilter
from c7n.filters.related import RelatedResourceFilter
from c7n.utils import type_schema


@resources.register('roleassignment')
class RoleAssignment(QueryResourceManager):

    class resource_type(object):
        service = 'azure.mgmt.authorization'
        client = 'AuthorizationManagementClient'
        enum_spec = ('role_assignments', 'list', None)
        get_spec = ('role_assignments', 'get_by_id', None)
        id = 'id'
        default_report_fields = (
            'name',
            'type',
            'properties.scope',
            'properties.roleDefinitionId'
        )


@resources.register('roledefinition')
class RoleDefinition(QueryResourceManager):

    class resource_type(object):
        s = Session()
        service = 'azure.mgmt.authorization'
        client = 'AuthorizationManagementClient'
        enum_spec = ('role_definitions', 'list', '/subscriptions/%s' % (s.subscription_id))
        get_spec = ('role_definitions', 'get_by_id', None)
        type = 'roleDefinition'
        id = 'id'
        default_report_fields = (
            'id',
            'name',
            'type',
            'properties.roleName',
            'properties.description',
            'properties.type',
            'properties.permissions'
        )


@RoleAssignment.filter_registry.register('role')
class UserRole(RelatedResourceFilter):
    """Filters role assignments based on role definitions

    :Example:

        .. code-block:: yaml

            policies:
               - name: assignments-by-role-definition
                 resource: azure.roleassignment
                 filters:
                    - type: role
                      key: properties.roleName
                      op: in
                      value: Owner
    """

    schema = type_schema('role', rinherit=ValueFilter.schema)

    RelatedResource = "c7n_azure.resources.access_control.RoleDefinition"
    RelatedIdsExpression = "properties.roleDefinitionId"
