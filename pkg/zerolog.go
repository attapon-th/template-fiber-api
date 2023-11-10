package pkg

import (
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// InitZeroLog initializes zerolog
func InitZeroLog() {
	viper.SetDefault("log_level", zerolog.InfoLevel)
	viper.SetDefault("log_output", "")
	viper.SetDefault("log_caller", false)

	lvl, err := zerolog.ParseLevel(viper.GetString("log_level"))
	if err != nil {
		lvl = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(lvl)
	zerolog.TimeFieldFormat = time.RFC3339Nano
	f, err := NewLogfileWriter(viper.GetString("log_output"))
	if err != nil {
		log.Fatal().Str("file", viper.GetString("log_output")).Msg(err.Error())
	}

	wr := zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02T15:04:05.999Z07:00"}, f)
	log.Logger = log.Output(wr)
	if viper.GetBool("log_caller") {
		setCaller()
		log.Logger = log.With().Caller().Logger()
	}
	log.Debug().Msg("log level set to " + lvl.String())
}

func setCaller() {
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short
		return file + ":" + strconv.Itoa(line)
	}
}
