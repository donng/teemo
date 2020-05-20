package main

import (
	"fmt"
	"github.com/donng/teemo/pkg/logging"
	"github.com/donng/teemo/pkg/setting"
	"github.com/donng/teemo/router"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	config := setting.Setting

	setDefaultConfig(config.App.Debug)

	r := router.InitRouter()
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Server.HttpPort),
		Handler:      r,
		ReadTimeout:  time.Duration(config.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.Server.WriteTimeout) * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}

func setDefaultConfig(debug bool) {
	if debug {
		gin.SetMode(gin.DebugMode)
		return
	}

	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	gin.DefaultWriter = logging.GetLogFile()
}
