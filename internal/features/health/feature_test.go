package users

import (
	"app/internal/app"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	_, e := app.NewTest(t)
	defer e.Teardown()

	e.GET("/health").Expect().IsOK()
}
