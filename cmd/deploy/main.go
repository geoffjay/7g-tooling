package deploy

import (
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var (
	branch string

	Command = &cobra.Command{
		Use:   "deploy",
		Short: "Work with branch deployments",
		Run:   deploy,
	}
)

func init() {
	Command.PersistentFlags().StringVarP(&branch, "branch", "b", "staging", "branch name")

	Command.AddCommand(initializeCmd)
}

func deploy(cmd *cobra.Command, args []string) {
	logrus.Debugf("running deployment for %s\n", branch)
}
