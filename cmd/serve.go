package cmd

import (
	"v2/pkg/config/boostrap"

	"github.com/spf13/cobra"
)

// logger.go (create this in your package, e.g., in the "config" or "internal" folder)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve app on the server",
	Long:  "Application will be server on host and port in config.yml",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	boostrap.Serve()
}
