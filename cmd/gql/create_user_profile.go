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
	createUserProfileCmd.PersistentFlags().StringVar(&email, "email", "", "email address")
	createUserProfileCmd.PersistentFlags().StringVar(&firstName, "first-name", "", "first name")
	createUserProfileCmd.PersistentFlags().StringVar(&lastName, "last-name", "", "last name")
	createUserProfileCmd.PersistentFlags().StringVar(&position, "position", "", "position")
	createUserProfileCmd.PersistentFlags().StringVar(&username, "username", "", "username")
}

func createUserProfile(cmd *cobra.Command, args []string) {
	logrus.Debug("gql > create-user-profile")

	config, err := gcontext.LoadConfig(".")
	if err != nil {
		logrus.Fatal(err)
	}

	if username == "" {
		logrus.Panic("A value for `username' is required")
	}

	// TODO: only add var when != ""
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
