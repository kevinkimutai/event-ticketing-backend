package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type authHandlerPort interface {
	IsAuthenticated(*fiber.Ctx) error
	//AllowedRoles(admin string) func(c *fiber.Ctx) error
}

type eventHandlerPort interface {
	CreateEvent(*fiber.Ctx) error
}

type ServerAdapter struct {
	port  string
	auth  authHandlerPort
	event eventHandlerPort
}

func New(
	port string,
	auth authHandlerPort,
	event eventHandlerPort) *ServerAdapter {
	return &ServerAdapter{
		port:  port,
		auth:  auth,
		event: event,
	}
}

func (s *ServerAdapter) StartServer() {
	//Initialize Fiber
	app := fiber.New()

	//Logger Middleware
	app.Use(logger.New())

	//Auth Middleware
	//Must be Authenticated
	app.Use(s.auth.IsAuthenticated)

	// Define routes
	app.Route("/api/v1/event", s.EventRouter)

	app.Listen(":" + s.port)

}
