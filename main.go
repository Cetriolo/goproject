package main

import (
	"v2/cmd"
	log "v2/logger"
)

func main() {
	log.Log.Info("Starting application...")
	cmd.Execute()

}
