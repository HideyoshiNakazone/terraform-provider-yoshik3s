package model

import "github.com/hashicorp/terraform-plugin-framework/types"

type YoshiK3SConnectionModel struct {
	Host                 types.String `tfsdk:"host"`
	Port                 types.String `tfsdk:"port"`
	User                 types.String `tfsdk:"user"`
	Password             types.String `tfsdk:"password"`
	PrivateKey           types.String `tfsdk:"private_key"`
	PrivateKeyPassphrase types.String `tfsdk:"private_key_passphrase"`
}
