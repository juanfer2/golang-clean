package servers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/juanfer2/golang-clean/src/shared/infrastructure/routes"
)

func StartServerApp() {
	app := fiber.New()
	app.Use(cors.New())

	routes.SetupRoutes(app)
	app.Listen(":4000")
}
