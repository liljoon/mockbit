resource "azurerm_postgresql_flexible_server" "psql" {
  name                = "${var.name_prefix}-psql"
  resource_group_name = azurerm_resource_group.rg.name
  location            = var.location

  sku_name = var.db_sku_name
  version  = "16"

  administrator_login    = var.db_administrator_login
  administrator_password = var.db_administrator_password

  depends_on = [azurerm_resource_group.rg]
  zone       = "1"
}

resource "azurerm_postgresql_flexible_server_firewall_rule" "psql_firewall" {
  server_id        = azurerm_postgresql_flexible_server.psql.id
  name             = "AllowAll"
  start_ip_address = "0.0.0.0"
  end_ip_address   = "255.255.255.255"
}