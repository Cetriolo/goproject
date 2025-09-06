package boostrap

import (
	"v2/internal/database/seeder"
	log "v2/logger"

	"v2/pkg/config"
	"v2/pkg/database"
)

func Seed() {
	config.Set()
	database.Connect()
	log.Log.Info("Configs: ", config.Get())

	log.Log.Info("Starting Seeding...")
	seeder.Seed()
}
