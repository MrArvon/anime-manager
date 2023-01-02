package main

import (
	"anime-manager/config"
	"anime-manager/models"
	"anime-manager/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router *gin.Engine
	port   string
	//db     *gorm.DB
)

func init() {
	port = config.GetEnv("PORT")
	if port == "" {
		fmt.Println("<=== Using Default Port ===>")
		port = "8080"
	}

	router = gin.Default()
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatalln(err)
	}

	//db = models.InitDB()
	routes.InitRoutes(router, models.InitDB())
}

func main() {
	errRun := router.Run(":" + port)
	if errRun != nil {
		log.Fatal(errRun)
	}
}
