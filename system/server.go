package system

import (
	"cicio.dev/class-service/route"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var SERVER *gin.Engine

func InitServer(database *gorm.DB) *gin.Engine {
	server := gin.Default()
	route.InitRoutes(server, database)

	SERVER = server
	return SERVER
}