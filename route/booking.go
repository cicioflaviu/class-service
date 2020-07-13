package route

import (
	"cicio.dev/class-service/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func InitializeBookingRoutes(router *gin.RouterGroup) {
	bookings := router.Group("bookings")
	bookings.POST("", CreateBooking)
}

func CreateBooking(context *gin.Context) {
	var booking model.Booking
	err := context.BindJSON(&booking)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = validate.Struct(booking)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if !database.NewRecord(booking) {
		context.JSON(http.StatusBadRequest, "Invalid booking, declaring an ID for a new booking is not supported.")
		return
	}

	if &booking.ClassId == nil {
		context.JSON(http.StatusBadRequest, "Invalid booking, class id not declared.")
		return
	}

	var class model.Class
	var classAssociation = database.First(&class, booking.ClassId)
	if classAssociation.RecordNotFound() {
		context.JSON(http.StatusForbidden, "Unable to create booking. Class with id "+
			strconv.FormatUint(uint64(booking.ClassId), 10)+" not found.")
		return
	}

	isAfterStartDate := booking.Date.After(class.StartDate) || booking.Date.Equal(class.StartDate)
	isBeforeEndDate := booking.Date.Before(class.EndDate) || booking.Date.Equal(class.EndDate)
	if !isAfterStartDate || !isBeforeEndDate {
		context.JSON(http.StatusForbidden, "Unable to create booking. The specified date is outside of the class interval.")
		return
	}

	var dayClass model.DayClass
	var dayClassAssociation = database.Where(model.DayClass{ClassId: class.ID, Date: booking.Date.Truncate(24 * time.Hour)}).First(&dayClass)
	if dayClassAssociation.RecordNotFound() {
		dayClass = model.DayClass{
			Date:          booking.Date.Truncate(24 * time.Hour),
			IsFullyBooked: false,
		}
		database.Create(&dayClass)
		database.Model(&class).Association("DayClasses").Append(&dayClass)
	}

	if dayClass.IsFullyBooked {
		context.JSON(http.StatusForbidden, "The class is full.")
		return
	}

	var allBookings []model.Booking
	database.Model(&dayClass).Related(&allBookings)
	if len(allBookings) == class.Capacity-1 {
		database.Model(&dayClass).Update(model.DayClass{IsFullyBooked: true})
	}
	database.Create(&booking)
	database.Model(&dayClass).Association("Bookings").Append(&booking)

	context.JSON(http.StatusCreated, booking)
}
