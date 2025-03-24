variable "region" {
  type = string
  nullable = false
}

variable "subscription_id" {
  type = string
  nullable = false
}

terraform {
  required_providers {
    azurerm = {
      source = "hashicorp/azurerm"
      version = "4.24.0"
    }
  }
}

provider "azurerm" {
  subscription_id = var.subscription_id
  features {}
}

resource "azurerm_resource_group" "dhbw_devops" {
  name     = "dhbw-devops"
  location = var.region
}

resource "azurerm_kubernetes_cluster" "dhbw_devops" {
  name                = "dhbw-devops-aks"
  location            = azurerm_resource_group.dhbw_devops.location
  resource_group_name = azurerm_resource_group.dhbw_devops.name
  dns_prefix          = "dhbwdevops"

  default_node_pool {
    name       = "default"
    node_count = 1
    vm_size    = "Standard_D2_v2"
  }

  identity {
    type = "SystemAssigned"
  }

  tags = {
    Environment = "Production"
  }
}

output "client_certificate" {
  value     = azurerm_kubernetes_cluster.dhbw_devops.kube_config[0].client_certificate
  sensitive = true
}

output "kube_config" {
  value = azurerm_kubernetes_cluster.dhbw_devops.kube_config_raw
  sensitive = true
}
