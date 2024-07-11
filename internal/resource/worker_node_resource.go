// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"context"
	"github.com/HideyoshiNakazone/terraform-provider-yoshik3s/internal/model"
	"github.com/HideyoshiNakazone/yoshi-k3s/pkg/cluster"
	"github.com/HideyoshiNakazone/yoshi-k3s/pkg/resources"
	"github.com/HideyoshiNakazone/yoshi-k3s/pkg/ssh_handler"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &YoshiK3SWorkerNodeResource{}
var _ resource.ResourceWithImportState = &YoshiK3SWorkerNodeResource{}

func NewYoshiK3SWorkerNodeResource() resource.Resource {
	return &YoshiK3SWorkerNodeResource{}
}

// YoshiK3SWorkerNodeResource defines the resource implementation.
type YoshiK3SWorkerNodeResource struct{}

func (r *YoshiK3SWorkerNodeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_worker_node"
}

func (r *YoshiK3SWorkerNodeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "K3S Master Node Resource",

		Attributes: model.YoshiK3SWorkerNodeResourceModelSchema,
	}
}

func (r *YoshiK3SWorkerNodeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	//	No configuration is needed for this resource.
}

func (r *YoshiK3SWorkerNodeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data model.YoshiK3SWorkerNodeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	client := r.createClientFromModel(data)
	if client == nil {
		resp.Diagnostics.AddError(
			"Failed to create a master node",
			"Invalid cluster configuration. Please check the cluster configuration.",
		)
		return
	}
	nodeConfig := r.createNodeConfigFromModel(data)
	if nodeConfig == nil {
		resp.Diagnostics.AddError(
			"Failed to create a master node",
			"Invalid node configuration. Please check the node configuration.",
		)
		return
	}
	options := r.createNodeOptionsFromModel(data)

	err := client.ConfigureWorkerNode(
		*nodeConfig,
		options,
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to create a master node", err.Error())
		return
	}

	//// Write logs using the tflog package
	//// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "created a resource")
	data.Id = types.StringValue(data.Connection.Attributes()["host"].String())
	//
	//// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *YoshiK3SWorkerNodeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data model.YoshiK3SWorkerNodeResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *YoshiK3SWorkerNodeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data model.YoshiK3SWorkerNodeResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	client := r.createClientFromModel(data)
	if client == nil {
		resp.Diagnostics.AddError(
			"Failed to create a master node",
			"Invalid cluster configuration. Please check the cluster configuration.",
		)
		return
	}
	nodeConfig := r.createNodeConfigFromModel(data)
	if nodeConfig == nil {
		resp.Diagnostics.AddError(
			"Failed to create a master node",
			"Invalid node configuration. Please check the node configuration.",
		)
		return
	}
	options := r.createNodeOptionsFromModel(data)

	err := client.ConfigureWorkerNode(
		*nodeConfig,
		options,
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to update master node", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *YoshiK3SWorkerNodeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data model.YoshiK3SWorkerNodeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	client := r.createClientFromModel(data)
	if client == nil {
		resp.Diagnostics.AddError(
			"Failed to create a master node",
			"Invalid cluster configuration. Please check the cluster configuration.",
		)
		return
	}
	nodeConfig := r.createNodeConfigFromModel(data)
	if nodeConfig == nil {
		resp.Diagnostics.AddError(
			"Failed to create a master node",
			"Invalid node configuration. Please check the node configuration.",
		)
		return
	}

	err := client.DestroyWorkerNode(
		*nodeConfig,
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to delete a master node", err.Error())
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *YoshiK3SWorkerNodeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *YoshiK3SWorkerNodeResource) createClientFromModel(data model.YoshiK3SWorkerNodeResourceModel) *cluster.K3sCluster {
	if data.Cluster.IsNull() || data.Cluster.IsUnknown() {
		return nil
	}

	var clusterModel model.YoshiK3SClusterResourceModel
	diags := data.Cluster.As(context.Background(), &clusterModel, basetypes.ObjectAsOptions{})
	if diags.HasError() {
		return nil
	}

	k3sVersion := clusterModel.ClusterVersion.ValueString()
	k3sToken := clusterModel.ClusterToken.ValueString()

	return cluster.NewK3sClientWithVersion(k3sVersion, k3sToken)
}

func (r *YoshiK3SWorkerNodeResource) createNodeConfigFromModel(data model.YoshiK3SWorkerNodeResourceModel) *resources.K3sWorkerNodeConfig {
	return resources.NewK3sWorkerNodeConfig(
		data.MasterNodeServerAddress.ValueString(),
		r.createSshConfigFromModel(data),
	)
}

func (r *YoshiK3SWorkerNodeResource) createSshConfigFromModel(data model.YoshiK3SWorkerNodeResourceModel) *ssh_handler.SshConfig {
	if data.Connection.IsNull() || data.Connection.IsUnknown() {
		return nil
	}

	var connectionModel model.YoshiK3SConnectionModel
	diags := data.Connection.As(context.Background(), &connectionModel, basetypes.ObjectAsOptions{})
	if diags.HasError() {
		return nil
	}

	var password string
	if connectionModel.Password.IsNull() || connectionModel.Password.IsUnknown() {
		password = ""
	} else {
		password = connectionModel.Password.ValueString()
	}

	var privateKey string
	if connectionModel.PrivateKey.IsNull() || connectionModel.PrivateKey.IsUnknown() {
		privateKey = ""
	} else {
		privateKey = connectionModel.PrivateKey.ValueString()
	}

	var privateKeyPassphrase string
	if connectionModel.PrivateKeyPassphrase.IsNull() || connectionModel.PrivateKeyPassphrase.IsUnknown() {
		privateKeyPassphrase = ""
	} else {
		privateKeyPassphrase = connectionModel.PrivateKeyPassphrase.ValueString()
	}

	return ssh_handler.NewSshConfig(
		connectionModel.Host.ValueString(),
		connectionModel.Port.ValueString(),
		connectionModel.User.ValueString(),
		password,
		privateKey,
		privateKeyPassphrase,
	)
}

func (r *YoshiK3SWorkerNodeResource) createNodeOptionsFromModel(model model.YoshiK3SWorkerNodeResourceModel) []string {
	if model.Options.IsNull() || model.Options.IsUnknown() {
		return []string{}
	}

	elements := make([]types.String, 0, len(model.Options.Elements()))
	diags := model.Options.ElementsAs(context.Background(), &elements, false)
	if diags.HasError() {
		return []string{}
	}

	var nodeOptions []string
	for _, element := range elements {
		nodeOptions = append(nodeOptions, element.ValueString())
	}

	return nodeOptions
}
