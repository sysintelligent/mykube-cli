package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// minikubeCmd represents the minikube command
var minikubeCmd = &cobra.Command{
	Use:   "minikube",
	Short: "Starts a local Kubernetes cluster.",
	RunE:  startMinikube,
}

func startMinikube(cmd *cobra.Command, args []string) error {
	prg := "minikube"
	arg1 := "--help"

	c := exec.Command(prg, arg1)
	stdout, err := c.Output()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Print(string(stdout))
	return nil
}

func init() {
	createCmd.AddCommand(minikubeCmd)
}
