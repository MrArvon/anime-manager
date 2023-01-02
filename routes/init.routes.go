package routes

import (
	"anime-manager/handlers"
	"anime-manager/repositories"
	"anime-manager/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(router *gin.Engine, db *gorm.DB) {
	v := router.Group("/v1")

	animeRepository := repositories.NewAnimeRepository(db)
	animeService := services.NewAnimeService(animeRepository)
	animeHandler := handlers.NewAnimeHandler(animeService)

	familyRepository := repositories.NewFamilyRepository(db)
	familyService := services.NewFamilyService(familyRepository)
	familyHandler := handlers.NewFamilyHandler(familyService)

	AnimeRoutes(v, animeHandler)
	FamilyRoutes(v, familyHandler)
	SwaggerRoutes(router)
}
