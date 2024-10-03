package config

type (
	Config struct {
		filePath string
	}

	App struct {
		// Listen specifies the address and port on which the application should listen for incoming requests.
		Listen string `toml:"listen"`
		// WorkerCount specifies the number of worker routines that the application should spawn.
		WorkerCount int `toml:"workerCount"`
	}

	DataBase struct {
		URL            string `toml:"url"`
		ConnectTimeout int    `toml:"connect_timeout"`
		PoolSize       int    `toml:"pool_size"`
	}

	// Redis holds configuration parameters for connecting to multiple Redis instances.
	Redis struct {
		Redis          []string `toml:"redis"`
		ConnectionSize int      `toml:"pool_size"`
		DialTimeout    int64    `toml:"dial_timeout"`
		WriteTimeout   int64    `toml:"write_timeout"`
		ReadTimeout    int64    `toml:"read_timeout"`
	}
)

func NewConfig(filePath string) *Config {
	return &Config{filePath: filePath}
}

func (c *Config) Load() (*Config, error) {

	return nil, nil
}
