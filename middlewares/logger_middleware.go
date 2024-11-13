package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

func Logger(c *fiber.Ctx) error {
	start := time.Now()

	err := c.Next()

	log.Printf("%s %s %d %s", c.Method(), c.Path(), c.Response().StatusCode(), time.Since(start))
	return err
}
