locals {
  ampls_resource_map = regex(
    "(?i)^/Subscriptions/(?P<subscription_id>[^/]+)/resourceGroups/(?P<resource_group>[^/]+)/providers/(?P<provider>[^/]+)/.+/(?P<resource_name>.+)$",
    var.monitor_private_link_scope_id
  )
}
