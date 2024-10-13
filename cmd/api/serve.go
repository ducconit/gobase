package api

import (
	"context"
	"github.com/ducconit/gobase/api"
	"github.com/ducconit/gobase/api/server"
	"github.com/ducconit/gobase/app"
	"github.com/ducconit/gobase/config"
	"github.com/ducconit/gobase/db"
	"github.com/ducconit/gobase/utils"
	"github.com/spf13/cobra"
	"log"
	"time"
)

var ServeCmd = &cobra.Command{
	Use:   "api:serve",
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

	d, err := db.NewFromConfig(cfg)
	if err != nil {
		return err
	}

	a := app.NewApp(app.WithConfig(cfg), app.WithDatabase(d))

	e := server.NewEchoServer(a)

	srv, err := api.NewServer(e)
	if err != nil {
		return err
	}

	utils.RegisterSignalDefaultHandler(func() {
		log.Println("shutting down gracefully, press Ctrl+C again to force")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Println(err)
			return
		}
	})

	return srv.Listen(addr)
}
