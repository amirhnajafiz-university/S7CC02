package http

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/ceit-aut/S7CC02/internal/services/crypto"
	"github.com/ceit-aut/S7CC02/internal/storage"

	"github.com/gofiber/fiber/v2"
)

// redis key for storing set.
const redisKey = "crypto"

// Handler manages the endpoints implementations.
type Handler struct {
	Crypto  *crypto.Client
	Storage *storage.Storage
}

// Health returns service status.
func (h *Handler) Health(ctx *fiber.Ctx) error {
	if err := h.Storage.Ping(ctx.Context()); err != nil {
		return ctx.SendString(fmt.Sprintf("%s\n%v\n", time.Now(), err))
	}

	return ctx.SendString(fmt.Sprintf("%s\nOK\n", time.Now()))
}

// Get request for getting crypto prizes.
func (h *Handler) Get(ctx *fiber.Ctx) error {
	var (
		response Response
	)

	log.Println("service accepted user request")

	if value, err := h.Storage.Read(redisKey); err == nil {
		// get crypto name
		response.Name = h.Crypto.GetCryptoName()
		// convert price to float32
		p, _ := strconv.ParseFloat(value, 32)
		response.Price = float32(p)
	} else {
		// get crypto status
		name, price, er := h.Crypto.Get()
		if er != nil {
			return er
		}

		response.Name = name
		response.Price = price

		// save to redis
		go func() {
			if e := h.Storage.Write(redisKey, fmt.Sprintf("%f", price)); e != nil {
				log.Printf("redis store failed:\n\t%v\n", e)
			}
		}()
	}

	return ctx.JSON(response)
}
