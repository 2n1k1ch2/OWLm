package configs

import "github.com/spf13/viper"

var (
	CFG *viper.Viper
)

func init() {
	CFG := viper.New()
	CFG.SetConfigName("config")
	CFG.SetConfigType("yaml")
	CFG.AddConfigPath("./configs")

}
