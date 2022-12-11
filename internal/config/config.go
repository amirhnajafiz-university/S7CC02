package config

import (
	"github.com/ceit-aut/cryptometer/internal/services/crypto"
	"github.com/ceit-aut/cryptometer/internal/storage"
)

// Config stores application configs.
type Config struct {
	Port    int            `koanf:"port"`
	Crypto  crypto.Config  `koanf:"crypto"`
	Storage storage.Config `koanf:"storage"`
}

// Load configs.
func Load() Config {
	return Default()
}
