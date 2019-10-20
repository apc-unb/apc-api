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
	Example: ` 

MAC && LINUX

./apc-api serve \
  --port 8080 \
  --mongo-host localhost \
  --mongo-port 27017 \
  --jwt-key SUPER_SECRET \
  --codeforces-key f3d968eea83ad8d5f21cad0365edcc200439c6f0 \
  --codeforces-secret b30c206b689d5ba004534c6780aa7be8e234a7f3 \
  --log-level debug

WINDOWS POWER SHELL

./apc-api serve --port 8080 --mongo-host localhost --mongo-port 27017 --jwt-key SUPER_SECRET --codeforces-key f3d968eea83ad8d5f21cad0365edcc200439c6f0 --codeforces-secret b30c206b689d5ba004534c6780aa7be8e234a7f3 --log-level debug


All command line options can be provided via environment variables by adding the prefix "DRAGONT_" 
and converting their names to upper case and replacing punctuation and hyphen with underscores. 
For example,

command line option                 environment variable
------------------------------------------------------------------
--mongo-host localhost              DRAGONT_MONGO_HOST=localhost
--mongo.port		             	DRAGONT_MONGO_PORT=27017
	`,
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
