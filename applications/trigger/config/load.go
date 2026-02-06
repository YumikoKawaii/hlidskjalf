package config

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/YumikoKawaii/shared/logger"
	"github.com/integralist/go-findroot/find"
	"github.com/spf13/viper"
)

// Load system env config
func Load() (*Application, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if testing.Testing() {
		root, err := find.Repo()
		if err != nil {
			logger.Infof("[せってい] - りぽじとりーるーとのしゅとくにしっぱい: %v", err)
			return nil, err
		}
		viper.AddConfigPath(root.Path)
	}

	viper.AddConfigPath("./")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()
	c := loadDefault()
	if configBuffer, err := json.Marshal(c); err != nil {
		logger.Infof("[せってい] - でふぉるとせっていのましゃるにしっぱい: %v", err)
		return nil, err
	} else if err := viper.ReadConfig(bytes.NewBuffer(configBuffer)); err != nil {
		logger.Infof("[せってい] - でふぉるとせっていのよみこみにしっぱい: %v", err)
		return nil, err
	}
	if err := viper.MergeInConfig(); err != nil {
		logger.Infof("[せってい] - せっていふぁいるのまーじにしっぱい（でふぉるとをしよう）: %v", err)
	}
	err := viper.Unmarshal(c)
	logger.Info("[せってい] - せっていよみこみかんりょう")
	return c, err
}
