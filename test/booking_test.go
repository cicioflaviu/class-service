package test

import (
	"cicio.dev/class-service/model"
	"encoding/json"
	gomocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestCreateBookingSuccessful(t *testing.T) {
	commonReply := []map[string]interface{}{{
		"id":         1,
		"name":       "Pilates",
		"start_date": time.Date(2020, 6, 1, 8, 0, 0, 0, time.UTC),
		"end_date":   time.Date(2020, 6, 5, 8, 0, 0, 0, time.UTC),
		"capacity":   "2"}}
	gomocket.Catcher.Reset().NewMock().WithQuery(`SELECT * FROM "classes"  WHERE`).WithReply(commonReply)

	bookingPostBody := `
	{
		"name":"Adam",
		"date":"2020-06-02T08:00:00Z",
		"classId":1
	}`

	w := performRequest(router, "POST", "/bookings", bookingPostBody)
	assert.Equal(t, http.StatusCreated, w.Code)

	var bookingResponse model.Booking
	err := json.Unmarshal([]byte(w.Body.String()), &bookingResponse)
	assert.Nil(t, err)

	var bookingRequest model.Booking
	err = json.Unmarshal([]byte(bookingPostBody), &bookingRequest)
	assert.Nil(t, err)

	assert.Equal(t, bookingRequest.Date, bookingResponse.Date)
	assert.Equal(t, bookingRequest.ClassId, bookingResponse.ClassId)
	assert.Equal(t, bookingRequest.Name, bookingResponse.Name)
	assert.NotNil(t, bookingResponse.ID)
	assert.Less(t, uint(0), bookingResponse.Model.ID)
}

func TestCreateBookingOutsideDateFail(t *testing.T) {
	commonReply := []map[string]interface{}{{
		"id":         1,
		"name":       "Pilates",
		"start_date": time.Date(2020, 6, 1, 8, 0, 0, 0, time.UTC),
		"end_date":   time.Date(2020, 6, 5, 8, 0, 0, 0, time.UTC),
		"capacity":   "2"}}
	gomocket.Catcher.Reset().NewMock().WithQuery(`SELECT * FROM "classes"  WHERE`).WithReply(commonReply)

	bookingPostBody := `
	{
		"name":"Adam",
		"date":"2020-07-02T08:00:00Z",
		"classId":1
	}`

	w := performRequest(router, "POST", "/bookings", bookingPostBody)
	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.Equal(t, "\"Unable to create booking. The specified date is outside of the class interval.\"", w.Body.String())
}

func TestCreateBookingFullClass(t *testing.T) {
	classMockReply := []map[string]interface{}{{
		"id":         1,
		"name":       "Pilates",
		"start_date": time.Date(2020, 6, 1, 8, 0, 0, 0, time.UTC),
		"end_date":   time.Date(2020, 6, 5, 8, 0, 0, 0, time.UTC),
		"capacity":   "1"}}
	gomocket.Catcher.Reset().NewMock().WithQuery(`SELECT * FROM "classes"  WHERE`).WithReply(classMockReply)

	dayClassMockReply := []map[string]interface{}{{
		"id":              1,
		"date":            time.Date(2020, 6, 2, 0, 0, 0, 0, time.UTC),
		"is_fully_booked": true,
		"class_id":        1}}
	gomocket.Catcher.NewMock().WithQuery(`SELECT * FROM "day_classes"  WHERE`).WithReply(dayClassMockReply)

	bookingPostBody := `
	{
		"name":"Adam",
		"date":"2020-06-02T08:00:00Z",
		"classId":1
	}`

	w := performRequest(router, "POST", "/bookings", bookingPostBody)
	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.Equal(t, "\"The class is full.\"", w.Body.String())
}
