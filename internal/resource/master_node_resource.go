// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"context"
	"fmt"
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
var _ resource.Resource = &YoshiK3SMasterNodeResource{}
var _ resource.ResourceWithImportState = &YoshiK3SMasterNodeResource{}

func NewYoshiK3SMasterNodeResource() resource.Resource {
	return &YoshiK3SMasterNodeResource{}
}

// YoshiK3SMasterNodeResource defines the resource implementation.
type YoshiK3SMasterNodeResource struct{}

func (r *YoshiK3SMasterNodeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_master_node"
}

func (r *YoshiK3SMasterNodeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "K3S Master Node Resource",

		Attributes: model.YoshiK3SMasterNodeResourceModelSchema,
	}
}

func (r *YoshiK3SMasterNodeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	//	No configuration is needed for this resource.
}

func (r *YoshiK3SMasterNodeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data model.YoshiK3SMasterNodeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Add a warning writting all data: Kubeconfig, Cluster, Connection, Options
	resp.Diagnostics.AddWarning("Info KUBECONFIG", fmt.Sprintf("Kubeconfig: %s", data.Kubeconfig))
	resp.Diagnostics.AddWarning("Info Cluster", fmt.Sprintf("Cluster: %s", data.Cluster.String()))
	resp.Diagnostics.AddWarning("Info Connection", fmt.Sprintf("Connection: %s", data.Connection.String()))
	resp.Diagnostics.AddWarning("Info Options", fmt.Sprintf("Options: %s", data.Options.String()))

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

	kubeconfig, err := client.ConfigureMasterNode(
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
	data.Kubeconfig = types.StringValue(string((*kubeconfig)[:]))

	//// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *YoshiK3SMasterNodeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data model.YoshiK3SMasterNodeResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *YoshiK3SMasterNodeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data model.YoshiK3SMasterNodeResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Add a warning writting all data: Kubeconfig, Cluster, Connection, Options
	resp.Diagnostics.AddWarning("Info", fmt.Sprintf("Kubeconfig: %s", data.Kubeconfig))
	resp.Diagnostics.AddWarning("Info", fmt.Sprintf("Cluster: %s", data.Cluster.String()))
	resp.Diagnostics.AddWarning("Info", fmt.Sprintf("Connection: %s", data.Connection.String()))
	resp.Diagnostics.AddWarning("Info", fmt.Sprintf("Options: %s", data.Options.String()))

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

	kubeconfig, err := client.ConfigureMasterNode(
		*nodeConfig,
		options,
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to update master node", err.Error())
		return
	}
	data.Kubeconfig = types.StringValue(string((*kubeconfig)[:]))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *YoshiK3SMasterNodeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data model.YoshiK3SMasterNodeResourceModel

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

	err := client.DestroyMasterNode(
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

func (r *YoshiK3SMasterNodeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *YoshiK3SMasterNodeResource) createClientFromModel(data model.YoshiK3SMasterNodeResourceModel) *cluster.K3sCluster {
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
	k3sServerAddress := clusterModel.ClusterServerAddress.ValueString()

	return cluster.NewK3sClientWithVersion(k3sVersion, k3sToken, k3sServerAddress)
}

func (r *YoshiK3SMasterNodeResource) createNodeConfigFromModel(data model.YoshiK3SMasterNodeResourceModel) *resources.NodeConfig {
	return resources.NewNodeConfig(
		data.Connection.Attributes()["host"].String(),
		r.createSshConfigFromModel(data),
	)
}

func (r *YoshiK3SMasterNodeResource) createSshConfigFromModel(data model.YoshiK3SMasterNodeResourceModel) *ssh_handler.SshConfig {
	connectionModel := r.parseSshConnectionModel(data.Connection)
	if connectionModel == nil {
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

func (r *YoshiK3SMasterNodeResource) parseSshConnectionModel(data types.Object) *model.YoshiK3SConnectionModel {
	if data.IsNull() || data.IsUnknown() {
		return nil
	}

	var connectionModel model.YoshiK3SConnectionModel
	diags := data.As(context.Background(), &connectionModel, basetypes.ObjectAsOptions{})
	if diags.HasError() {
		return nil
	}

	return &connectionModel
}

func (r *YoshiK3SMasterNodeResource) createNodeOptionsFromModel(model model.YoshiK3SMasterNodeResourceModel) []string {
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
