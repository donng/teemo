package router

import (
	"github.com/donng/teemo/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	addRoutes(r)

	return r
}

func addRoutes(r *gin.Engine) {
	router := r.Group("/api")

	router.Use()
	{
		router.GET("/wikipedia", api.GetWikiPedia)
	}
}
