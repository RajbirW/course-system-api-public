package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"RegisterApplication/internal/entity"
	"RegisterApplication/internal/handler"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Clean database - drop all tables (UNCOMMENT IF YOU WANT TO RESET THE DATABASE)
	// err = db.Migrator().DropTable(
	// 	&entity.User{},
	// 	&entity.Document{},
	// 	&entity.Course{},
	// 	&entity.Section{},
	// 	&entity.Topic{},
	// 	&entity.Enrollment{},
	// )
	// if err != nil {
	// 	panic("failed to drop tables: " + err.Error())
	// }

	// Recreate all tables
	err = db.AutoMigrate(
		&entity.User{},
		&entity.Document{},
		&entity.Course{},
		&entity.Section{},
		&entity.Topic{},
		&entity.Enrollment{},
	)
	if err != nil {
		panic("failed to migrate database schema: " + err.Error())
	}

	// Seed initial data
	seedDatabase(db)

	r := gin.Default()
	handler.RegisterHandlers(r, db)
	r.Run(":8080")
}

// Seeding function
func seedDatabase(db *gorm.DB) {
	// Example: Create an admin user
	admin := entity.User{
		Username: "admin",
		Password: "admin123",
		Role:     "admin",
	}
	db.Create(&admin)
}