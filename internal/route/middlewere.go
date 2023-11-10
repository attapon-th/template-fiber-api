package route

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	memory "github.com/gofiber/storage/memory/v2"
)

func newCORS() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: "*",
	})
}

// import
func newLimit() fiber.Handler {
	return limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        100,
		Expiration: 10 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return fiber.NewError(400, "Too many requests")
		},
		Storage: memory.New(memory.Config{
			GCInterval: 10 * time.Second,
		}),
	})
}

func newAccessLog() fiber.Handler {
	// var wr io.Writer = os.Stdout
	// f, err := pkg.NewLogfileWriter(viper.GetString("log_access"))
	// if err != nil {
	// 	log.Fatal().Str("file", viper.GetString("log_access")).Msg(err.Error())
	// }
	// if f != nil {
	// 	wr = zerolog.MultiLevelWriter(os.Stdout, f)
	// }
	// return logger.New(logger.Config{
	// 	Format:        "[${time}][${method}][${status}][${latency}][${ip}] - ${url}\n",
	// 	TimeFormat:    "2006-01-02T15:04:05.999Z07:00",
	// 	TimeZone:      "Asia/Bangkok",
	// 	DisableColors: true,
	// 	Output:        wr,
	// })
	return nil
}

func newCompress() fiber.Handler {
	return compress.New(compress.Config{
		// set default compression all request
		Level: compress.LevelBestSpeed, // 1
	})
}
