package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Config struct {
	RootDir string
	Port    string
}

type WebPool struct {
	WebURLs []string `json:"webUrls"`
	Current *string  `json:"-"`
}

func (p *WebPool) setNextIdx() error {
	for idx, url := range p.WebURLs {
		if *p.Current == url {
			idx += 1
			if idx > (len(p.WebURLs) - 1) {
				p.Current = &p.WebURLs[0]
				return nil
			} else {
				p.Current = &p.WebURLs[idx]
				return nil
			}
		}
	}
	return errors.New("Failed to find next index!")
}

type LoadBalancer struct {
	Pool    *WebPool
	Configs *Config
	App     *fiber.App
}

func (lb *LoadBalancer) setupConfigs() error {
	configs := Config{}
	root, ok := os.LookupEnv("ROOT_DIR")
	if !ok {
		return errors.New("Failed to find 'ROOT_DIR' env!")
	}
	configs.RootDir = root
	port, ok := os.LookupEnv("PORT")
	if !ok {
		return errors.New("Failed to find 'PORT' env!")
	}
	configs.Port = port
	lb.Configs = &configs
	return nil
}

func (lb *LoadBalancer) setupWebPool(file string) error {
	path := lb.Configs.RootDir + file
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &lb.Pool)
	if err != nil {
		return err
	}
	return nil
}

func (lb *LoadBalancer) pingBalancedHandler(c *fiber.Ctx) error {
	if lb.Pool.Current == nil {
		lb.Pool.Current = &lb.Pool.WebURLs[0]
	} else {
		if err := lb.Pool.setNextIdx(); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON("message:'bad server pool'")
		}
	}
	uri := string(c.Request().URI().RequestURI())
	agent := fiber.Get(*lb.Pool.Current + uri)
	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}
	var pong fiber.Map
	if err := json.Unmarshal(body, &pong); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err,
		})
	}
	return c.Status(statusCode).JSON(pong)
}

func (lb *LoadBalancer) setupApp() error {
	lb.App = fiber.New()
	lb.App.Use(logger.New())
	lb.App.Get("/ping", lb.pingBalancedHandler)
	if err := lb.App.Listen(lb.Configs.Port); err != nil {
		return err
	}
	return nil
}

func main() {
	lb := LoadBalancer{}
	if err := lb.setupConfigs(); err != nil {
		panic(err)
	}
	if err := lb.setupWebPool("config.json"); err != nil {
		panic(err)
	}
	if err := lb.setupApp(); err != nil {
		panic(err)
	}
	fmt.Printf("[LB] [*] Listening on port{%s}...", lb.Configs.Port)
}
