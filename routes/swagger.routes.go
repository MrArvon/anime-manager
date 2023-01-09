package routes

import (
	"anime-manager/docs"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/logrusorgru/aurora/v4"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SwaggerRoutes(engine *gin.Engine) {
	engine.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.BasePath = "/v1"
	fmt.Println("Swagger\t:", aurora.Green("http://localhost:8080/swagger/index.html"))
	fmt.Println("Home\t:", aurora.BrightYellow("http://localhost:8080/v1/api/anime"))
}
