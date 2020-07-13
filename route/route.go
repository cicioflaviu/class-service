package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

var validate *validator.Validate

func InitRoutes(server *gin.Engine, db *gorm.DB) {
	baseRoute := server.Group("/")
	database = db

	InitializeClassRoutes(baseRoute)
	InitializeBookingRoutes(baseRoute)

	validate = validator.New()
}

