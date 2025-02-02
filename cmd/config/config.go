package config

type Config struct {
}

func NewConfigFromFlags() (*Config, error) {
	return &Config{}, nil
}
