package config

type (
	Config struct {
		filePath string
	}
)

func NewConfig(filePath string) *Config {
	return &Config{filePath: filePath}
}

func (c *Config) Load() (*Config, error) {

	return nil, nil
}
