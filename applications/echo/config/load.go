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
	logger.Info("[せってい] - せっていめいをせっていちゅう")
	viper.SetConfigName("config")
	logger.Info("[せってい] - せっていたいぷをせっていちゅう")
	viper.SetConfigType("yaml")
	logger.Info("[せってい] - せっていぱすをついかちゅう")
	viper.AddConfigPath(".")
	if testing.Testing() {
		logger.Info("[せってい] - てすとかんきょうをけんしゅつ")
		root, err := find.Repo()
		if err != nil {
			logger.Infof("[せってい] - りぽじとりーるーとのしゅとくにしっぱい: %v", err)
			return nil, err
		}
		logger.Info("[せってい] - りぽじとりーぱすをついか")
		viper.AddConfigPath(root.Path)
	}

	viper.AddConfigPath("./")
	logger.Info("[せってい] - かんきょうへんすうりぷれーさーをせってい")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	logger.Info("[せってい] - じどうかんきょうへんすうよみこみをゆうこうか")
	viper.AutomaticEnv()
	logger.Info("[せってい] - でふぉるとせっていをよみこみちゅう")
	c := loadDefault()
	logger.Info("[せってい] - でふぉるとせっていをましゃるちゅう")
	if configBuffer, err := json.Marshal(c); err != nil {
		logger.Infof("[せってい] - でふぉるとせっていのましゃるにしっぱい: %v", err)
		return nil, err
	} else if err := viper.ReadConfig(bytes.NewBuffer(configBuffer)); err != nil {
		logger.Infof("[せってい] - でふぉるとせっていのよみこみにしっぱい: %v", err)
		return nil, err
	}
	logger.Info("[せってい] - せっていふぁいるをまーじちゅう")
	if err := viper.MergeInConfig(); err != nil {
		logger.Infof("[せってい] - せっていふぁいるのまーじにしっぱい（でふぉるとをしよう）: %v", err)
	}
	logger.Info("[せってい] - せっていをあんましゃるちゅう")
	err := viper.Unmarshal(c)
	logger.Info("[せってい] - せっていよみこみかんりょう")
	return c, err
}
