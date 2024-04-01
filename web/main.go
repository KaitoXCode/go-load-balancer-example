package main

import (
	"fmt"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

type WebApp struct {
	URL     *url.URL // web service url
	Healthy bool     // if is healthy
	Config  *Config
	App     *fiber.App
}

func (a *WebApp) setupRoutes() {
	a.App.Get("/ping", func(c *fiber.Ctx) error {
		response := fiber.Map{
			"message": "pong",
		}
		return c.Status(fiber.StatusOK).JSON(response)
	})
	return
}

func setupConfigs() *Config {
	config := &Config{}
	if err := config.loadEnvs(); err != nil {
		panic(err)
	}
	if err := config.loadConfig("config.json"); err != nil {
		panic(err)
	}
	return config
}

func main() {
	web := &WebApp{}
	web.Config = setupConfigs()
	web.App = fiber.New()
	web.setupRoutes()
	web.App.Listen(web.Config.Port)
	fmt.Printf("[WEB] [*] Listening on port{%s}...", web.Config.Port)
}
