package git

import (
	gg "github.com/go-git/go-git"
	"github.com/sirupsen/logrus"
)

func Clone(url, path string) {
	logrus.Debugf("cloning repository %s to %s\n", url, path)
	options := gg.CloneOptions{
		URL:               url,
		Auth:              nil,
		RemoteName:        "",
		ReferenceName:     "",
		SingleBranch:      false,
		NoCheckout:        false,
		Depth:             0,
		RecurseSubmodules: 0,
		Progress:          nil,
		Tags:              0,
	}
	// TODO: change path
	// gg.Clone(nil, nil, options)
	logrus.Debug(options)
}

func Checkout(branch string) {
	logrus.Debugf("checking out branch %s\n", branch)
}

func hasBranch(branch string) {
	logrus.Debugf("checking if branch %s exists\n", branch)
}
