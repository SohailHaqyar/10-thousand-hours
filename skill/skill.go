package skill

import (
	"fmt"

	"github.com/SohailHaqyar/10-hours/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Skill struct {
	gorm.Model
	Name  string `json:"name"`
	Hours int    `json:"hours"`
}

func SetupRoutes(app *fiber.App) {
	app.Get("/skills", getSkills)
	app.Post("/skills", addSkill)
	app.Put("/skills/hour/:id", incrementHour)
	app.Put("/skills/:id", updateHour)
	app.Delete("/skills/:id", deleteSkill)
}

func getSkills(c *fiber.Ctx) error {
	var skills []Skill
	db := database.DatabaseConfig
	db.Find(&skills)
	return c.JSON(skills)
}

func addSkill(c *fiber.Ctx) error {

	db := database.DatabaseConfig
	skill := new(Skill)

	// parse the body and make sure the type matches the struct otherwise send an error
	if err := c.BodyParser(skill); err != nil {
		c.Status(503).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}
	// Here witht he c.BodyParser we basically put the content of the body into the skill variable

	db.Create(&skill)
	return c.JSON(skill)

}

func incrementHour(c *fiber.Ctx) error {
	db := database.DatabaseConfig
	id := c.Params("id")
	var skill Skill

	db.Find(&skill, id)
	skill.Hours = skill.Hours + 1

	db.Save(&skill)
	return c.JSON(skill)

}

func deleteSkill(c *fiber.Ctx) error {

	db := database.DatabaseConfig
	id := c.Params("id")
	var skill Skill

	db.First(&skill, id)

	if skill.Name == "" {
		c.Status(400).Send([]byte("Skill you're trying to delete not found"))
	}

	db.Delete(&skill, id)
	return c.Send([]byte(fmt.Sprintf("Skill with id: %s deleted successfully", id)))
}

func updateHour(c *fiber.Ctx) error {
	db := database.DatabaseConfig
	id := c.Params("id")

	body := new(Skill)

	if err := c.BodyParser(body); err != nil {
		return err
	}

	var skill Skill

	db.First(&skill, id)

	if skill.Name == "" {
		c.Status(400).Send([]byte("Skill you're trying to update not found"))
	}

	if body.Name != "" {
		skill.Name = body.Name
	}
	if body.Hours != 0 {
		skill.Hours = body.Hours
	}

	db.Save(&skill)
	return c.JSON(fiber.Map{
		"message": "Skill updated successfully",
		"id":      id,
		"skill":   skill,
	})

}
