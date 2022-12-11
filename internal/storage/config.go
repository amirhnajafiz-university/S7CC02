package storage

// Config stores information of redis cluster configs.
type Config struct {
	Host     string `koanf:"host"`
	Password string `koanf:"password"`
	Timeout  int    `koanf:"timeout"`
}
