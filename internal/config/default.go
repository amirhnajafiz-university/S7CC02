package config

import (
	"github.com/ceit-aut/cryptometer/internal/services/crypto"
	"github.com/ceit-aut/cryptometer/internal/storage"
)

// Default config values.
func Default() Config {
	return Config{
		Port: 6030,
		Crypto: crypto.Config{
			Crypto: "bitcoin",
			Value:  "usd",
		},
		Storage: storage.Config{
			Host:     "localhost:6337",
			Password: "",
			Timeout:  5,
		},
	}
}
