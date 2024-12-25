package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type config struct {
	Redis         redisConfig       `mapstructure:",squash"`
	RabbitMq      rabbitMqConfig    `mapstructure:",squash"`
	Database      databaseConfig    `mapstructure:",squash"`
	AuthConsole   authConsoleConfig `mapstructure:",squash"`
	Env           string            `mapstructure:"ENV"`
	ServiceName   string            `mapstructure:"SERVICE_NAME"`
	GlobalTimeout int               `mapstructure:"GLOBAL_TIMEOUT"`
}

type redisConfig struct {
	Host            string `mapstructure:"REDIS_HOST"`
	Port            string `mapstructure:"REDIS_PORT"`
	Password        string `mapstructure:"REDIS_PASSWORD"`
	PoolMaxSize     int    `mapstructure:"REDIS_POOL_MAX_SIZE"`
	PoolMinIdleSize int    `mapstructure:"REDIS_POOL_MIN_IDLE_SIZE"`
}

type rabbitMqConfig struct {
	Host     string `mapstructure:"RABBITMQ_HOST"`
	Port     string `mapstructure:"RABBITMQ_PORT"`
	Username string `mapstructure:"RABBITMQ_USERNAME"`
	Password string `mapstructure:"RABBITMQ_PASSWORD"`
}

type databaseConfig struct {
	Username string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	DbName   string `mapstructure:"DB_NAME"`
}

type authConsoleConfig struct {
	AuthIssuer   string `mapstructure:"CONSOLE_AUTH_ISSUER"`
	AuthAudience string `mapstructure:"CONSOLE_AUTH_AUDIENCE"`
	AuthCertsUrl string `mapstructure:"CONSOLE_AUTH_CERTS_URL"`
}

var viperInstance *viper.Viper
var configInstance config

func Env(filenames ...string) config {
	initViper(filenames...)
	return configInstance
}

func Viper(filenames ...string) *viper.Viper {
	initViper(filenames...)
	return viperInstance
}

func initViper(filenames ...string) {
	if viperInstance == nil {
		viperInstance = viper.New()

		if len(filenames) > 0 {
			viperInstance.SetConfigFile(filenames[0])
		} else {
			viperInstance.SetConfigFile(".env")
		}

		viperInstance.AutomaticEnv()
		err := viperInstance.ReadInConfig()
		if err != nil {
			fmt.Println("Error Config", err)
			return
		}

		err = viperInstance.Unmarshal(&configInstance)
		if err != nil {
			fmt.Println("Error Config", err)
			return
		}

	}
}
