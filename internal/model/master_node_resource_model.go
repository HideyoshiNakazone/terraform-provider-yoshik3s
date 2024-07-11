package model

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// YoshiK3SMasterNodeResourceModel describes the resource data model.
type YoshiK3SMasterNodeResourceModel struct {
	Id types.String `tfsdk:"id"`

	ServerAddress types.String `tfsdk:"server_address"`

	Cluster    types.Object `tfsdk:"cluster"`
	Connection types.Object `tfsdk:"node_connection"`

	Options types.List `tfsdk:"node_options"`
}

var YoshiK3SMasterNodeResourceModelSchema = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		MarkdownDescription: "The ID of the master node.",
		Computed:            true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"server_address": schema.StringAttribute{
		MarkdownDescription: "The address of the master node.",
		Computed:            true,
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
		Attributes:          YoshiK3SConnectionModelSchema,
	},
	"node_options": schema.ListAttribute{
		MarkdownDescription: "The options of the master node.",
		ElementType:         types.StringType,
		Optional:            true,
	},
}
