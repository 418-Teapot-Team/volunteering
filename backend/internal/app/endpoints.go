package app

import (
	"github.com/gin-gonic/gin"
)

func StartServer(app *Handler, port string) {

	router := gin.Default()
	router.Use(CORSMiddleware())
	auth := router.Group("/auth/v1")
	auth.Use(CORSMiddleware())
	{
		auth.POST("/sign-up", app.signUp)
		auth.POST("/sign-in", app.signIn)

	}

	api := router.Group("/api/v1", app.authMiddleware)
	{
		api.GET("/who-am-i", app.whoAmI)
	}

	router.Run(":" + port)
}
