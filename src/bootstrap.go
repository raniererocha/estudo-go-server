package bootstrap

import (
	"server/src/routes"

	"github.com/gin-gonic/gin"
)

func Start(port string) {
	app := gin.Default()
	routes.RegisterRoutes(app)
	app.Run(port)
}
