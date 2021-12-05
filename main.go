package main

import (
	"fmt"

	"github.com/SohailHaqyar/10-hours/database"
	"github.com/SohailHaqyar/10-hours/skill"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error
	URI := "postgres://xntvoppnmowgqo:acfd30f057c1f17f9126449c73147a13c5b514686a18dbfd7104259b2eef4576@ec2-34-242-89-204.eu-west-1.compute.amazonaws.com:5432/dcljsqcnian6k5"
	database.DatabaseConfig, err = gorm.Open(postgres.Open(URI), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Error while connecting to the database")
	}
	fmt.Println("Connected to the database successfully")
	database.DatabaseConfig.AutoMigrate(skill.Skill{})
}

func setupRoutes(app *fiber.App) {
	skill.SetupRoutes(app)
}
func main() {
	app := fiber.New()

	app.Use(cors.New())

	initDatabase()
	setupRoutes(app)

	app.Listen(":4000")
	sqldb, _ := database.DatabaseConfig.DB()
	defer sqldb.Close()

}
