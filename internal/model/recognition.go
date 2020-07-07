package model

type Recognition struct {
	BaseModelSeq
	Recipient User
	Provider  User
	Badge     RecognitionBadge
	Message   *string `gorm:"size:1024"`
}

type RecognitionBadge struct {
	BaseModelSeq
	Name        *string
	Description *string `gorm:"size:1024"`
}
