package api

import (
	"github.com/ducconit/gobase/app"
	"github.com/ducconit/gobase/config"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the api server",
	Long:  "Start the api server.",
	RunE: func(cmd *cobra.Command, _ []string) error {
		configPath, _ := cmd.Flags().GetString("config")
		bindingAddr, _ := cmd.Flags().GetString("binding")

		return serveApiServer(configPath, bindingAddr)
	},
}

func init() {
	ServeCmd.Flags().StringP("binding", "b", "", "binding address")
}

func serveApiServer(cfgPath, addr string) error {
	cfg := config.NewFileStore(&config.FileStoreConfig{
		Path: cfgPath,
	})

	if err := cfg.Load(); err != nil {
		return err
	}
	a := app.NewApp(app.WithConfig(cfg))

	srv := app.NewServerApi(a)

	return srv.Run(addr)
}
