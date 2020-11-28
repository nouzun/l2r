package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nouzun/l2r/pkg/api"
	db "github.com/nouzun/l2r/pkg/database"

	log "github.com/sirupsen/logrus"
)

func main() {

	app := api.Application{}

	database, err := db.ConnectDatabase()
	if err != nil {
		log.Error(err)
	}

	router := gin.Default()
	app.Init(router, database)

	log.Infof("Starting application")
	err = router.Run(":10000")

	if err == nil {
		log.Error("router.run caused errors")
		log.Error(err)
		os.Exit(1)
	}

}
