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
	GetEvents(*fiber.Ctx) error
	GetEvent(*fiber.Ctx) error
}

type categoryHandlerPort interface {
	CreateCategory(*fiber.Ctx) error
	GetCategories(*fiber.Ctx) error
}
type ticketTypeHandlerPort interface {
	CreateTicketType(*fiber.Ctx) error
	GetTicketTypesByEvent(c *fiber.Ctx) error
}

type ticketHandlerPort interface {
	CreateTicketOrder(*fiber.Ctx) error
}
type ServerAdapter struct {
	port       string
	auth       authHandlerPort
	event      eventHandlerPort
	category   categoryHandlerPort
	tickettype ticketTypeHandlerPort
	ticket     ticketHandlerPort
}

func New(
	port string,
	auth authHandlerPort,
	event eventHandlerPort,
	category categoryHandlerPort,
	tickettype ticketTypeHandlerPort,
	ticket ticketHandlerPort,
) *ServerAdapter {
	return &ServerAdapter{
		port:       port,
		auth:       auth,
		event:      event,
		category:   category,
		tickettype: tickettype,
		ticket:     ticket,
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
	app.Route("/api/v1/category", s.CategoryRouter)
	app.Route("/api/v1/ticket-type", s.TicketTypeRouter)
	app.Route("/api/v1/ticket-order", s.TicketRouter)

	app.Listen(":" + s.port)

}
