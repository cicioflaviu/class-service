package test

import (
	"cicio.dev/class-service/model"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCreateClassSuccessful(t *testing.T) {
	postBody := `
	{
		"name":"Pilates",
		"startDate":"2020-06-01T08:00:00Z",
		"endDate":"2020-06-05T06:00:00Z",
		"capacity":10
	}`

	w := performRequest(router, "POST", "/classes", postBody)

	assert.Equal(t, http.StatusCreated, w.Code)

	var responseClass model.Class
	err := json.Unmarshal([]byte(w.Body.String()), &responseClass)

	var requestClass model.Class
	err = json.Unmarshal([]byte(postBody), &requestClass)

	assert.Nil(t, err)
	assert.Equal(t, requestClass.StartDate, responseClass.StartDate)
	assert.Equal(t, requestClass.EndDate, responseClass.EndDate)
	assert.Equal(t, requestClass.Capacity, responseClass.Capacity)
	assert.Equal(t, requestClass.Name, responseClass.Name)
	assert.NotNil(t, responseClass.ID)
	assert.Less(t, uint(0), responseClass.Model.ID)
}

func TestCreateClassWithNoNameFail(t *testing.T) {
	postBody := `
	{
		"name":"",
		"startDate":"2020-06-01T08:00:00Z",
		"endDate":"2020-06-05T06:00:00Z",
		"capacity":10
	}`

	w := performRequest(router, "POST", "/classes", postBody)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateClassWithNegativeCapacityFail(t *testing.T) {
	postBody := `
	{
		"name":"Pilates",
		"startDate":"2020-06-01T08:00:00Z",
		"endDate":"2020-06-05T06:00:00Z",
		"capacity":-1
	}`

	w := performRequest(router, "POST", "/classes", postBody)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}