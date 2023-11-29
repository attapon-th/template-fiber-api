package pkg

import "github.com/spf13/viper"

func init() {
	// only use viper.SetDefault

	viper.SetDefault("enveronment", "production")
	viper.SetDefault("dev", 0) // debug mode
	viper.SetDefault("domain", "localhost")
	viper.SetDefault("host", "0.0.0.0")
	viper.SetDefault("port", "8888")
	viper.SetDefault("prefork", "1")
	viper.SetDefault("prefix", "/") // url prefix

	// log
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("LOG_DIR", "storage/logs") // set directory for store log files

	// LOG_FILE
	// console:   print log with os.Stdout
	// $Filename: save log into filename
	viper.SetDefault("LOG_FILE", "console") // console or file name
	// viper.SetDefault("LOG_FILE", "log.log") // console or file name

	// LOG_CALLER
	// 1:  is show caller code
	// 0: disabled
	viper.SetDefault("LOG_CALLER", 0) //  0 or 1 is show caller code

	// LOG_SQL
	// discard:   Discard logger will print any log to io.Discard
	// default:   gorm default log
	// $Filename: save traceRecoder into filename (if file extension = `.log`)
	viper.SetDefault("LOG_SQL", "default")
}
