/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	resourcePrefix = "portlabs"
	modulePath     = "github.com/haarchri/provider-portlabs"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("getport.io"),
		ujconfig.WithIncludeList(ResourcesWithExternalNameConfig()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		// null.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// ResourcesWithExternalNameConfig returns the list of resources that have external
// name configured in ExternalNameConfigs table.
func ResourcesWithExternalNameConfig() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// Expected format is regex and we'd like to have exact matches.
		l[i] = name + "$"
		i++
	}
	return l
}
