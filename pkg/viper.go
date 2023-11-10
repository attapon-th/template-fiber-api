package pkg

import "github.com/spf13/viper"

// SetDefaultConfig set default for viper
func SetDefaultConfig() {

	viper.SetDefault("dev", 0)
	viper.SetDefault("host", "0.0.0.0")
	viper.SetDefault("port", "8888")
	viper.SetDefault("prefork", "1")
	viper.SetDefault("prefix", "/")

	// log
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("LOG_DIR", "storage/logs")
	viper.SetDefault("LOG_OUTPUT", "log.log")
	viper.SetDefault("LOG_ACCESS", "access.log")
	viper.SetDefault("LOG_CALLER", 0)
}
