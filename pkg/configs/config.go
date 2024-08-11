package configs

import (
	"github.com/spf13/viper"
)

type EnvConfigs struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBPort     int    `mapstructure:"DB_PORT"`
	ApiPort    int    `mapstructure:"API_PORT"`
}

func LoadEnv() EnvConfigs {
	envConfigs := EnvConfigs{}

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&envConfigs)
	if err != nil {
		panic(err)
	}

	return envConfigs
}
