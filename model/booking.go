package model

import "time"

type Booking struct {
	Model

	Name string `json:"name" validate:"required"`
	Date time.Time `json:"date" validate:"required"`
	ClassId uint `json:"classId" gorm:"-" validate:"required"`
	DayClassId uint `json:"-"`
}