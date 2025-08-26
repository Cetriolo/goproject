package config

import (
	log "v2/logger"

	"github.com/spf13/viper"
)

func Set() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("config") // path to look for the config file in
	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Log.Println("Config file not found; ignoring error")
		} else {
			log.Log.Error("Error reading config file: ", err)
			log.Log.Fatal("Exiting application...")
		}
	}
	if err := viper.Unmarshal(&configurations); err != nil {
		log.Log.Fatal("Unable to decode into struct: ", err)
	}
}
