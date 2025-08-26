package migration

import (
	articleModels "v2/internal/modules/article/models"
	userModels "v2/internal/modules/user/models"
	log "v2/logger"
	"v2/pkg/database"
)

func Migrate() {

	db := database.Connection()
	// Migrate the schema

	err := db.AutoMigrate(&userModels.User{}, &articleModels.Article{})
	if err != nil {
		log.Log.Fatal("Migration failed:", err)
	}
}
