package crypto

// Config stores data for api connection.
type Config struct {
	Crypto string `koanf:"name"`
	Value  string `koanf:"value"`
}
