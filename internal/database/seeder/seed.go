package seeder

import (
	"fmt"
	articleModels "v2/internal/modules/article/models"
	userModels "v2/internal/modules/user/models"
	log "v2/logger"
	"v2/pkg/database"

	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	db := database.Connection()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		log.Log.Fatal("Hashing password failed:", err)
	}
	user := userModels.User{Name: "John Doe", Email: "john@example.com", Password: string(hashedPassword)}
	// Seed the database with initial data
	if err := db.Create(&user).Error; err != nil {
		log.Log.Fatal("Seeding user failed:", err)
	}
	log.Log.Info("User seeded:", user)

	for i := 1; i <= 5; i++ {
		article := articleModels.Article{Title: fmt.Sprintf("Article %d", i), Content: fmt.Sprintf("This is the content of article %d.", i), UserID: user.ID}
		if err := db.Create(&article).Error; err != nil {
			log.Log.Fatal("Seeding article failed:", err)
		}
		log.Log.Info("Article seeded:", article)
	}
}
