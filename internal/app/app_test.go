package app

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAppHomeHandler(t *testing.T){
	app := New()
	app.Start()

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	app.router.ServeHTTP(rec, req)

	assert.Equal(t, rec.Body.String(), "Домашняя страница")
}