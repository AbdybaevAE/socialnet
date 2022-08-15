package migrations

import "embed"

//go:embed postgresql
var MigrationFS embed.FS
