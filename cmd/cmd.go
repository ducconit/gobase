package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gobase",
	Short: "Skeleton for go projects",
	Long:  "Skeleton for go projects.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		checkForRootUser()
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "config.yml", "config file path")
}

func Execute() error {
	return rootCmd.Execute()
}

// checkForRootUser logs a warning if the process is running as root
func checkForRootUser() {
	if os.Geteuid() == 0 {
		log.Println("Running as root is not recommended. Please use a non-root user.")
	}
}
