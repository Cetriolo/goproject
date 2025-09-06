package cmd

import (
	"v2/pkg/boostrap"

	"github.com/spf13/cobra"
)

// logger.go (create this in your package, e.g., in the "config" or "internal" folder)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Long:  "This command will run the database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		migrate()
	},
}

func migrate() {
	boostrap.Migrate()
}
