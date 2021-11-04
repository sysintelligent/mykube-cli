package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

const (
	memory string = "16384"
	cpus   string = "4"
)

// minikubeCmd represents the minikube command
var minikubeCmd = &cobra.Command{
	Use:   "minikube",
	Short: "Starts a local minikube cluster",
	RunE:  startMinikube,
}

func startMinikube(cmd *cobra.Command, args []string) error {
	// minikube start --memory=16384 --cpus=4 --kubernetes-version=v1.21.3
	startCmd := []string{"start",
		"--memory", memory,
		"--cpus", cpus,
	}

	fmt.Println("Create Minikube instance ...")

	RunCmd(startCmd)
	return nil
}

func RunCmd(arg []string) error {
	cmd := exec.Command("minikube", arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func init() {
	provisionCmd.AddCommand(minikubeCmd)
}
