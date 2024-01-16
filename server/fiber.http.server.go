package server

import (
	"errors"
	"fmt"
	"runtime"
	"time"

	"github.com/attapon-th/template-fiber-api/models"
	"github.com/attapon-th/template-fiber-api/routes"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func jsonMarshal(v interface{}) ([]byte, error) {
	return json.MarshalWithOption(v, json.UnorderedMap())
}

// NewFiber creates a new Fiber
func Listen(config ...fiber.Config) {
	var cfg fiber.Config
	if len(config) == 0 {
		cfg = fiber.Config{
			AppName:                      "API Service",
			ServerHeader:                 "API Service",
			BodyLimit:                    1 * 1024 * 1024, // 1MB
			ReadBufferSize:               10 * 1024,
			ReadTimeout:                  10 * time.Second,
			WriteTimeout:                 30 * time.Second,
			Prefork:                      false,
			DisablePreParseMultipartForm: true,
			DisableStartupMessage:        true,
			CaseSensitive:                true,
			EnablePrintRoutes:            true,
			Views:                        nil,
			JSONEncoder:                  jsonMarshal,
			JSONDecoder:                  json.Unmarshal,
			ErrorHandler:                 FiberErrorHandler,
		}
	} else {
		cfg = config[0]
	}

	viper.SetDefault("prefork", 1)
	viper.SetDefault("port", 8888)
	viper.SetDefault("host", "0.0.0.0")
	if n := viper.GetInt("prefork"); n > 1 {
		cfg.Prefork = true
		maxProcessing := runtime.GOMAXPROCS(n)
		if maxProcessing > n {
			maxProcessing = n
		}
		log.Info().Bool("prefork", cfg.Prefork).Int("maxProcessing", maxProcessing).Msg("Prefork enabled")
	}

	app := fiber.New(cfg)

	// Create Router
	routes.NewRouters(app)

	// start server
	host := viper.GetString("host")
	port := viper.GetInt("port")
	log.Info().Str("host", host).Int("port", port).Msg("Server started")
	err := app.Listen(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatal().Err(err).Msg("Server error")
	}
	log.Info().Msg("Server stopped")
}

// FiberErrorHandler handles errors
func FiberErrorHandler(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError
	msg := "Internal Server Error"

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		msg = e.Error()
	}
	c.Set("Content-Type", fiber.MIMEApplicationJSONCharsetUTF8)

	// Send custom error page
	_ = c.Status(code).JSON(models.NewAPIResponse(false, msg))

	// Return from handler
	return nil
}
