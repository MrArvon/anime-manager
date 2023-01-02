package routes

import (
	"anime-manager/handlers"
	"github.com/gin-gonic/gin"
)

func AnimeRoutes(incomingRoutes *gin.RouterGroup, handler handlers.AnimeHandler) {
	group := incomingRoutes.Group("api")
	//group.Use(AdminChecker)
	group.GET("/anime", handler.GetAllAnime)
	group.POST("/anime", handler.CreateAnime)
}
