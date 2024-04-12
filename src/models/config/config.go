package config

type Config struct {
}

func ReadConfig(p string) (*Config, error) {
	return &Config{}, nil
}
