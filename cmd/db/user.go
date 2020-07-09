package db

import (
	gcontext "github.com/geoffjay/7g-tooling/internal/context"
	"github.com/geoffjay/7g-tooling/internal/model"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	userCmd = &cobra.Command{
		Use:   "user",
		Short: "Read users",
		Run:   user,
	}
)

func init() {
}

func user(cmd *cobra.Command, args []string) {
	config, err := gcontext.LoadConfig(".")
	if err != nil {
		logrus.Fatal(err)
	}

	db, err := gcontext.Connect(config)
	if err != nil {
		logrus.Fatal("Failed to connect to database:", err)
	}

	var users []model.User
	db.Model(&model.User{}).Find(&users)
	logrus.Infof("%+v", users)

	_ = gcontext.Disconnect(db)
}
