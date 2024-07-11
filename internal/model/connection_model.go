package model

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type YoshiK3SConnectionModel struct {
	Host                 types.String `tfsdk:"host"`
	Port                 types.String `tfsdk:"port"`
	User                 types.String `tfsdk:"user"`
	Password             types.String `tfsdk:"password"`
	PrivateKey           types.String `tfsdk:"private_key"`
	PrivateKeyPassphrase types.String `tfsdk:"private_key_passphrase"`
}

var YoshiK3SConnectionModelSchema = map[string]schema.Attribute{
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
		Sensitive:           true,
	},
	"private_key": schema.StringAttribute{
		MarkdownDescription: "The SSH private key of the master node.",
		Optional:            true,
		Sensitive:           true,
	},
	"private_key_passphrase": schema.StringAttribute{
		MarkdownDescription: "The passphrase for the SSH private key of the master node.",
		Optional:            true,
		Sensitive:           true,
	},
}
