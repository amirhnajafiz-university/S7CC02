package cmd

import (
	"fmt"
	"log"

	"github.com/ceit-aut/cryptometer/internal/config"
	"github.com/ceit-aut/cryptometer/internal/port/http"
	"github.com/ceit-aut/cryptometer/internal/services/crypto"
	"github.com/ceit-aut/cryptometer/internal/storage"
	"github.com/gofiber/fiber/v2"
)

// Execute application.
func Execute() {
	// load configs
	cfg := config.Load()

	// create new fiber app
	app := fiber.New()

	// create redis client
	rdb := storage.New(cfg.Storage)

	// create crypto service
	crp := crypto.New(cfg.Crypto)

	// create handler
	h := http.Handler{
		Crypto:  crp,
		Storage: rdb,
	}

	// define endpoints
	app.Get("/api", h.Get)

	// start server
	log.Println(app.Listen(fmt.Sprintf(":%d", cfg.Port)))
}
