package route

import (
	"cicio.dev/class-service/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

var database *gorm.DB

func InitializeClassRoutes(router *gin.RouterGroup) {
	classes := router.Group("classes")
	classes.POST("", CreateClass)
}

func CreateClass(context *gin.Context) {
	var class model.Class
	err := context.BindJSON(&class)

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	err = validate.Struct(class)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if !database.NewRecord(class) {
		context.JSON(http.StatusBadRequest, "Invalid class, declaring an ID for a new class is not supported.")
		return
	}

	database.Create(&class)
	context.JSON(http.StatusCreated, class)
}