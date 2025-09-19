package config

import (
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	// Logs LogConfig
	// DB   PostgresConfig
	JWT  JwtConfig
	Port string
}

// TODO: For later
// type LogConfig struct {
// 	Style string
// 	Level string
// }

// type PostgresConfig struct {
// 	Username string
// 	Password string
// 	URL      string
// 	Port     string
// }

type JwtConfig struct {
	Key string
}

var instance *Config
var once sync.Once

func GetConfig() (*Config, error) {
	envFile, err := godotenv.Read(".env")
	if err != nil {
		return nil, err
	}

	once.Do(func() {
		instance = &Config{
			Port: envFile["SERVER_PORT"],
			// Logs: LogConfig{
			// 	Style: os.Getenv("LOG_STYLE"),
			// 	Level: os.Getenv("LOG_LEVEL"),
			// },
			// DB: PostgresConfig{
			// 	Username: os.Getenv("POSTGRES_USER"),
			// 	Password: os.Getenv("POSTGRES_PWD"),
			// 	URL:      os.Getenv("POSTGRES_URL"),
			// 	Port:     os.Getenv("POSTGRES_PORT"),
			// },
			JWT: JwtConfig{
				Key: envFile["JWT_KEY"],
			},
		}
	})

	return instance, nil
}
