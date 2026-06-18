package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Sukanthan06/code-execution-engine/api/routes"
)

func main() {
	router := gin.Default()

	routes.SetupRoutes((router))

	router.Run(":8080")
}
