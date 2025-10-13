package config

import (
	"path"
	"runtime"
	"sync"

	"github.com/joho/godotenv"
)

var RoothPath = rootDir()

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	return path.Join(path.Dir(b), "../..")
}

type Config struct {
	// Logs LogConfig
	// DB   PostgresConfig
	JWT  JwtConfig
	DB   DbConfig
	Port string
}

// TODO: For later
// type LogConfig struct {
// 	Style string
// 	Level string
// }

type DbConfig struct {
	Username   string
	Password   string
	URL        string
	Port       string
	SQLitePath string
}

type JwtConfig struct {
	Key string
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		envFile, err := godotenv.Read(path.Join(RoothPath, ".env"))
		if err != nil {
			panic(err)
		}

		instance = &Config{
			Port: envFile["SERVER_PORT"],
			// Logs: LogConfig{
			// 	Style: os.Getenv("LOG_STYLE"),
			// 	Level: os.Getenv("LOG_LEVEL"),
			// },
			DB: DbConfig{
				Username:   envFile["SERVER_PORT"],
				Password:   envFile["SERVER_PORT"],
				URL:        envFile["SERVER_PORT"],
				Port:       envFile["SERVER_PORT"],
				SQLitePath: envFile["SQLITE_PATH"],
			},
			JWT: JwtConfig{
				Key: envFile["JWT_KEY"],
			},
		}
	})

	return instance
}
