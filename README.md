# Ticketing Golang Backend

Backend for an event-ticketing management system written in Go.

## Tech and Tools Used

- [Go-Fiber](https://docs.gofiber.io/).(Express-like web framework)
- [Auth0](https://auth0.com/docs/quickstart/backend).(Authorization & Authentication)
- [Golang-Migrate](https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md)(golang migration tool)
- [Postgresql](https://hub.docker.com/_/postgres).(SQL database)
- [sqlc](https://sqlc.dev/).(Generate SQL type-safe Go code.)
- [fpdf](https://github.com/jung-kurt/gofpdf).(Generate PDF files in Go.)
- [go-qrcode](https://github.com/skip2/go-qrcode).(Generate qr-codes in Go)

## Architecture

Hexagonal Architecture

## Features

- Authentication and Authorization using Auth0
- CRUD events
- CRUD tickets
- Pagination
- Query Parameters
- Generate PDF tickets
- Generate QR codes in tickets to allow for admittance and security

## Getting Started

- Setup your Auth0 at [Auth0](https://auth0.com/docs/quickstart/backend)
- Run make postgresdb :setup database as docker container
- Run make createdb (create db)
- Run make migrateup (migrate database migrations to get latest db schema)
- Use your **own .env file**
- Ensure you have the following environmental variables
- **APP_PORT**
- **POSTGRES_USERNAME**
- **POSTGRES_PASSWORD**
- **DB_PORT**
- **DB_HOST**
- **DB_NAME**
- **AUTH0_URL**
- **AUTH0_CLIENTID**
- **AUTH0_CLIENT_SECRET**

- Run **go run cmd/main.go**

## Database

- Schema
  ![db-schema](https://firebasestorage.googleapis.com/v0/b/creadable-22c39.appspot.com/o/Screenshot%20from%202024-07-07%2003-21-42.png?alt=media&token=ddbb8fda-ab27-4b28-8b7f-9ee001293b64)

## Sample Generated Ticket PDF

- Find **ticket** at [sample](https://github.com/kevinkimutai/event-ticketing-backend/blob/main/summertides-ticket-7.pdf)
  ![ticket-pdf](https://firebasestorage.googleapis.com/v0/b/creadable-22c39.appspot.com/o/Screenshot%20from%202024-07-11%2020-01-35.png?alt=media&token=c01267c4-c895-4fe8-b2b2-21535799adaf)

## Whats Remaining

- Intergration with email service (sendgrid) for sending of notifications
- Payment service (stripe etc)
