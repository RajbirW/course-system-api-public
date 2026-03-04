package handler

import (
	"RegisterApplication/internal/repository"
	"RegisterApplication/internal/usecase"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"os"
)

func RegisterHandlers(r *gin.Engine, db *gorm.DB) {
	// Repositories

	os.MkdirAll("uploads", os.ModePerm)

	userRepo := repository.NewUserRepository(db)
	courseRepo := repository.NewCourseRepository(db)
	sectionRepo := repository.NewSectionRepository(db)
	topicRepo := repository.NewTopicRepository(db)
	enrollmentRepo := repository.NewEnrollmentRepository(db)
	documentRepo := repository.NewDocumentRepository(db)

	// Usecases
	authUC := usecase.NewAuthUsecase(userRepo)
	courseUC := usecase.NewCourseUsecase(courseRepo)
	sectionUC := usecase.NewSectionUsecase(sectionRepo)
	topicUC := usecase.NewTopicUsecase(topicRepo)
	enrollmentUC := usecase.NewEnrollmentUsecase(enrollmentRepo)
	documentUC := usecase.NewDocumentUsecase(documentRepo)

	// Handlers
	courseHandler := NewCourseHandler(courseUC)
	sectionHandler := NewSectionHandler(sectionUC)
	topicHandler := NewTopicHandler(topicUC)
	enrollmentHandler := NewEnrollmentHandler(enrollmentUC)
	documentHandler := NewDocumentHandler(documentUC)

	// Public routes
	r.POST("/register", func(c *gin.Context) {
		var input struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Role     string `json:"role"`
		}
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		user, err := authUC.Register(input.Username, input.Password, input.Role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	})

	r.POST("/login", func(c *gin.Context) {
		var input struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		token, err := authUC.Login(input.Username, input.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	// Protected routes with auth middleware
	authGroup := r.Group("/")
	authGroup.Use(authMiddleware(userRepo))
	{
		authGroup.GET("/read", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Authenticated route success"})
		})

		// Courses and nested routes
		coursesGroup := authGroup.Group("/courses")
		{
			coursesGroup.POST("", courseHandler.Create)
			coursesGroup.GET("", courseHandler.GetAll)
			
			// Single course operations and nested resources
			courseGroup := coursesGroup.Group("/:id")
			{
				courseGroup.GET("", courseHandler.GetByID)
				courseGroup.DELETE("", courseHandler.Delete)

				// Sections under this course
				courseGroup.GET("/sections", sectionHandler.GetByCourseID)
				courseGroup.POST("/sections", sectionHandler.Create)

				// Topics under a section in this course
				courseGroup.GET("/sections/:sectionID/topics", topicHandler.GetBySectionID)
				courseGroup.POST("/sections/:sectionID/topics", topicHandler.Create)
			}
		}

		// Enrollments (assuming not nested)
		enrollmentsGroup := authGroup.Group("/enrollments")
		{
			enrollmentsGroup.POST("", enrollmentHandler.Enroll)
			enrollmentsGroup.GET("", enrollmentHandler.GetMyCourses)
		}

		// Inside authGroup:
		documentGroup := authGroup.Group("/documents")
		{
			documentHandler.RegisterRoutes(documentGroup)
		}
	}
}

func authMiddleware(repo *repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		user, err := repo.FindByToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		c.Set("user", map[string]interface{}{
			"ID":   user.ID,
			"Role": user.Role,
		})
		c.Set("role", user.Role)
		fmt.Println("Authenticated user role:", user.Role)
		c.Next()
	}
}