terraform {
  required_providers {
    yoshik3s = {
      source = "registry.terraform.io/HideyoshiNakazone/yoshik3s"
    }
  }
}


provider "yoshik3s" {}


resource "yoshik3s_cluster" "example_cluster" {
  name        = "example-cluster"
  token       = "example_token"
  k3s_version = "v1.30.2+k3s2"
}


resource "yoshik3s_master_node" "example_master_node" {
  cluster = yoshik3s_cluster.example_cluster

  node_connection = {
    host     = "localhost"
    port     = "2222"
    user     = "sshuser"
    password = "password"
  }

  node_options = [
    "--disable traefik",
    "--node-label node_type=master",
    "--snapshotter native",
  ]
}


locals {
  example_cluster_workers = {
    worker1 = {
      host = "localhost"
      port = "3333"
    }
  }
}


resource "yoshik3s_worker_node" "example_worker_node" {
  master_server_address = "master_node"

  cluster = yoshik3s_cluster.example_cluster

  for_each = local.example_cluster_workers

  node_connection = {
    host     = each.value.host
    port     = each.value.port
    user     = "sshuser"
    password = "password"
  }

  node_options = [
    "--node-label node_type=worker",
    "--snapshotter native",
  ]
}