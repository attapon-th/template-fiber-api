package route

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func newJWK() fiber.Handler {
	return jwtware.New(jwtware.Config{
		JWKSetURLs:     []string{viper.GetString("JWK_URL")},
		SuccessHandler: nil,
		ContextKey:     "user_token",
	})
}
