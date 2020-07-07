package model

type Objective struct {
	BaseModelSeq
	Name        *string `gorm:"not null"`
	Type        int
	Description *string `gorm:"size:1024"`
	//Due Date
	//Creation Date
	//Open/Closed
	//Privacy
	//Objective Department
	//Labels
	//Parent Objective ID(s)
	//Child Objective ID(s)
	//Child Objective Weight(s)
	//Creator name
	//Creator ID
	//Owner name(s)
	//Owner ID(s)
	//Stakeholder name(s)
	//Stakeholder ID(s)
	//Follower name(s)
	//Follower ID(s)
	//KR name(s)
	//KR Unit(s)
	//KR Weight(s)
	//KR Target Value(s)
	//KR Start Value(s)
}
