package config

// Config stores application configs.
type Config struct {
}

// Load configs.
func Load() Config {
	return Default()
}
