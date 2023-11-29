package pkg

import (
	"crypto/sha256"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/attapon-th/template-fiber-api/helper"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbs sync.Map

	defaultGormLoggerConfig = logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Warn,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	}
)

// ConnectPostgreSQL connect postgresql database
func ConnectPostgreSQL(dsn string, cfgs ...*gorm.Config) (*gorm.DB, error) {

	h := sha256.New224()
	hStr := helper.B2S(h.Sum([]byte(dsn)))

	if d, ok := dbs.Load(hStr); ok && d != nil {
		if db, ok := d.(*gorm.DB); ok && db != nil {
			return db, nil
		}
	}
	cfg := &gorm.Config{}
	if len(cfgs) == 0 {
		cfg = &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			DisableNestedTransaction:                 true,
			Logger:                                   getGormLogger(),
		}
	} else {
		cfg = cfgs[0]
	}

	db, err := gorm.Open(postgres.Open(dsn), cfg)
	if err != nil {
		return nil, err
	}
	dbs.Store(hStr, db)
	return db, nil

}

func getGormLogger() logger.Interface {
	logSQL := viper.GetString("log_sql")
	l := logger.Discard
	if logSQL == "default" {
		l = logger.New(newGormLoggerFileWriter(viper.GetString("log_file")), defaultGormLoggerConfig)
	} else if strings.HasSuffix(logSQL, ".log") {
		cfg := defaultGormLoggerConfig
		cfg.LogLevel = logger.Info
		l = logger.New(newGormLoggerFileWriter(logSQL), cfg)
	} else {
		l = logger.Discard
	}
	return l
}

type gormLoggerWriter struct {
	log  zerolog.Logger
	mode string
}

func (g *gormLoggerWriter) Printf(s string, a ...any) {
	g.log.Log().Str("logger", "gorm").Str("mode", g.mode).Msgf(s, a...)
}

func newGormLoggerConsoleWriter() *gormLoggerWriter {
	cs := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: zerolog.TimeFieldFormat,
		NoColor:    false,
	}
	return &gormLoggerWriter{
		log:  zerolog.New(cs).With().Timestamp().Logger(),
		mode: "console",
	}
}

func newGormLoggerFileWriter(filename string) *gormLoggerWriter {
	cs := NewDiodeCronWriter(filename)
	return &gormLoggerWriter{
		log:  zerolog.New(cs).With().Timestamp().Logger(),
		mode: "filemode",
	}
}
