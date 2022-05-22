package init

import (
	"by291.fun/scaffold/global"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func ViperInit() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "read config err")
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		return errors.Wrap(err, "unmarshal config err")
	}

	return nil
}
