package main

import (
	_ "app/apidocs"
	"app/internal/app"
	"app/internal/features"
)

// @title Service API
// @version 1.0
// @description Documentation for our API
// @termsOfService http://foo.bar/api

// @contact.name API Support
// @contact.url http://foo.bar/support
// @contact.email support@domain.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description "Type 'Bearer TOKEN' to correctly set the API Key"

// @BasePath /v1
func main() {
	srv := app.Create(features.All)
	srv.Run()
}
