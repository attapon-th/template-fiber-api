package pkg

import (
	"os"
	"path"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/utahta/go-cronowriter"
)

// CronFileWriter create log file The file path is constructed based on current date and time, like cronolog.
func CronFileWriter(filename string) *cronowriter.CronoWriter {
	return cronowriter.MustNew(filename+".%Y%m%d", cronowriter.WithInit())
}

// NewLogfileWriter create log file
func NewLogfileWriter(filename string) (*cronowriter.CronoWriter, error) {
	if filename == "" {
		return nil, nil
	}
	logDir := viper.GetString("log_dir")
	fout := filename
	fout = path.Join(logDir, fout)
	err := os.MkdirAll(path.Dir(fout), os.ModePerm)
	if err != nil {
		log.Fatal().Str("file", fout).Msg(err.Error())
		return nil, err
	}
	f := CronFileWriter(fout)
	return f, nil
}
