package gql

import (
	"github.com/geoffjay/7g-tooling/internal/client"
	gcontext "github.com/geoffjay/7g-tooling/internal/context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	email     string
	firstName string
	lastName  string
	position  string
	username  string

	createUserProfileCmd = &cobra.Command{
		Use:   "create-user-profile",
		Short: "Create a new user profile",
		Run:   createUserProfile,
	}
)

func init() {
	createUserProfileCmd.PersistentFlags().StringVarP(&email, "email", "", "", "email address")
	createUserProfileCmd.PersistentFlags().StringVarP(&firstName, "first-name", "", "", "first name")
	createUserProfileCmd.PersistentFlags().StringVarP(&lastName, "last-name", "", "", "last name")
	createUserProfileCmd.PersistentFlags().StringVarP(&position, "position", "", "", "position")
	createUserProfileCmd.PersistentFlags().StringVarP(&username, "username", "", "", "username")
}

func createUserProfile(cmd *cobra.Command, args []string) {
	logrus.Debug("gql > get-objective")

	config, err := gcontext.LoadConfig(".")
	if err != nil {
		logrus.Fatal(err)
	}

	// TODO: only add var when != ""
	// TODO: make username required field
	userProfile := map[string]interface{}{
		"email":     email,
		"firstName": firstName,
		"lastName":  lastName,
		"position":  position,
		"username":  username,
	}
	variables := map[string]interface{}{
		"userProfile": userProfile,
	}
	client.Query("create-user-profile", variables, config)
}
