package handler

import (
	"RegisterApplication/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EnrollmentHandler struct {
	usecase *usecase.EnrollmentUsecase
}

func NewEnrollmentHandler(u *usecase.EnrollmentUsecase) *EnrollmentHandler {
	return &EnrollmentHandler{usecase: u}
}

func (h *EnrollmentHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/courses/:id/enroll", h.Enroll)
	r.GET("/my-courses", h.GetMyCourses)
}

func (h *EnrollmentHandler) Enroll(c *gin.Context) {
	if c.GetString("role") != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can enroll users into courses"})
		return
	}

	var input struct {
		UserID uint `json:"user_id"`
	}
	
	if err := c.BindJSON(&input); err != nil || input.UserID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	courseID, _ := strconv.Atoi(c.Param("id"))
	if err := h.usecase.EnrollUser(input.UserID, uint(courseID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User enrolled successfully"})
}

func (h *EnrollmentHandler) GetMyCourses(c *gin.Context) {
	user := c.MustGet("user")
	userID := user.(map[string]interface{})["ID"].(uint)

	courses, err := h.usecase.GetUserCourses(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, courses)
}
