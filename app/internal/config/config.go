package config

type Config struct {
	IP       string `env:"IP,required"`
	HTTPPort string `env:"HTTP_PORT,required"`
	Workers  int
}

// New создает экземпляр Config и заполняет его значениями переменных окружения.
func New(workers int) *Config {
	cfg := &Config{
		Workers:  workers,
		IP:       "0.0.0.0",
		HTTPPort: "8000",
	}

	return cfg
}
