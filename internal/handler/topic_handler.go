package handler

import (
	"RegisterApplication/internal/entity"
	"RegisterApplication/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TopicHandler struct {
	usecase *usecase.TopicUsecase
}

func NewTopicHandler(u *usecase.TopicUsecase) *TopicHandler {
	return &TopicHandler{usecase: u}
}

func (h *TopicHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/sections/:sectionID/topics", h.Create)
	r.GET("/sections/:sectionID/topics", h.GetBySectionID)
}

func (h *TopicHandler) Create(c *gin.Context) {
	if c.GetString("role") != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can add topics"})
		return
	}
	sectionID, _ := strconv.Atoi(c.Param("sectionID"))

	var topic entity.Topic
	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	topic.SectionID = uint(sectionID)

	if err := h.usecase.Create(&topic); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, topic)
}

func (h *TopicHandler) GetBySectionID(c *gin.Context) {
	sectionID, _ := strconv.Atoi(c.Param("sectionID"))
	topics, err := h.usecase.GetBySectionID(uint(sectionID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, topics)
}
