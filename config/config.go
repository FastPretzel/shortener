package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Addr         string
	DBURL        string
	InMemoryMode bool
	PathLog      string
	PortGRPC     string
	PortHTTP     string
}

func New() *Config {
	viper.SetConfigFile(".env")
	//viper.SetConfigType("yaml")

	//viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		//if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		//fmt.Fprintf(os.Stderr, "config file not found: %s", err.Error())
		//} else {
		//fmt.Fprintf(os.Stderr, "failed to read config: %s", err.Error())
		//}
		fmt.Fprintf(os.Stderr, "failed to read config: %s", err.Error())
		os.Exit(1)
	}

	return &Config{
		Addr:         viper.GetString("API_ADDRESS"),
		DBURL:        viper.GetString("DB_URL"),
		PathLog:      viper.GetString("API_LOG_PATH"),
		InMemoryMode: viper.GetBool("API_STORAGE_MODE"),
		PortGRPC:     viper.GetString("API_PORT_GRPC"),
		PortHTTP:     viper.GetString("API_PORT_HTTP"),
	}
}
