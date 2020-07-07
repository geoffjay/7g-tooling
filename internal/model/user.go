package model

type User struct {
	BaseModelSeq
	FirstName         *string
	LastName          *string
	Email             *string
	Phone             *string
	ManagerID         int
	ManagerName       *string
	State             bool
	Title             *string
	DepartmentsJoined *string
	LocationsJoined   *string
	RolesJoined       *string

	Departments []Department `gorm:"many2many:user_departments;association_autoupdate:false;association_autocreate:false"`
	Locations   []Location   `gorm:"many2many:user_locations;association_autoupdate:false;association_autocreate:false"`
	Roles       []Role       `gorm:"many2many:user_roles;association_autoupdate:false;association_autocreate:false"`
}
