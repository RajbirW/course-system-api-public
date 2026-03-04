package handler

import (
	"RegisterApplication/internal/entity"
	"RegisterApplication/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SectionHandler struct {
	usecase *usecase.SectionUsecase
}

func NewSectionHandler(u *usecase.SectionUsecase) *SectionHandler {
	return &SectionHandler{usecase: u}
}

func (h *SectionHandler) RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("", h.Create)       // POST /courses/:courseID/sections
	rg.GET("", h.GetByCourseID) // GET /courses/:courseID/sections
}

func (h *SectionHandler) Create(c *gin.Context) {
	if c.GetString("role") != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can add sections"})
		return
	}
	courseID, _ := strconv.Atoi(c.Param("courseID"))

	var section entity.Section
	if err := c.ShouldBindJSON(&section); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	section.CourseID = uint(courseID)

	if err := h.usecase.Create(&section); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, section)
}

func (h *SectionHandler) GetByCourseID(c *gin.Context) {
	courseID, _ := strconv.Atoi(c.Param("courseID"))
	sections, err := h.usecase.GetByCourseID(uint(courseID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sections)
}
