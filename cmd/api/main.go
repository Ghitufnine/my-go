// Package main is the entry point of the application.
//
//	@title			My Go API
//	@version		1.0
//	@description	REST API built with Go, Fiber, PostgreSQL, Redis, MongoDB, and RabbitMQ.
//	@termsOfService	http://swagger.io/terms/
//
//	@contact.name	API Support
//	@contact.email	support@example.com
//
//	@license.name	MIT
//	@license.url	https://opensource.org/licenses/MIT
//
//	@host		localhost:8080
//	@BasePath	/api
//
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and the JWT token.
package main

import (
	_ "github.com/ghitufnine/my-go/docs"

	"github.com/ghitufnine/my-go/cmd/container"
)

func main() {
	container.Container()
}
