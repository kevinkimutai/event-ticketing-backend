package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type authHandlerPort interface {
	IsAuthenticated(*fiber.Ctx) error
	AllowedRoles(allowedrole string) func(*fiber.Ctx) error
}

type eventHandlerPort interface {
	CreateEvent(*fiber.Ctx) error
	GetEvents(*fiber.Ctx) error
	GetEvent(*fiber.Ctx) error
	UpdateEvent(*fiber.Ctx) error
}

type categoryHandlerPort interface {
	CreateCategory(*fiber.Ctx) error
	GetCategories(*fiber.Ctx) error
}
type ticketTypeHandlerPort interface {
	CreateTicketType(*fiber.Ctx) error
	GetTicketTypesByEvent(c *fiber.Ctx) error
	GetTicketByTicketTypeID(c *fiber.Ctx) error
}

type ticketHandlerPort interface {
	CreateTicketOrder(*fiber.Ctx) error
	GetOrder(c *fiber.Ctx) error
	GetTicketOrderItem(*fiber.Ctx) error
	// GetTicketsByEvent(*fiber.Ctx) error
}
type LocationHandlerPort interface {
	GetLocationByID(*fiber.Ctx) error
	GetLocations(c *fiber.Ctx) error
}

type UserHandlerPort interface {
	GetUser(c *fiber.Ctx) error
}
type AttendeeHandlerPort interface {
	GetAttendee(c *fiber.Ctx) error
	GetAttendeeEvents(c *fiber.Ctx) error
}

type OrganiserHandlerPort interface {
	GetOrganiserByUserID(c *fiber.Ctx) error
	GetOrganiserEvent(c *fiber.Ctx) error
	DownloadOrganiserEvent(c *fiber.Ctx) error
}

type UsherHandlerPort interface {
	GetTicketOrderDetails(c *fiber.Ctx) error
	AdmitTicketOrder(c *fiber.Ctx) error
}

type ServerAdapter struct {
	port       string
	auth       authHandlerPort
	event      eventHandlerPort
	category   categoryHandlerPort
	tickettype ticketTypeHandlerPort
	ticket     ticketHandlerPort
	location   LocationHandlerPort
	user       UserHandlerPort
	attendee   AttendeeHandlerPort
	organiser  OrganiserHandlerPort
	usher      UsherHandlerPort
}

func New(
	port string,
	auth authHandlerPort,
	event eventHandlerPort,
	category categoryHandlerPort,
	tickettype ticketTypeHandlerPort,
	ticket ticketHandlerPort,
	location LocationHandlerPort,
	user UserHandlerPort,
	attendee AttendeeHandlerPort,
	organiser OrganiserHandlerPort,
	usher UsherHandlerPort,

) *ServerAdapter {
	return &ServerAdapter{
		port:       port,
		auth:       auth,
		event:      event,
		category:   category,
		tickettype: tickettype,
		ticket:     ticket,
		location:   location,
		user:       user,
		attendee:   attendee,
		organiser:  organiser,
		usher:      usher,
	}
}

func (s *ServerAdapter) StartServer() {

	//Initialize Fiber
	app := fiber.New()

	//Serve Static Files
	app.Static("/", "./public")

	//Logger Middleware
	app.Use(logger.New())

	//Telemetry Observability Middleware
	//app.Use(s.telemetry.OtelFiberMiddleware)

	// Cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000,https://www.ticketpass.site",
		AllowMethods: "GET,POST,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Authorization, Accept",
	}))

	//Auth Middleware
	//Must be Authenticated
	app.Use(s.auth.IsAuthenticated)

	// Define routes
	app.Route("/api/v1/event", s.EventRouter)
	app.Route("/api/v1/category", s.CategoryRouter)
	app.Route("/api/v1/ticket-type", s.TicketTypeRouter)
	app.Route("/api/v1/ticket-order", s.TicketRouter)
	app.Route("/api/v1/location", s.LocationRouter)
	app.Route("/api/v1/user", s.UserRouter)
	app.Route("/api/v1/attendee", s.AttendeeRouter)
	app.Route("/api/v1/organiser", s.OrganiserRouter)
	app.Route("/api/v1/usher", s.UsherRouter)

	app.Listen(":" + s.port)

}
