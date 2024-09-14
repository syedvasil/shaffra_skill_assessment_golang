package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/syedvasil/shaffra_skill_assessment_golang/config"
	dbmongo "github.com/syedvasil/shaffra_skill_assessment_golang/database/mongo"
	_ "github.com/syedvasil/shaffra_skill_assessment_golang/docs"
	"github.com/syedvasil/shaffra_skill_assessment_golang/internal/middleware"
)

// @title Users API - Shaffra
// @version 1.0
// @description API documentation for the Users API.

// @contact.name   Syed
// @contact.url    https://www.linkedin.com/in/syed-vasil/
// @contact.email  syedvasil@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	// Load configuration
	cfg := config.Config()

	// Initialize MongoDB connection
	//dbConn, err := dbmongo.InitDB(cfg.DB.URI)
	dbConn, err := dbmongo.InitDBConnect(cfg.DB.URI)
	if err != nil {
		fmt.Printf("%+v \n", err)
		log.Fatal(err)
	}

	// Create a new Gin router
	server := gin.Default()

	router := server.RouterGroup

	// Swagger documentation endpoint
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Apply middleware
	router.Use(middleware.LogProcessTime())
	//server.Use(middleware.AuthMiddleware(dbConn))  // skipping as not part of the requirements

	// Define API routes
	v1 := router.Group("/api/v1")

	// Setup API routes groups
	SetupV1UserRoutes(dbConn, v1)

	// Start the server
	err2 := server.Run(":" + strconv.Itoa(int(cfg.App.Port)))
	if err2 != nil {
		return
	}
}
