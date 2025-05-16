package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maxonbejenari/testWebApp/config"
	"github.com/maxonbejenari/testWebApp/database"
	"github.com/maxonbejenari/testWebApp/handlers"
	"github.com/maxonbejenari/testWebApp/middleware"
	"log"
)

func main() {
	// load variables form env
	config.LoadEnv()
	// connect to out database
	database.ConnectDB()

	r := gin.Default()
	// we will need cors when we add FrontEnd
	r.Use(middleware.CORSMiddleware())

	api := r.Group("/api")
	{
		api.POST("/login", handlers.Login(database.DB))
		api.POST("/patients", middleware.AuthMiddleware(), handlers.CreatePatient(database.DB))
		api.GET("/patients", middleware.AuthMiddleware(), handlers.ListPatients(database.DB))
		api.GET("/patients/:id", middleware.AuthMiddleware(), handlers.GetPatient(database.DB))
		api.PUT("/patients/:id", middleware.AuthMiddleware(), handlers.UpdatePatient(database.DB))
		api.DELETE("/patients/:id", middleware.AuthMiddleware(), handlers.DeletePatient(database.DB))
	}

	log.Fatal(r.Run(":8080"))
}
