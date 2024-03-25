package config

import (
	"github.com/crossplane/upjet/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"postgresql_database":                  config.IdentifierFromProvider,
	"postgresql_default_privileges":        config.IdentifierFromProvider,
	"postgresql_extension":                 config.IdentifierFromProvider,
	"postgresql_function":                  config.IdentifierFromProvider,
	"postgresql_grant":                     config.IdentifierFromProvider,
	"postgresql_grant_role":                config.IdentifierFromProvider,
	"postgresql_physical_replication_slot": config.IdentifierFromProvider,
	"postgresql_publication":               config.IdentifierFromProvider,
	"postgresql_replication_slot":          config.IdentifierFromProvider,
	"postgresql_role":                      config.IdentifierFromProvider,
	"postgresql_schema":                    config.IdentifierFromProvider,
	"postgresql_server":                    config.IdentifierFromProvider,
	"postgresql_subscription":              config.IdentifierFromProvider,
	"postgresql_user_mapping":              config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
