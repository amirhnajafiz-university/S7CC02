package crypto

// Config stores data for api connection.
type Config struct {
	Crypto string `koanf:"crypto"`
	Value  string `koanf:"value"`
}
