package main

import (
	"github.com/donng/teemo/cmd/wikipedia"
	"github.com/robfig/cron/v3"
	"log"
)

func main() {
	c := cron.New()
	c.AddFunc("@daily", func() {
		log.Println("Run sync wikipedia ...")
		wikipedia.Sync()
	})

	c.Start()
	select {}
}
