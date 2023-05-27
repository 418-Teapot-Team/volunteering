package app

import (
	"github.com/gin-gonic/gin"
)

func StartServer(app *Handler, port string) {

	router := gin.Default()

	auth := router.Group("/auth/v1")
	auth.Use(CORSMiddleware())
	{
		auth.POST("/sign-up", app.signUp)
		auth.POST("/sign-in", app.signIn)
	}

	router.Run(":" + port)
}
