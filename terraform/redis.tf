resource "azurerm_redis_cache" "redis" {
  name                = "${var.name_prefix}-redis"
  resource_group_name = azurerm_resource_group.rg.name
  location            = azurerm_resource_group.rg.location

  capacity = 0
  family   = "C"
  sku_name = "Basic"

  non_ssl_port_enabled = true

  public_network_access_enabled = true
}