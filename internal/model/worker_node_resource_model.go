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
		Description:         nodeResourceDescriptions["id"],
		MarkdownDescription: nodeResourceDescriptions["id"],
		Computed:            true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"master_server_address": schema.StringAttribute{
		Description:         "The address of the master node.",
		MarkdownDescription: "The address of the master node.",
		Required:            true,
	},
	"cluster": schema.SingleNestedAttribute{
		Description:         nodeResourceDescriptions["cluster"],
		MarkdownDescription: nodeResourceDescriptions["cluster"],
		Required:            true,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: clusterResourceDescriptions["id"],
				Description:         clusterResourceDescriptions["id"],
				Optional:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: clusterResourceDescriptions["name"],
				Description:         clusterResourceDescriptions["name"],
				Optional:            true,
			},
			"token": schema.StringAttribute{
				MarkdownDescription: clusterResourceDescriptions["token"],
				Description:         clusterResourceDescriptions["token"],
				Required:            true,
			},
			"k3s_version": schema.StringAttribute{
				MarkdownDescription: clusterResourceDescriptions["k3s_version"],
				Description:         clusterResourceDescriptions["k3s_version"],
				Optional:            true,
			},
		},
	},
	"node_connection": schema.SingleNestedAttribute{
		Description:         nodeResourceDescriptions["node_connection"],
		MarkdownDescription: nodeResourceDescriptions["node_connection"],
		Required:            true,
		Attributes:          YoshiK3SConnectionModelSchema,
	},
	"node_options": schema.ListAttribute{
		Description:         nodeResourceDescriptions["node_options"],
		MarkdownDescription: nodeResourceDescriptions["node_options"],
		ElementType:         types.StringType,
		Optional:            true,
	},
}
