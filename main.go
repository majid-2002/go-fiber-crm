package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/majid-2002/go-fiber-crm/database"
	"github.com/majid-2002/go-fiber-crm/lead"
	"os"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	passDb := os.Getenv("DB_PASS")

	dsn := "abdul_majid:" + passDb + "@tcp(localhost:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"

	database.DBConn, err = gorm.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connection to database was successful")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3003)
	fmt.Println("Server is running on port 3003")
	defer database.DBConn.Close()
}
