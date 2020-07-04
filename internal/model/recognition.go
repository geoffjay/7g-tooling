package model

type Recognition struct {
	//Recognition ID
	//Recipient ID
	//Recipient Name
	//Recipient Department
	//Recipient Location
	//Provider ID
	//Provider Name
	//Provider Department
	//Provider Location
	//Badge Name
	//Recognition Message
}

type RecognitionBadge struct {
	Name        string `xlsx:"0"`
	Description string `xlsx:"1"`
}
