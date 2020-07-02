package model

type User struct {
	ID                int    `xlsx:"0"`
	FirstName         string `xlsx:"1"`
	LastName          string `xlsx:"2"`
	Email             string `xlsx:"3"`
	Phone             string `xlsx:"4"`
	ManagerID         int    `xlsx:"5"`
	ManagerName       string `xlsx:"6"`
	State             bool   `xlsx:"7"`
	Title             string `xlsx:"8"`
	DepartmentsJoined string `xlsx:"9"`
	LocationsJoined   string `xlsx:"10"`
	RolesJoined       string `xlsx:"11"`

	Departments []Department
	Locations   []Location
	Roles       []Role
}
