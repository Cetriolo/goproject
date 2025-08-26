package cmd

import (
	"v2/pkg/config/boostrap"

	"github.com/spf13/cobra"
)

// logger.go (create this in your package, e.g., in the "config" or "internal" folder)

func init() {
	rootCmd.AddCommand(seedCmd)
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Run database seeding",
	Long:  "This command will run the database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		seed()
	},
}

func seed() {
	boostrap.Seed()
}
