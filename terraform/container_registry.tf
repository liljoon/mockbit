resource "azurerm_container_registry" "acr" {
  name                = "${var.name_prefix}acr"
  location            = var.location
  resource_group_name = azurerm_resource_group.rg.name
  sku                 = "Basic"
}