package util

import "github.com/spf13/viper"

type Config struct {
	DBDriver string `mapstructure:"db_driver"`
	DBSource string `mapstructure:"db_source"`
	ServerAddress string `mapstructure:"server_address"`
}

type ConfigWithEnvironments struct {
	Development Config `mapstructure:"development"`
	Test Config `mapstructure:"test"`
	Production Config `mapstructure:"production"`
}

func LoadConfig(path string, env string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	var allConfig ConfigWithEnvironments
	err = viper.Unmarshal(&allConfig)
	switch env {
	case "development":
		config = allConfig.Development
	case "test":
		config = allConfig.Test
	case "production":
		config = allConfig.Production
	}
	return
}
