package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kevinkimutai/ticketingapp/internal/adapters/auth"
	"github.com/kevinkimutai/ticketingapp/internal/adapters/db"
	handler "github.com/kevinkimutai/ticketingapp/internal/adapters/handlers"
	"github.com/kevinkimutai/ticketingapp/internal/adapters/pdf"
	"github.com/kevinkimutai/ticketingapp/internal/adapters/rabbitmq"
	"github.com/kevinkimutai/ticketingapp/internal/adapters/server"
	application "github.com/kevinkimutai/ticketingapp/internal/app/api"
)

func main() {
	// Init Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
	}

	//Env Variables
	APP_PORT := os.Getenv("APP_PORT")
	RABBITMQ_SERVER_URL := os.Getenv("RABBITMQ_SERVER")
	POSTGRES_USERNAME := os.Getenv("POSTGRES_USERNAME")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	DATABASE_HOST := os.Getenv("DB_HOST")
	DATABASE_PORT := os.Getenv("DB_PORT")
	DATABASE_NAME := os.Getenv("DB_NAME")

	DBURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		POSTGRES_USERNAME,
		POSTGRES_PASSWORD,
		DATABASE_HOST,
		DATABASE_PORT,
		DATABASE_NAME)

	//RABBITMQSERVER := os.Getenv("RABBITMQ_SERVER")

	//Dependency injection

	//Database
	//Connect To DB
	dbAdapter := db.NewDB(DBURL)

	//RabbitMQ
	//msgQueue := queue.NewRabbitMQServer(RABBITMQSERVER)

	//Repositories
	pdfService := pdf.NewPDF()
	queue := rabbitmq.NewRabbitMQServer(RABBITMQ_SERVER_URL)

	eventRepo := application.NewEventRepo(dbAdapter)
	categoriesRepo := application.NewCategoriesRepo(dbAdapter)
	ticketTypeRepo := application.NewTicketTypeRepo(dbAdapter)
	ticketOrderRepo := application.NewTicketOrderRepo(dbAdapter, pdfService, queue)

	//Services
	//telemetryService := telemetry.NewTelemetryService()

	eventService := handler.NewEventService(eventRepo)
	categoryService := handler.NewCategoryService(categoriesRepo)
	ticketTypeService := handler.NewTicketTypeService(ticketTypeRepo)
	ticketOrderService := handler.NewTicketOrderService(ticketOrderRepo)

	authService, err := auth.New(dbAdapter, queue)
	if err != nil {
		log.Fatal(err)
	}

	//Server
	server := server.New(
		APP_PORT,
		authService,
		eventService,
		categoryService,
		ticketTypeService,
		ticketOrderService,
	)

	server.StartServer()

}
