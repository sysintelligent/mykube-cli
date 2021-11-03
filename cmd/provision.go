package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// provisionCmd represents the provision command
var provisionCmd = &cobra.Command{
	Use:   "provision",
	Short: "Provisions a local Kubernetes cluster",
	RunE:  provision,
}

func provision(cmd *cobra.Command, args []string) error {
	fmt.Println("provision called")
	return nil
}

func init() {
	rootCmd.AddCommand(provisionCmd)
}
