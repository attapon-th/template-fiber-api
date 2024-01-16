package routes

import (
	"net/url"

	"github.com/attapon-th/template-fiber-api/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func newSwagerRoute(r fiber.Router, apiPathPrefix string) {
	if baseURL, err := url.Parse(viper.GetString("base_url")); err == nil {
		// Custom Swagger URL
		docs.SwaggerInfo.BasePath = apiPathPrefix
		docs.SwaggerInfo.Host = baseURL.Host
		docs.SwaggerInfo.Schemes = []string{baseURL.Scheme}
		log.Debug().Str("host", baseURL.Host).Str("schema", baseURL.Scheme).Str("prefix", apiPathPrefix).Msg("set swagger url.")
	}
	swaggerConfig := swagger.ConfigDefault
	r.Get("/*", swagger.New(swaggerConfig))

}
