# Terraform Provider YoshiK3S

This is a Terraform provider for managing [K3s](https://k3s.io/) clusters. 
It is based on the [YoshiK3S](https://github.com/HideyoshiNakazone/yoshi-k3s) Golang Package.

## Features
    
- [x] Create a K3s cluster
- [x] Create a K3s master node
- [x] Create a K3s worker node
- [x] Delete a K3s cluster
- [x] Delete a K3s master node
- [x] Delete a K3s worker node
- [x] Update a K3s cluster
- [x] Update a K3s master node
- [x] Update a K3s worker node
- [ ] ~~Validation of the Master Node on which the Worker Node is being created~~ 
  - (Not possible due to the lack of state management)

### Disclaimer

This project does not aim to be a full-fledged provider, but rather a simple way to manage K3s clusters using Terraform.
Therefore, all state management is done by Terraform itself and the provider does not have any state management capabilities,
by consequence it is possible to use the provider to create a invalid worker node without a master node or cluster resource - this is not a recommended action.

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.21

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:

```shell
go install
```

## Using the provider

```hcl
terraform {
  required_providers {
    yoshik3s = {
      source = "HideyoshiNakazone/yoshik3s"
      version = "{VERSION}"
    }
  }
}

provider "yoshik3s" {
  # No configuration needed
}
```

### Configuring the Cluster

This resource is used to share the `token` and `k3s_version` between the master and worker nodes, 
it is used to update the `k3s_version` or the `token` of all nodes at once. It is not necessary to use this resource,
but it is recommended to use it to avoid inconsistency between nodes.

```hcl
resource "yoshik3s_cluster" "example_cluster" {
  name        = "example-cluster"
  token       = "{K3S_TOKEN}"
  k3s_version = "{K3S_VERSION}"
}
```

### Configuring the Master Node

This resource is used to create and manage the configuration of a K3s master node.

```hcl
resource "yoshik3s_master_node" "example_master_node" {
  cluster = {
    token       = yoshik3s_cluster.example_cluster.token
    k3s_version = yoshik3s_cluster.example_cluster.k3s_version
  }

  node_connection = {
    host                    = "{NODE_CONNECTION_HOST}"
    port                    = "{NODE_CONNECTION_PORT}"
    user                    = "{NODE_CONNECTION_USER}"
    password                = "{NODE_CONNECTION_PASSWORD}"
    private_key             = "{NODE_CONNECTION_PRIVATE_KEY}"
    private_key_passphrase  = "{NODE_CONNECTION_PRIVATE_KEY_PASSPHRASE}"
  }

  node_options = [
    "{OPTION1}",
    "{OPTION2}",
    "{OPTION3}"
  ]
}
```

You must provide either `password` or `private_key` and `private_key_passphrase` in `node_connection`.

### Configuring the Worker Node

This resource is used to create and manage the configuration of a K3s worker node.

```hcl
resource "yoshik3s_worker_node" "example_worker_node" {
  master_server_address = "{MASTER_SERVER_ADDRESS}"

  cluster = {
    token       = yoshik3s_cluster.example_cluster.token
    k3s_version = yoshik3s_cluster.example_cluster.k3s_version
  }

  node_connection = {
    host                    = "{NODE_CONNECTION_HOST}"
    port                    = "{NODE_CONNECTION_PORT}"
    user                    = "{NODE_CONNECTION_USER}"
    password                = "{NODE_CONNECTION_PASSWORD}"
    private_key             = "{NODE_CONNECTION_PRIVATE_KEY}"
    private_key_passphrase  = "{NODE_CONNECTION_PRIVATE_KEY_PASSPHRASE}"
  }

  node_options = [
    "{OPTION1}",
    "{OPTION2}",
    "{OPTION3}"
  ]
}
```
This resource requires the `master_server_address` to be set to the address of the master node, 
it must be a valid **ip address** or a valid **host name**.

You must provide either `password` or `private_key` and `private_key_passphrase` in `node_connection`.


## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```shell
make testacc
```
