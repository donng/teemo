package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/donng/teemo/model"
	"github.com/gin-gonic/gin"
)

func GetWikiPedia(c *gin.Context) {
	date := time.Now().Format("20060102")

	wiki, err := model.GetWikiPediaByDate(date)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "1",
			"data": wiki,
			"message": "",
		})
		panic(fmt.Sprintf("get wikipedia error, err: %s", err))
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": wiki,
		"message": "",
	})
}
