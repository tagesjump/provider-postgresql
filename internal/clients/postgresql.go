package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/tagesjump/provider-postgresql/apis/v1beta1"
)

const (
	// (Optional) The driver to use. Valid values are: postgres: Default value, use lib/pq, awspostgres: Use GoCloud for AWS, gcppostgres: Use GoCloud for GCP
	scheme = "scheme"

	// (Required) The address for the postgresql server connection, see GoCloud for specific format.
	host = "host"

	// (Optional) The port for the postgresql server connection. The default is 5432.
	port = "port"

	// (Optional) Database to connect to. The default is postgres.
	database = "database"

	// (Required) Username for the server connection.
	username = "username"

	// (Optional) Password for the server connection.
	password = "password"

	// (Optional) Username of the user in the database if different than connection username (See user name maps).
	databaseUsername = "database_username"

	// (Optional) Should be set to false if the user to connect is not a PostgreSQL superuser (as is the case in AWS RDS or GCP SQL). In this case, some features might be disabled (e.g.: Refreshing state password from database).
	superuser = "superuser"

	// (Optional) Set the priority for an SSL connection to the server. Valid values for sslmode are (note: prefer is not supported by Go's lib/pq)): disable - No SSL, require - Always SSL (the default, also skip verification) verify-ca - Always SSL (verify that the certificate presented by the server was signed by a trusted CA) verify-full - Always SSL (verify that the certification presented by the server was signed by a trusted CA and the server host name matches the one in the certificate) Additional information on the options and their implications can be seen in the libpq(3) SSL guide.
	sslmode = "sslmode"

	// (Optional) - The SSL server root certificate file path. The file must contain PEM encoded data.
	sslrootcert = "sslrootcert"

	// (Optional) Maximum wait for connection, in seconds. The default is 180s. Zero or not specified means wait indefinitely.
	connectTimeout = "connect_timeout"

	// (Optional) Set the maximum number of open connections to the database. The default is 20. Zero means unlimited open connections.
	maxConnections = "max_connections"

	// (Optional) Specify a hint to Terraform regarding the expected version that the provider will be talking with. This is a required hint in order for Terraform to talk with an ancient version of PostgreSQL. This parameter is expected to be a PostgreSQL Version or current. Once a connection has been established, Terraform will fingerprint the actual version. Default: 9.0.0.
	expectedVersion = "expected_version"

	// (Optional) If set to true, call the AWS RDS API to grab a temporary password, using AWS Credentials from the environment (or the given profile, see aws_rds_iam_profile)
	awsRdsIamAuth = "aws_rds_iam_auth"

	// (Optional) The AWS IAM Profile to use while using AWS RDS IAM Auth.
	awsRdsIamProfile = "aws_rds_iam_profile"

	// (Optional) The AWS region to use while using AWS RDS IAM Auth.
	awsRdsIamRegion = "aws_rds_iam_region"

	// (Optional) If set to true, call the Azure OAuth token endpoint for temporary token
	azureIdentityAuth = "azure_identity_auth"

	// (Optional) (Required if azure_identity_auth is true) Azure tenant ID read more
	azureTenantId = "azure_tenant_id"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal postgresql credentials as JSON"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// set provider configuration
		ps.Configuration = map[string]any{}

		for _, setting := range []string{
			scheme, host, port, database, username, password, databaseUsername,
			superuser, sslmode, sslrootcert, connectTimeout, maxConnections,
			expectedVersion, awsRdsIamAuth, awsRdsIamProfile, awsRdsIamRegion,
			azureIdentityAuth, azureTenantId,
		} {
			if value, ok := creds[setting]; ok {
				ps.Configuration[setting] = value
			}
		}

		return ps, nil
	}
}
