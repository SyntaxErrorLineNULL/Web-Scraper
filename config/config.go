package config

import (
	"github.com/BurntSushi/toml"
)

type (
	// Config holds the overall configuration for the application, including settings for the application itself,
	// the database, and Redis. It uses TOML (Tom's Obvious, Minimal Language) tags to define the structure of the configuration file.
	Config struct {
		// filePath holds the path to the configuration file that will be loaded into the Config struct.
		filePath string
		// App holds the application-specific configuration, such as the address to listen on and the number of worker routines.
		App App `toml:"app"`
		// DataBase holds database connection settings, such as the URL, connection timeout, and pool size.
		DataBase DataBase `toml:"database"`
		// Redis holds the Redis connection configuration, including Redis server addresses, pool size, and timeout settings.
		Redis Redis `toml:"redis"`
	}

	// App holds configuration parameters related to the application itself.
	App struct {
		// Listen specifies the address and port on which the application should listen for incoming requests.
		Listen string `toml:"listen"`
		// WorkerCount specifies the number of worker routines that the application should spawn.
		WorkerCount int `toml:"workerCount"`
	}

	// DataBase defines the database connection settings.
	DataBase struct {
		// URL is the database connection string used to connect to the database.
		URL string `toml:"url"`
		// ConnectTimeout is the timeout (in seconds) for establishing a connection to the database.
		ConnectTimeout int `toml:"connect_timeout"`
		// PoolSize specifies the number of database connections in the connection pool.
		PoolSize int `toml:"pool_size"`
	}

	// Redis holds configuration parameters for connecting to multiple Redis instances.
	Redis struct {
		// Redis is a list of Redis server addresses.
		Redis []string `toml:"redis"`
		// ConnectionSize specifies the size of the Redis connection pool.
		ConnectionSize int `toml:"pool_size"`
		// DialTimeout specifies the maximum duration (in seconds) to wait for a Redis connection to be established.
		DialTimeout int64 `toml:"dial_timeout"`
		// WriteTimeout specifies the maximum duration (in seconds) to wait for writing data to a Redis connection.
		WriteTimeout int64 `toml:"write_timeout"`
		// ReadTimeout specifies the maximum duration (in seconds) to wait for reading data from a Redis connection.
		ReadTimeout int64 `toml:"read_timeout"`
	}
)

// NewConfig is a constructor function that returns a new Config object with the specified file path.
// The file path will later be used to load configuration settings from a TOML file.
func NewConfig(filePath string) *Config {
	// Return a new instance of Config with the provided file path for the configuration file.
	return &Config{filePath: filePath}
}

// Load reads the configuration from the TOML file specified by the filePath field in the Config struct.
// This method attempts to decode the file into the Config struct, ensuring that all necessary configuration
// values are loaded for the application. If an error occurs during file reading or decoding, the method
// returns an error indicating the problem.
func (c *Config) Load() (*Config, error) {
	// Create a variable `cfg` of type Config that will store the decoded configuration values.
	// This struct will hold the configuration read from the TOML file.
	var cfg Config

	// Attempt to decode the TOML file located at the `filePath` field of the Config object `c`.
	// The file is expected to contain key-value pairs matching the structure of the `Config` type.
	// The decoded data is stored in the `cfg` variable.
	_, err := toml.DecodeFile(c.filePath, &cfg)

	// If the TOML file cannot be read or decoded (e.g., due to a missing file or syntax errors),
	// the error is captured in `err`. In that case, return `nil` for the configuration and the error
	// to indicate failure. This prevents the application from proceeding with an invalid configuration.
	if err != nil {
		return nil, err
	}

	// If no error occurred during decoding, return a pointer to the successfully populated `cfg` struct.
	// Also return `nil` for the error, indicating that the configuration was loaded successfully.
	return &cfg, nil
}
