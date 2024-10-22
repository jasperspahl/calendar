package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
)

type ColorConfig struct {
	Accent    lipgloss.Color
	LightGray lipgloss.Color
	Muted     lipgloss.Color
}

type Config struct {
	Colors ColorConfig
}

var defaultConfig = Config{
	Colors: ColorConfig{
		Accent:    lipgloss.Color("3"),
		LightGray: lipgloss.Color("#888888"),
		Muted:     lipgloss.Color("9"),
	},
}

var config *Config

func init() {
	configPath, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("Unable to load get configPath: %v", err)
	}
	configPath += "/calendar"
	viper.AutomaticEnv()
	viper.SetEnvPrefix("CALENDAR")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.SetConfigPermissions(0600)
	viper.AddConfigPath(configPath)

	viper.SetDefault("colors", defaultConfig.Colors)

	if err := viper.ReadInConfig(); err != nil {
		var errConfigNotFound viper.ConfigFileNotFoundError
		if !errors.As(err, &errConfigNotFound) {
			log.Fatalf("Unexpected error accoured when trying to load config: %v", err)
		}

		err := os.MkdirAll(configPath, os.ModePerm)
		if err != nil {
			log.Fatalf("Unable to create config directory: %v", err)
		}

		viper.SetConfigFile(fmt.Sprintf("%s/config.yml", configPath))
		if err := viper.WriteConfig(); err != nil {
			log.Fatalf("Unable to write config file: %v", err)
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to load config: %v", err)
	}
}

func GetConfig() *Config {
	return config
}
