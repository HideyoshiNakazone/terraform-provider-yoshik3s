// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	internalresource "github.com/HideyoshiNakazone/terraform-provider-yoshi-k3s/internal/resource"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

type YoshiK3SProvider struct {
	version string
}

var _ provider.Provider = &YoshiK3SProvider{}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &YoshiK3SProvider{
			version: version,
		}
	}
}

// Metadata returns the provider type name.
func (p *YoshiK3SProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "yoshik3s"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *YoshiK3SProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (p *YoshiK3SProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *YoshiK3SProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		internalresource.NewYoshiK3SClusterResource,
		internalresource.NewYoshiK3SMasterNodeResource,
	}
}

func (p *YoshiK3SProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}
