package main

import (
	"GO-1/config"
	"GO-1/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var taskCollection *mongo.Collection

func main() {
	// koneksi
	config.ConnectDB()
	r := gin.Default()

	// panggil fungsi route
	routes.SetupTaskRouters(r)
	r.Run(":8080")
}
