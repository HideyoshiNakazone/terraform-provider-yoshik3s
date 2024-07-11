package model

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// YoshiK3SClusterResourceModel describes the resource data model.
type YoshiK3SClusterResourceModel struct {
	Id types.String `tfsdk:"id"`

	ClusterName    types.String `tfsdk:"name"`
	ClusterToken   types.String `tfsdk:"token"`
	ClusterVersion types.String `tfsdk:"k3s_version"`
}

var YoshiK3SClusterResourceModelSchema = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		MarkdownDescription: "The ID of the K3S Cluster.",
		Computed:            true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"name": schema.StringAttribute{
		MarkdownDescription: "The name of the K3S Cluster.",
		Optional:            true,
	},
	"token": schema.StringAttribute{
		MarkdownDescription: "The token of K3S to be used in the configuration of the K3S Cluster.",
		Required:            true,
		Sensitive:           true,
	},
	"k3s_version": schema.StringAttribute{
		MarkdownDescription: "The version of K3S to be used in the configuration of the K3S Cluster.",
		Optional:            true,
	},
}
