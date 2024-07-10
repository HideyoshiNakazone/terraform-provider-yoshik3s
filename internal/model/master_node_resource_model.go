package model

import "github.com/hashicorp/terraform-plugin-framework/types"

// YoshiK3SMasterNodeResourceModel describes the resource data model.
type YoshiK3SMasterNodeResourceModel struct {
	Token   types.String `tfsdk:"token"`
	Version types.String `tfsdk:"version"`

	Connection types.Object `tfsdk:"node_connection"`
}
