package main

import (
	"log"

	"github.com/indrariksa/ws-indra2024/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/indrariksa/ws-indra2024/url"

	"github.com/gofiber/fiber/v2"
	_ "github.com/indrariksa/ws-indra2024/docs"
)

// @title TES SWAGGER ULBI
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/indrariksa
// @contact.email indra@ulbi.ac.id

// @host ws-indra2024-878f7e6fab92.herokuapp.com
// @BasePath /
// @schemes https http
func main() {
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
