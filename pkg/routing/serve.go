package routing

import (
	"fmt"
	log "v2/logger"
	"v2/pkg/config"
)

func Serve() {

	r := GetRouter()
	configs := config.Get()
	err := r.Run(fmt.Sprintf("%s:%d", configs.Server.Host, configs.Server.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Log.Fatal("Error in routing: ", err)
	}

}
