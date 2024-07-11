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

var connectResourceDescriptions = map[string]string{
	"host":                   "The hostname or IP address of the master node.",
	"port":                   "The SSH port of the master node.",
	"user":                   "The SSH user of the master node.",
	"password":               "The SSH password of the master node.",
	"private_key":            "The SSH private key of the master node.",
	"private_key_passphrase": "The passphrase for the SSH private key of the master node.",
}

var YoshiK3SConnectionModelSchema = map[string]schema.Attribute{
	"host": schema.StringAttribute{
		Description:         connectResourceDescriptions["host"],
		MarkdownDescription: connectResourceDescriptions["host"],
		Required:            true,
	},
	"port": schema.StringAttribute{
		Description:         connectResourceDescriptions["port"],
		MarkdownDescription: connectResourceDescriptions["port"],
		Required:            true,
	},
	"user": schema.StringAttribute{
		Description:         connectResourceDescriptions["user"],
		MarkdownDescription: connectResourceDescriptions["user"],
		Required:            true,
	},
	"password": schema.StringAttribute{
		Description:         connectResourceDescriptions["password"],
		MarkdownDescription: connectResourceDescriptions["password"],
		Optional:            true,
		Sensitive:           true,
	},
	"private_key": schema.StringAttribute{
		Description:         connectResourceDescriptions["private_key"],
		MarkdownDescription: connectResourceDescriptions["private_key"],
		Optional:            true,
		Sensitive:           true,
	},
	"private_key_passphrase": schema.StringAttribute{
		Description:         connectResourceDescriptions["private_key_passphrase"],
		MarkdownDescription: connectResourceDescriptions["private_key_passphrase"],
		Optional:            true,
		Sensitive:           true,
	},
}
