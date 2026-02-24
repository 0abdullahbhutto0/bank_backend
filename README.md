# Bank Ledger API 🏦

A robust backend service for a banking ledger, built with Go and PostgreSQL. 

Currently in active development.

## Tech Stack
* **Language:** Go 1.26
* **Database:** PostgreSQL (via Docker)
* **DB Driver & Queries:** `pgx/v5` + `sqlc`
* **Configuration:** Viper
* **CI/CD:** GitHub Actions

---

## Local Development Setup

### 1. Prerequisites
* [Go](https://go.dev/dl/) installed.
* [Docker](https://www.docker.com/) installed and running.
* `Make` installed (optional, but recommended for running commands).

### 2. Configuration
This project uses Viper for environment configuration. You must set up your local environment file before running the server or tests.

1. Copy the example config file:
   cp app.env.example app.env
Open app.env and update the DB_SOURCE.

3. Database Setup
Start a local Postgres instance using Docker:

# Start the container
docker run --name postgres18 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:18-alpine

# Create the database
docker exec -it postgres18 createdb --username=root --owner=root simple_bank

4. Running the Code
To run the unit tests:

go test -v ./...
To start the development server:

go run main.go