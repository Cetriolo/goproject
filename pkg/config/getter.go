package config

import (
	"v2/config" // Update this import path to the actual location of your config package
)

func Get() config.Config {
	return configurations
}
