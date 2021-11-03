package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Provisions a local Kubernetes cluster.",
	RunE:  create,
}

func create(cmd *cobra.Command, args []string) error {
	fmt.Println("create called")
	return nil
}

func init() {
	rootCmd.AddCommand(createCmd)
}
