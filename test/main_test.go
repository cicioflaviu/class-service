package test

import (
	"cicio.dev/class-service/system"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	mocket "github.com/selvatico/go-mocket"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var router *gin.Engine
var db *gorm.DB

func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run()
	tearDown()

	os.Exit(retCode)
}

func setUp() {
	db = mockDatabase()
	router = system.InitServer(db)
}

func tearDown() {
	_ = db.Close()
}

func performRequest(r http.Handler, method, path string, body string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, ioutil.NopCloser(strings.NewReader(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func mockDatabase() *gorm.DB {
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true
	db, _ := gorm.Open(mocket.DriverName, "connection_string") // Can be any connection string
	return db
}
