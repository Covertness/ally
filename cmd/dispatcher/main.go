package main

import (
	"log"

	"github.com/Covertness/ally/pkg/config"
	"github.com/Covertness/ally/pkg/order"
	"github.com/Covertness/ally/pkg/transaction"

	"github.com/Covertness/ally/pkg/messagequeue"
	"github.com/Covertness/ally/pkg/storage"

	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := storage.InitDB(config.GetDBDialect(), config.GetDBConnectArgs())
	if err != nil {
		log.Fatalf("db init err: %v", err)
		return
	}
	defer db.Close()

	err = messagequeue.InitQueue()
	if err != nil {
		log.Fatalf("message queue init err: %v", err)
		return
	}

	jobrunner.Start()

	err = order.JobSchedule()
	if err != nil {
		log.Fatalf("schedule order err: %v", err)
		return
	}

	err = transaction.JobSchedule()
	if err != nil {
		log.Fatalf("schedule transaction err: %v", err)
		return
	}

	r := gin.Default()
	// Resource to return the JSON data
	r.GET("/jobrunner/json", jobJSON)
	// Load template file location relative to the current working directory
	r.LoadHTMLGlob("views/jobrunner/*")
	// Returns html page at given endpoint based on the loaded
	// template from above
	r.GET("/jobrunner/html", jobHTML)
	r.Run(":9002")

	log.Println("dispatcher is working...")
}

func jobJSON(c *gin.Context) {
	// returns a map[string]interface{} that can be marshalled as JSON
	c.JSON(200, jobrunner.StatusJson())
}

func jobHTML(c *gin.Context) {
	// Returns the template data pre-parsed
	c.HTML(200, "Status.html", jobrunner.StatusPage())
}
