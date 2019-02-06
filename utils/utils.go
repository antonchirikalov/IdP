package utils

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

var (
	v *viper.Viper
)


func init() {
	v = viper.New()
	v.SetConfigType("yaml")

	v.AddConfigPath(".././conf/")
	v.AddConfigPath("./conf/")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := v.ReadInConfig()
	v.SetEnvPrefix(v.GetString("appname"))
	if err != nil {
		log.Fatal(err)
	}
}

func GetConfig() *viper.Viper {
	return v
}



