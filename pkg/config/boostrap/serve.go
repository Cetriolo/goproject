package boostrap

import (
	log "v2/logger"

	"v2/pkg/config"
	"v2/pkg/database"
	htmlHandler "v2/pkg/html"
	"v2/pkg/routing"
	"v2/pkg/static"
)

func Serve() {
	config.Set()
	database.Connect()
	log.Log.Info("Configs: ", config.Get())
	routing.Init()
	static.LoadStaticFiles(routing.GetRouter())
	htmlHandler.LoadHtml(routing.GetRouter())
	routing.RegisterRoutes()
	log.Log.Info("Starting application...")
	routing.Serve()
}
