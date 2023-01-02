package routes

import (
	"anime-manager/handlers"
	"github.com/gin-gonic/gin"
)

//v1 := router.Group("/v1")

func FamilyRoutes(incomingRoutes *gin.RouterGroup, handler handlers.FamilyHandler) {
	group := incomingRoutes.Group("api")
	//group.Use(AdminChecker)
	group.GET("/family", handler.GetAllFamily)
	group.GET("/family/:id", handler.GetFamilyByID)
	group.POST("/family", handler.CreateFamily)
	group.DELETE("/family/:id", handler.DeleteFamily)
}
