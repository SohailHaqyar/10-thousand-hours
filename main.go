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
	URI := "postgres://pufttqfdektzqa:e3d39ca00bf8794f743f9d572a3e35762146b9a698080af6f4e6e4ecf663f08c@ec2-44-193-111-218.compute-1.amazonaws.com:5432/dfsrp72rdk52su"

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
