package postgres

func CreateClientFactory() *PostgrestClient {
	configClient := CreateConfigClient()
	migrationClient := NewMigrationClient()
	postgrestClient := NewPostgrestClient(configClient, &migrationClient)

	return &postgrestClient
}
