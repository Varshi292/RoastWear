package main

import (
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func main() {
	app := fiber.New()
	app.Static("/", "./frontend/build")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})
	db, _ := openDB("./backend/db/users.db")
	mod := &models.UserModel{DB: db}
	//initializeDB(mod.DB)

	mod.ShowTable("users")

	log.Fatal(app.Listen(":7777"))
}

func openDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func initializeDB(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}

func createTestUsers(mod *models.UserModel) {
	mod.CreateUser("luffy", "luuffy@ao.lcom", "maetrt")
	mod.CreateUser("red-hair", "hakiman@cd.gov", "worldSstrongestSwordsman1997")
	mod.CreateUser("buggy_the_star_clown", "buggy@crossguild.net", "buggyB@115")
	mod.CreateUser("sir_crocodile", "sir_crocodile@crossguild.net", "utopia")
	mod.CreateUser("dracule_mihawk", "draculemihawk@crossguild.net", "9N[2:Jn5-Tr+!5UFKXql,7U,k2A+738a7w2kom£<qETt^%\"DHd")
}
