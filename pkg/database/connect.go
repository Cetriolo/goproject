package database

import (
	"fmt"
	log "v2/logger"
	"v2/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	configs := config.Get()
	log.Log.Info("Connecting to database...")
	log.Log.Info("Database Configs: ", configs)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configs.DB.Username,
		configs.DB.Password,
		configs.DB.Host,
		configs.DB.Port,
		configs.DB.Name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Log.Fatal("error connecting to database:", err)
		return
	}

	DB = db

}
