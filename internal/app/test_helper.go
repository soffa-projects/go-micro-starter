package app

import (
	"github.com/soffa-projects/go-micro/di"
	"github.com/soffa-projects/go-micro/micro"
	"github.com/soffa-projects/go-micro/tests"
	"testing"
)

func NewTest(t *testing.T, features ...micro.Feature) (*micro.App, tests.HttpExpect) {
	micro.Set(micro.DatabaseUrl, "file::memory:?parseTime=true")
	a := Create(features)
	return a, tests.HttpTest(t, a.Env.Router.Handler(), func() {
		a.Cleanup()
		di.Clear()
		micro.Reset()
	})
}
