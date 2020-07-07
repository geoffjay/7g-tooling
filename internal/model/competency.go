package model

type Competency struct {
	BaseModelSeq
	Title       *string
	Description *string `gorm:"size:1024"`
	//Core Competency
	Levels []Level `gorm:"many2many:competency_levels;association_autoupdate:false;association_autocreate:false"`
}

type Level struct {
	BaseModelSeq
	Name        *string
	Description *string `gorm:"size:1024"`
}
