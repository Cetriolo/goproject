package boostrap

import (
	"v2/internal/database/migration"
	log "v2/logger"

	"v2/pkg/config"
	"v2/pkg/database"
)

func Migrate() {
	config.Set()
	database.Connect()
	log.Log.Info("Configs: ", config.Get())

	log.Log.Info("Starting Migration...")
	migration.Migrate()
}
