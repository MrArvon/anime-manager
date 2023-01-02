package handlers

import (
	"anime-manager/models"
	"anime-manager/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type AnimeHandler interface {
	CreateAnime(c *gin.Context)
	GetAllAnime(c *gin.Context)
}

type animeHandler struct {
	service services.AnimeService
}

func NewAnimeHandler(service services.AnimeService) *animeHandler {
	return &animeHandler{service}
}

// GetAllAnime godoc
// @Summary Call All Data Anime
// @Tags Data Anime
// @Accept json
// @Produce json
// @Router /api/anime [get]
func (h *animeHandler) GetAllAnime(c *gin.Context) {
	res, err := h.service.GetAllAnime()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"row":     len(res),
		"status":  http.StatusOK,
		"message": "success",
	})
}

// CreateAnime godoc
// @Summary Add New Anime
// @Tags Data Anime
// @Accept json
// @Produce json
// @Param request body models.Anime true "query params"
// @Router /api/anime [post]
func (h *animeHandler) CreateAnime(c *gin.Context) {
	var anime models.Anime
	err := c.ShouldBindJSON(&anime)
	if err != nil {
		errMsgs := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errMsg := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errMsgs = append(errMsgs, errMsg)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  errMsgs,
			"status": http.StatusBadRequest,
		})
		return
	}

	res, err := h.service.CreateAnime(anime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"status":  http.StatusOK,
		"message": "success",
	})
}
