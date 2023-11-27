package pkg

import "github.com/spf13/viper"

func init() {
	// only use viper.SetDefault

	viper.SetDefault("enveronment", "production")
	viper.SetDefault("dev", 0) // debug mode
	viper.SetDefault("host", "0.0.0.0")
	viper.SetDefault("port", "8888")
	viper.SetDefault("prefork", "1")
	viper.SetDefault("prefix", "/") // url prefix

	// log
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("LOG_DIR", "storage/logs") // set directory for store log files
	viper.SetDefault("LOG_FILE", "console")     // console or file name
	// viper.SetDefault("LOG_FILE", "log.log") // console or file name
	viper.SetDefault("LOG_CALLER", 0) //  0 or 1 is show caller code

}
