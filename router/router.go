package router

import (
	"fmt"
	"github.com/donng/teemo/api"
	"github.com/gin-gonic/gin"
	"time"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
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
