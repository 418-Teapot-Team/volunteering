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

		api.GET("/projects", app.getProjects)
		api.POST("/projects", app.createProject)

		api.DELETE("/projects", app.deleteProject)
		api.DELETE("/projects/task", app.deleteProjectTask)

		api.GET("/tasks", app.getUserTasks)
		api.POST("/tasks", app.createTask)
		api.DELETE("/tasks", app.deleteTask)

		api.POST("tasks/share", app.shareTask)
		api.GET("tasks/shared", app.getSharedTasks)

		api.POST("/done-volunteer", app.markAsDoneVolunteer)
		api.POST("/done-employer", app.markAsDoneEmployer)

		// apply
		// accept / deny
		// statistic
		// mark as done (points)
		// verified and unverified
		// if verified, volunteer gets bonuses

	}

	router.Run(":" + port)
}
