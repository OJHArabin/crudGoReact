package main

import (
	"net/http"

	"./routers"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "done",
			})
		})
	}
	api.POST("/add/users", routers.Create)
	api.GET("/user/:id", routers.AuthMiddleware(), routers.Read)
	api.GET("/users/list", routers.AuthMiddleware(), routers.ReadAll)
	api.PUT("/update/user/:id", routers.AuthMiddleware(), routers.Update)
	api.DELETE("/delete/user/:id", routers.AuthMiddleware(), routers.Delete)

	//listen and serve
	router.Run(":4000")

}
