package http

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ceit-aut/cryptometer/internal/services/crypto"
	"github.com/ceit-aut/cryptometer/internal/storage"
	"github.com/gofiber/fiber/v2"
)

// redis key for storing set.
const redisKey = "crypto"

// Handler manages the endpoints implementations.
type Handler struct {
	Crypto  crypto.Client
	Storage *storage.Storage
}

// Get request for getting crypto prizes.
func (h *Handler) Get(ctx *fiber.Ctx) error {
	var (
		response Response
	)

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
