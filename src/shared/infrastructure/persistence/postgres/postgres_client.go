package postgres

type PostgrestClient struct {
	ConfigClient    *ConfigClient
	MigrationClient *MigrationClient
}

func NewPostgrestClient(
	configClient *ConfigClient,
	migrationClient *MigrationClient,
) PostgrestClient {
	return PostgrestClient{
		ConfigClient:    configClient,
		MigrationClient: migrationClient,
	}
}
