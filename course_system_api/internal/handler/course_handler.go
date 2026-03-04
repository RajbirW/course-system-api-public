package handler

import (
	"RegisterApplication/internal/entity"
	"RegisterApplication/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	usecase *usecase.CourseUsecase
}

func NewCourseHandler(u *usecase.CourseUsecase) *CourseHandler {
	return &CourseHandler{
		usecase: u,
	}
}

func (h *CourseHandler) RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("", h.Create)       // POST /courses
	rg.GET("", h.GetAll)        // GET /courses
	rg.GET("/:id", h.GetByID)   // GET /courses/:id
	rg.DELETE("/:id", h.Delete) // DELETE /courses/:id
}

func (h *CourseHandler) Create(c *gin.Context) {
	if c.GetString("role") != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can create courses"})
		return
	}

	var course entity.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.usecase.Create(&course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, course)
}

func (h *CourseHandler) GetAll(c *gin.Context) {
	courses, err := h.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, courses)
}

func (h *CourseHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	course, err := h.usecase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}
	c.JSON(http.StatusOK, course)
}

func (h *CourseHandler) Delete(c *gin.Context) {
	if c.GetString("role") != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can delete courses"})
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.usecase.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Course deleted"})
}
