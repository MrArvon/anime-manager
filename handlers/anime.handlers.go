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
	GetAnimeByID(c *gin.Context)
	UpdateAnime(c *gin.Context)
	DeleteAnime(c *gin.Context)
}

type animeHandler struct {
	service services.AnimeService
}

func NewAnimeHandler(service services.AnimeService) *animeHandler {
	return &animeHandler{service}
}

// CreateAnime godoc
// @Summary Add New Anime
// @Tags Data Anime
// @Accept json
// @Produce json
// @Param request body models.Anime true "query params"
// @Router /api/anime [post]
func (ah *animeHandler) CreateAnime(c *gin.Context) {
	var anime models.Anime
	err := c.ShouldBindJSON(&anime)
	if err != nil {
		var errMsgs []string
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

	res, err := ah.service.CreateAnime(anime)
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

// GetAllAnime godoc
// @Summary Call All Data Anime
// @Tags Data Anime
// @Accept json
// @Produce json
// @Router /api/anime [get]
func (ah *animeHandler) GetAllAnime(c *gin.Context) {
	res, err := ah.service.GetAllAnime()
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

// GetAnimeByID godoc
// @Summary Call Single Data Anime
// @Tags Data Anime
// @Accept json
// @Produce json
// @Param id path string true "search anime by id"
// @Router /api/anime/{id} [get]
func (ah *animeHandler) GetAnimeByID(c *gin.Context) {
	id := c.Param("id")
	res, err := ah.service.GetAnimeByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"status":  http.StatusOK,
		"message": "success",
	})
}

// UpdateAnime godoc
// @Summary Update Single Data Anime
// @Tags Data Anime
// @Accept json
// @Produce json
// @Param id path string true "update anime by id"
// @Param request body models.AnimeRequest false "query params"
// @Router /api/anime/{id} [put]
func (ah *animeHandler) UpdateAnime(c *gin.Context) {
	id := c.Param("id")

	var anime models.AnimeRequest
	err := c.ShouldBindJSON(&anime)
	if err != nil {
		var errMsgs []string
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

	res, err := ah.service.UpdateAnime(id, anime)
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

// DeleteAnime godoc
// @Summary Delete Anime
// @Tags Data Anime
// @Accept json
// @Produce json
// @Param id path string true "delete anime"
// @Router /api/anime/{id} [delete]
func (ah *animeHandler) DeleteAnime(c *gin.Context) {
	id := c.Param("id")
	err := ah.service.DeleteAnime(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "success",
		})
	}
}
