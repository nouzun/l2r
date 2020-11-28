package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nouzun/l2r/pkg/database"
	"github.com/nouzun/l2r/pkg/service"
)

type Application struct {
	Service service.Service
}

func (app *Application) Init(router *gin.Engine, database *database.Database) {

	app.Service = service.Service{
		DB: database,
	}

	fs := http.Dir("./swagger-ui/")
	router.StaticFS("/swagger-ui/", fs)

	v0UserGroup := router.Group(("v0/"))
	{
		v0UserGroup.GET("/words", app.GetWords)
	}

}

func (app *Application) GetWords(context *gin.Context) {

	words, err := app.Service.GetWords(context)
	if err != nil {
		return
	}

	context.JSON(200, words)
}
