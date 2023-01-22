package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Addr        string `mapstructure:"API_ADDRESS"`
	DBURL       string `mapstructure:"DB_URL"`
	PathLog     string `mapstructure:"API_LOGPATH"`
	StorageMode string `mapstructure:"API_STORAGE_MODE"`
	PortGRPC    string `mapstructure:"API_PORT_GRPC"`
	PortHTTP    string `mapstructure:"API_PORT_HTTP"`
}

func New(path string) *Config {
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		//if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		//fmt.Fprintf(os.Stderr, "config file not found: %s", err.Error())
		//} else {
		//fmt.Fprintf(os.Stderr, "failed to read config: %s", err.Error())
		//}
		fmt.Fprintf(os.Stderr, "failed to read config: %s\n", err.Error())
		os.Exit(1)
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to unmarshal config: %s\n", err.Error())
		os.Exit(1)
	}

	return &cfg
	//Addr:        viper.GetString("API_ADDRESS"),
	//DBURL:       viper.GetString("DB_URL"),
	//PathLog:     viper.GetString("API_LOG_PATH"),
	//StorageMode: viper.GetString("API_STORAGE_MODE"),
	//PortGRPC:    viper.GetString("API_PORT_GRPC"),
	//PortHTTP:    viper.GetString("API_PORT_HTTP"),
	//}
}
