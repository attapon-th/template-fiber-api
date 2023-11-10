/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// servCmd represents the serv command
var servCmd = &cobra.Command{
	Use:   "serv",
	Short: "Start server for api (Listen on 0.0.0.0:8888)",
	Long:  `Start API Server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serv called")
		// pkg.SetDefaultConfig()
		// pkg.InitZeroLog()
		// fb := pkg.NewFiber()

		// // create routes
		// route.Init(fb.Fiber())

		// log.Info().Int("port", fb.Port).Msg("Server started")
		// log.Info().Msgf("Version: %s", Version)

		// if viper.GetBool("dev") && !fiber.IsChild() {
		// 	pringLog()
		// }
		// fb.Listen()
	},
}

func init() {
	rootCmd.AddCommand(servCmd)
}

func pringLog() {
	for _, v := range os.Environ() {
		if strings.HasPrefix(v, "FB_") {
			k := strings.TrimPrefix(v, "FB_")
			k = k[0:strings.Index(k, "=")]
			// viper.BindEnv(k)
			k = strings.ToLower(k)
			fmt.Printf("%-20s: %v\n", k, viper.Get(k))
		}
	}
}
