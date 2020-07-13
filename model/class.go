package model

import "time"

type Class struct {
	Model

	Name       string     `json:"name" validate:"required"`
	StartDate  time.Time  `json:"startDate"  validate:"required"`
	EndDate    time.Time  `json:"endDate" validate:"required"`
	Capacity   int        `json:"capacity" validate:"required,min=0"`
	DayClasses []DayClass `json:"-"`
}
