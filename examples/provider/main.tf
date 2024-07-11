terraform {
    required_providers {
        yoshik3s = {
            source = "registry.terraform.io/HideyoshiNakazone/yoshi-k3s"
        }
    }
}


provider "yoshik3s" {}


resource "yoshik3s_cluster" "example_cluster" {
    name = "example-cluster"
    token = "example_token"
    k3s_version = "v1.30.2+k3s2"
}


resource "yoshik3s_master_node" "example_master_node" {
    cluster = {
        token = yoshik3s_cluster.example_cluster.token
        k3s_version = yoshik3s_cluster.example_cluster.k3s_version
    }

    node_connection = {
        host = "localhost"
        port = "2222"
        user = "sshuser"
        password = "password"
    }

    node_options = [
		"--disable traefik",
		"--node-label node_type=master",
		"--snapshotter native",
    ]
}