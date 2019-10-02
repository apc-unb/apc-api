package cmd

import (
	"github.com/apc-unb/apc-api/web"
	"github.com/apc-unb/apc-api/web/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the HTTP REST APIs server",
	RunE: func(cmd *cobra.Command, args []string) error {
		webBuilder := new(config.WebBuilder).InitFromViper(viper.GetViper())
		server := new(web.Server).InitFromWebBuilder(webBuilder)
		return server.Run()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	config.AddFlags(serveCmd.Flags())

	err := viper.GetViper().BindPFlags(serveCmd.Flags())
	if err != nil {
		panic(err)
	}
}
