package handler

import (
	"RegisterApplication/internal/entity"
	"RegisterApplication/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
)

type DocumentHandler struct {
	usecase *usecase.DocumentUsecase
}

func NewDocumentHandler(u *usecase.DocumentUsecase) *DocumentHandler {
	return &DocumentHandler{usecase: u}
}

func (h *DocumentHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/upload", h.Upload)
	r.GET("", h.GetAll)
	r.GET("/:id", h.GetByID)
	r.DELETE("/:id", h.Delete)
}

func (h *DocumentHandler) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	savePath := filepath.Join("uploads", file.Filename)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	user := c.MustGet("user").(map[string]interface{}) // getting user from middleware

	doc := entity.Document{
		Filename: file.Filename,
		Path:     savePath,
		UserID:   user["ID"].(uint),
	}

	if err := h.usecase.Upload(&doc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save metadata"})
		return
	}

	c.JSON(http.StatusCreated, doc)
}

func (h *DocumentHandler) GetAll(c *gin.Context) {
	docs, err := h.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch documents"})
		return
	}
	c.JSON(http.StatusOK, docs)
}

func (h *DocumentHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	doc, err := h.usecase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
		return
	}
	c.JSON(http.StatusOK, doc)
}

func (h *DocumentHandler) Delete(c *gin.Context) {
	if c.GetString("role") != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can delete documents"})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.usecase.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete document"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Document deleted"})
}
