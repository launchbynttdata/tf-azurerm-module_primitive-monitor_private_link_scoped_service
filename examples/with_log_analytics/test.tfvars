resource_names_map = {
  resource_group = {
    name       = "rg"
    max_length = 80
  }
  monitor_private_link_scope = {
    name       = "ampls"
    max_length = 80
  }
  log_analytics_workspace = {
    name       = "law"
    max_length = 80
  }
}
instance_env            = 0
instance_resource       = 0
logical_product_family  = "launch"
logical_product_service = "ampls"
class_env               = "gotest"
location                = "eastus"
