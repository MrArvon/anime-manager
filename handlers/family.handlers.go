package handlers

import (
	"anime-manager/models"
	"anime-manager/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type FamilyHandler interface {
	CreateFamily(c *gin.Context)
	GetAllFamily(c *gin.Context)
	GetFamilyByID(c *gin.Context)
	UpdateFamily(c *gin.Context)
	DeleteFamily(c *gin.Context)
}

type familyHandler struct {
	service services.FamilyService
}

func NewFamilyHandler(service services.FamilyService) *familyHandler {
	return &familyHandler{service}
}

// CreateFamily godoc
// @Summary Add New Family
// @Tags Data Family
// @Accept json
// @Produce json
// @Param request body models.Family true "query params"
// @Router /api/family [post]
func (fh *familyHandler) CreateFamily(c *gin.Context) {
	var family models.Family
	err := c.ShouldBindJSON(&family)
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

	res, err := fh.service.CreateFamily(family)
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

// GetAllFamily godoc
// @Summary Call All Data Family
// @Tags Data Family
// @Accept json
// @Produce json
// @Router /api/family [get]
func (fh *familyHandler) GetAllFamily(c *gin.Context) {
	res, err := fh.service.GetAllFamily()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"row":     len(res),
		"status":  http.StatusOK,
		"message": "success",
	})
}

// GetFamilyByID godoc
// @Summary Call Single Data Family
// @Tags Data Family
// @Accept json
// @Produce json
// @Param id path string true "search family by id"
// @Router /api/family/{id} [get]
func (fh *familyHandler) GetFamilyByID(c *gin.Context) {
	id := c.Param("id")
	res, err := fh.service.GetFamilyByID(id)
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

// UpdateFamily godoc
// @Summary Update Single Data Family
// @Tags Data Family
// @Accept json
// @Produce json
// @Param id path string true "update family by id"
// @Param request body models.FamilyRequest false "query params"
// @Router /api/family/{id} [put]
func (fh *familyHandler) UpdateFamily(c *gin.Context) {
	id := c.Param("id")

	var family models.FamilyRequest
	err := c.ShouldBindJSON(&family)
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

	res, err := fh.service.UpdateFamily(id, family)
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

// DeleteFamily godoc
// @Summary Delete Family
// @Tags Data Family
// @Accept json
// @Produce json
// @Param id path string true "delete family"
// @Router /api/family/{id} [delete]
func (fh *familyHandler) DeleteFamily(c *gin.Context) {
	id := c.Param("id")
	err := fh.service.DeleteFamily(id)
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
