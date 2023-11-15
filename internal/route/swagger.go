package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func setupSwager(r fiber.Router, prefix string) {
	// if baseURL, err := url.Parse(viper.GetString("base_url")); err == nil {
	// 	docs.SwaggerInfo.BasePath = prefix
	// 	docs.SwaggerInfo.Host = baseURL.Host
	// 	docs.SwaggerInfo.Schemes = []string{baseURL.Scheme}
	// 	log.Debug().Str("host", baseURL.Host).Str("schema", baseURL.Scheme).Str("prefix", prefix).Msg("set swagger url.")
	// }
	swaggerConfig := swagger.ConfigDefault
	r.Get(prefix+"/swagger/*", swagger.New(swaggerConfig)) // default
}
