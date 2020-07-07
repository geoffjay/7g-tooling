package model

type OneOnOne struct {
	BaseModelSeq
	//1-on-1 Template Name
	//1-on-1 Template ID
	//1-on-1 Creation Date
	//1-on-1 Scheduled Date
	//1-on-1 Completion Date
	Facilitator User
	TeamMember  User
	//Question 1
	//Facilitator Notes
	//Team Member Notes
	//Question 2
	//Facilitator Notes
	//Team Member Notes
	//Question 3
	//Facilitator Notes
	//Team Member Notes
	//Question 4
	//Facilitator Notes
	//Team Member Notes
	//Question 5
	//Facilitator Notes
	//Team Member Notes
}

type OneOnOneTemplate struct {
	BaseModelSeq
	Name        *string `gorm:"not null"`
	Description *string `gorm:"size:1024"`
	//1on1 Frequency
	//Question Name
	//Question Description `gorm:"size:1024"`
}
