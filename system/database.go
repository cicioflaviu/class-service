package system

import (
	"cicio.dev/class-service/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func InitDatabase() *gorm.DB {
	database, err := gorm.Open("sqlite3", "file::memory:?cache=shared")

	if err != nil {
		fmt.Println("Database error: ", err)
	}
	database.DB().SetMaxIdleConns(10)
	database.AutoMigrate(&model.Class{}, &model.Booking{}, &model.DayClass{})

	DB = database
	return DB
}
