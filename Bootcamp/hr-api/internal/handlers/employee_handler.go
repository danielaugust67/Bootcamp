package handlers

import (
	"net/http"
	"strconv"

	"github.com/danielaugust67/hr-api/internal/dto"
	"github.com/danielaugust67/hr-api/internal/services"
	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	Service *services.EmployeeService
}

func NewEmployeeHandler(service *services.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{Service: service}
}

func (h *EmployeeHandler) GetAll(c *gin.Context) {
	result, err := h.Service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *EmployeeHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := h.Service.GetByID(c.Request.Context(), int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *EmployeeHandler) Create(c *gin.Context) {
	var req dto.EmployeeDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.Create(c.Request.Context(), &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "created"})
}

func (h *EmployeeHandler) Update(c *gin.Context) {
	var req dto.EmployeeDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.Update(c.Request.Context(), &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (h *EmployeeHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.Delete(c.Request.Context(), int32(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
