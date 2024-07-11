package model

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// YoshiK3SWorkerNodeResourceModel describes the resource data model.
type YoshiK3SWorkerNodeResourceModel struct {
	Id types.String `tfsdk:"id"`

	MasterNodeServerAddress types.String `tfsdk:"master_server_address"`

	Cluster    types.Object `tfsdk:"cluster"`
	Connection types.Object `tfsdk:"node_connection"`

	Options types.List `tfsdk:"node_options"`
}

var YoshiK3SWorkerNodeResourceModelSchema = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		MarkdownDescription: "The ID of the master node.",
		Computed:            true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"master_server_address": schema.StringAttribute{
		MarkdownDescription: "The address of the master node.",
		Required:            true,
	},
	"cluster": schema.SingleNestedAttribute{
		MarkdownDescription: "The cluster to which the master node belongs.",
		Required:            true,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The ID of the K3S Cluster.",
				Optional:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the K3S Cluster.",
				Optional:            true,
			},
			"token": schema.StringAttribute{
				MarkdownDescription: "The token	of the cluster to which the master node belongs.",
				Required:            true,
			},
			"k3s_version": schema.StringAttribute{
				MarkdownDescription: "The version of the cluster to which the master node belongs.",
				Optional:            true,
			},
		},
	},
	"node_connection": schema.SingleNestedAttribute{
		MarkdownDescription: "The connection details of the master node.",
		Required:            true,
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				MarkdownDescription: "The hostname or IP address of the master node.",
				Required:            true,
			},
			"port": schema.StringAttribute{
				MarkdownDescription: "The SSH port of the master node.",
				Required:            true,
			},
			"user": schema.StringAttribute{
				MarkdownDescription: "The SSH user of the master node.",
				Required:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "The SSH password of the master node.",
				Optional:            true,
			},
			"private_key": schema.StringAttribute{
				MarkdownDescription: "The SSH private key of the master node.",
				Optional:            true,
			},
			"private_key_passphrase": schema.StringAttribute{
				MarkdownDescription: "The passphrase for the SSH private key of the master node.",
				Optional:            true,
			},
		},
	},
	"node_options": schema.ListAttribute{
		MarkdownDescription: "The options of the master node.",
		ElementType:         types.StringType,
		Optional:            true,
	},
}
