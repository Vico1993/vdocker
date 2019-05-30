package tools

import (
	"github.com/spf13/viper"
)

// Config is the struct to handle vdocker config json file
type Config struct {
	TokenTelegram string
}

// LoadConfig will read the config file and build the Conf variable
func LoadConfig() (Config, error) {
	// Available PATH
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.vdocker")
	// config file name
	viper.SetConfigName("vdocker")
	viper.SetConfigType("json")

	// Read Config file
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	// Unmarshal the JSON file into a Config structure
	var conf Config
	err = viper.Unmarshal(&conf)
	if err != nil {
		return Config{}, err
	}

	return conf, nil
}
