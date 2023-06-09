package app

import (
	"mine-kube/cmd/app/cmd"
	"os"
)

func Run() error {
	cmd := cmd.NewCloudCommand(os.Stdin, os.Stdout, os.Stderr)
	return cmd.Execute()
}
