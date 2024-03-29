package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/oneaushaf/go-broiler/database"
	"github.com/oneaushaf/go-broiler/initializers"
	"github.com/oneaushaf/go-broiler/routes"
)

func init() {
	initializers.LoadEnv()
	database.ConnectDatabase()
	database.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.MaxMultipartMemory = 8 << 20  // 8 MiB

	routes.AuthRoutes(r)
	routes.UserRoutes(r)
	routes.FarmRoutes(r)
	routes.RanchRoutes(r)
	routes.BatchRoutes(r)
	routes.WeighingRoutes(r)

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
