package cmd

import (
	"github.com/ducconit/gobase/cmd/api"
	"github.com/spf13/cobra"
)

var ApiCmd = &cobra.Command{
	Use:   "api",
	Short: "Related commands for api",
	Long:  "Related commands for api.",
}

func init() {
	ApiCmd.AddCommand(api.ServeCmd)
	ApiCmd.RunE = api.ServeCmd.RunE

	rootCmd.RunE = ApiCmd.RunE
	rootCmd.AddCommand(ApiCmd)
}
