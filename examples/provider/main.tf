terraform {
    required_providers {
        yoshik3s = {
            source = "registry.terraform.io/HideyoshiNakazone/yoshi-k3s"
        }
    }
}

provider "yoshik3s" {}


resource "yoshik3s_master_node" "master_node" {
    version = "v1.21.4+k3s1"
    token = "secret_cluster_token"

    node_connection = {
        host = "127.0.0.1"
        port = "2222"
        user = "sshuser"
        password = "password"
    }
}