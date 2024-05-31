package main

import (
	"github.com/fiatfour/go-hex/adapters"
	"github.com/fiatfour/go-hex/core"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	// Initialize the database connection
	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&core.Order{})

	// Set up the core service and adapters
	orderRepo := adapters.NewGormOrderRepository(db)
	orderService := core.NewOrderService(orderRepo)
	orderHandler := adapters.NewHttpOrderHandler(orderService)

	// Define routes
	app.Post("/order", orderHandler.CreateOrderHandler)

	// Start the server
	app.Listen("localhost:8080")
}
