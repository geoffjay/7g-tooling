package config

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/geoffjay/7g-tooling/internal/model"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

type Populate struct {
	Locations         []*model.Location         `mapstructure:"locations"`
	Departments       []*model.Department       `mapstructure:"departments"`
	Users             []*model.User             `mapstructure:"users"`
	Objectives        []*model.Objective        `mapstructure:"objectives"`
	OneOnOnes         []*model.OneOnOne         `mapstructure:"one-on-ones"`
	OneOnOneTemplates []*model.OneOnOneTemplate `mapstructure:"one-on-one-templates"`
	RecognitionBadges []*model.RecognitionBadge `mapstructure:"recognition-badges"`
	Recognitions      []*model.Recognition      `mapstructure:"recognitions"`
	CompetencyLevels  []*model.Level            `mapstructure:"competency-levels"`
	Competencies      []*model.Competency       `mapstructure:"competencies"`
	Roles             []*model.Role             `mapstructure:"roles"`
	RoleTemplates     []*model.RoleTemplate     `mapstructure:"role-templates"`
}

func LoadPopulate() (*Populate, error) {
	// TODO: tmpdir should be configurable
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	tmpDir := fmt.Sprintf("%s/.config/7g/tmp", home)
	file := fmt.Sprintf("%s/populate.yml", tmpDir)
	config := viper.New()
	config.SetConfigFile(file)
	if err = config.ReadInConfig(); err != nil {
		logrus.Error("Failed to read populate config:", err)
		return nil, err
	}

	var p Populate
	if err = config.Unmarshal(&p); err != nil {
		logrus.Error("Failed to populate data from file:", err)
		return nil, err
	}

	return &p, nil
}

func (p *Populate) Summary() map[string]int {
	return map[string]int{
		"locations":            len(p.Locations),
		"departments":          len(p.Departments),
		"users":                len(p.Users),
		"objectives":           len(p.Objectives),
		"one-on-ones":          len(p.OneOnOnes),
		"one-on-one-templates": len(p.OneOnOneTemplates),
		"recognition-badges":   len(p.RecognitionBadges),
		"recognitions":         len(p.Recognitions),
		"competency-levels":    len(p.CompetencyLevels),
		"competencies":         len(p.Competencies),
		"roles":                len(p.Roles),
		"role-templates":       len(p.RoleTemplates),
	}
}
