// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

module "resource_names" {
  source  = "terraform.registry.launch.nttdata.com/module_library/resource_name/launch"
  version = "~> 1.0"

  for_each = var.resource_names_map

  logical_product_family  = var.logical_product_family
  logical_product_service = var.logical_product_service
  region                  = var.location
  class_env               = var.class_env
  cloud_resource_type     = each.value.name
  instance_env            = var.instance_env
  instance_resource       = var.instance_resource
  maximum_length          = each.value.max_length
}

module "resource_group" {
  source  = "terraform.registry.launch.nttdata.com/module_primitive/resource_group/azurerm"
  version = "~> 1.0"

  name     = module.resource_names["resource_group"].minimal_random_suffix
  location = var.location

  tags = merge(var.tags, { resource_name = module.resource_names["resource_group"].standard })
}

module "log_analytics_workspace" {
  source  = "terraform.registry.launch.nttdata.com/module_primitive/log_analytics_workspace/azurerm"
  version = "~> 1.0"

  name                = module.resource_names["log_analytics_workspace"].minimal_random_suffix
  resource_group_name = module.resource_group.name
  location            = var.location

  sku = "PerGB2018"

  tags = merge(var.tags, { resource_name = module.resource_names["log_analytics_workspace"].standard })

  depends_on = [module.resource_group]
}

module "monitor_private_link_scope" {
  source  = "terraform.registry.launch.nttdata.com/module_primitive/azure_monitor_private_link_scope/azurerm"
  version = "~> 1.0"

  name                = module.resource_names["monitor_private_link_scope"].minimal_random_suffix
  resource_group_name = module.resource_group.name

  linked_resource_ids = {}

  tags = merge(var.tags, { resource_name = module.resource_names["monitor_private_link_scope"].standard })

  depends_on = [module.resource_group]
}

module "monitor_private_link_scoped_service" {
  source = "../.."

  name = module.log_analytics_workspace.name

  monitor_private_link_scope_id = module.monitor_private_link_scope.private_link_scope_id
  resource_id                   = module.log_analytics_workspace.id

  depends_on = [module.resource_group, module.log_analytics_workspace, module.monitor_private_link_scope]
}
