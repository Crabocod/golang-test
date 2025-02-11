package handler

import (
	"net/http"

	"github.com/Crabocod/golang-test/internal/model"
	"github.com/Crabocod/golang-test/internal/service"
	"github.com/gin-gonic/gin"
)

type HashHandler interface {
	CreateHash(c *gin.Context)
}

type hashHandler struct {
	service service.HashService
}

func NewHashHandler(service service.HashService) HashHandler {
	return &hashHandler{service: service}
}

func (h *hashHandler) CreateHash(c *gin.Context) {
	var request model.HashRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.service.CreateHash(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
