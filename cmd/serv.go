/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// servCmd represents the serv command
var servCmd = &cobra.Command{
	Use:   "serv",
	Short: "Start server for api (Listen on 0.0.0.0:8888)",
	Long:  `Start API Server`,
	Run: func(cmd *cobra.Command, args []string) {

		// pkg.InitZeroLog()
		// fb := pkg.NewFiber()

		// // create routes
		// routes.NewRouters(fb.Fiber())

		// log.Info().Int("port", fb.Port).Msg("Server started")
		// log.Info().Msgf("Version: %s", Version)

		// if viper.GetBool("dev") && !fiber.IsChild() {
		// 	pringLog()
		// }
		// if err := fb.Listen(); err != nil {
		// 	log.Fatal().Err(err).Send()
		// }
	},
}

func init() {
	rootCmd.AddCommand(servCmd)
}

func pringLog() {
	for _, k := range viper.AllKeys() {
		log.Debug().Interface(k, viper.Get(k)).Send()
	}
}
