package model

import "time"

type DayClass struct {
	Model

	Date          time.Time `json:"date"`
	IsFullyBooked bool      `json:"isFullyBooked"`
	ClassId       uint      `json:"classId"`
	Class         Class     `json:"-" gorm:"ForeignKey: ClassId"`
	Bookings      []Booking `json:"bookings"  gorm:"foreignkey:DayClassId"`
}
