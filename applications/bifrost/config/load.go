package config

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/YumikoKawaii/shared/logger"
	"github.com/spf13/viper"
)

// Load system env config
func Load() (*Application, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.AddConfigPath("./")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()
	c := loadDefault()
	if configBuffer, err := json.Marshal(c); err != nil {
		logger.Infof("[config] failed to marshal default config: %v", err)
		return nil, err
	} else if err := viper.ReadConfig(bytes.NewBuffer(configBuffer)); err != nil {
		logger.Infof("[config] failed to read default config: %v", err)
		return nil, err
	}
	if err := viper.MergeInConfig(); err != nil {
		logger.Infof("[config] failed to merge config file (using defaults): %v", err)
	}
	err := viper.Unmarshal(c)
	logger.Info("[config] config loaded successfully")
	return c, err
}
