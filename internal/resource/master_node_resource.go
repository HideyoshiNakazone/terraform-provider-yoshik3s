// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"context"
	"fmt"
	"github.com/HideyoshiNakazone/terraform-provider-yoshi-k3s/internal/model"
	"github.com/HideyoshiNakazone/yoshi-k3s/pkg/client"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &YoshiK3SMasterNodeResource{}
var _ resource.ResourceWithImportState = &YoshiK3SMasterNodeResource{}

func NewYoshiK3SMasterNodeResource() resource.Resource {
	return &YoshiK3SMasterNodeResource{}
}

// YoshiK3SMasterNodeResource defines the resource implementation.
type YoshiK3SMasterNodeResource struct {
	k3sClient *client.K3sClient
}

func (r *YoshiK3SMasterNodeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_master_node"
}

func (r *YoshiK3SMasterNodeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "K3S Master Node Resource",

		Attributes: map[string]schema.Attribute{
			"version": schema.StringAttribute{
				MarkdownDescription: "The version of K3S to install on the master node.",
				Required:            false,
			},
			"token": schema.StringAttribute{
				MarkdownDescription: "The token used to join the master node to the cluster.",
				Required:            true,
			},
			"node_connection": schema.SingleNestedAttribute{
				Required: true,
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
						Required:            false,
					},
					"private_key": schema.StringAttribute{
						MarkdownDescription: "The SSH private key of the master node.",
						Required:            false,
					},
					"private_key_passphrase": schema.StringAttribute{
						MarkdownDescription: "The passphrase for the SSH private key of the master node.",
						Required:            false,
					},
				},
			},
		},
	}
}

func (r *YoshiK3SMasterNodeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	k3sClient, ok := req.ProviderData.(*client.K3sClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.K3sClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.k3sClient = k3sClient
}

func (r *YoshiK3SMasterNodeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data model.YoshiK3SMasterNodeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create example, got error: %s", err))
	//     return
	// }

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "created a resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *YoshiK3SMasterNodeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data model.YoshiK3SMasterNodeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read example, got error: %s", err))
	//     return
	// }

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *YoshiK3SMasterNodeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data model.YoshiK3SMasterNodeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update example, got error: %s", err))
	//     return
	// }

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *YoshiK3SMasterNodeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data model.YoshiK3SMasterNodeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete example, got error: %s", err))
	//     return
	// }
}

func (r *YoshiK3SMasterNodeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
