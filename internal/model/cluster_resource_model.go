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

	ClusterName          types.String `tfsdk:"name"`
	ClusterToken         types.String `tfsdk:"token"`
	ClusterServerAddress types.String `tfsdk:"server_address"`
	ClusterVersion       types.String `tfsdk:"k3s_version"`
}

var clusterResourceDescriptions = map[string]string{
	"id":             "The ID of the K3S Cluster.",
	"name":           "The name of the K3S Cluster.",
	"token":          "The token of K3S to be used in the configuration of the K3S Cluster.",
	"server_address": "The server address of the K3S Cluster.",
	"k3s_version":    "The version of K3S to be used in the configuration of the K3S Cluster.",
}

var YoshiK3SClusterResourceModelSchema = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		MarkdownDescription: clusterResourceDescriptions["id"],
		Description:         clusterResourceDescriptions["id"],
		Computed:            true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
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
		Sensitive:           true,
	},
	"server_address": schema.StringAttribute{
		MarkdownDescription: clusterResourceDescriptions["server_address"],
		Description:         clusterResourceDescriptions["server_address"],
		Required:            true,
	},
	"k3s_version": schema.StringAttribute{
		MarkdownDescription: clusterResourceDescriptions["k3s_version"],
		Description:         clusterResourceDescriptions["k3s_version"],
		Optional:            true,
	},
}
