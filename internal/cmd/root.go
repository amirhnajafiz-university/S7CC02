package cmd

import (
	"fmt"

	"github.com/ceit-aut/cryptometer/internal/config"
)

// Execute application.
func Execute() {
	// load configs
	cfg := config.Load()

	fmt.Println(cfg)
}
