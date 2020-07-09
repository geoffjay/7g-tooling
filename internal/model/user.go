package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type User struct {
	gorm.Model
	SgID              int
	FirstName         *string
	LastName          *string
	Email             *string
	Phone             *string
	ManagerID         int
	ManagerName       *string
	State             *string
	Title             *string
	DepartmentsJoined *string
	LocationsJoined   *string
	RolesJoined       *string

	Departments []Department `gorm:"many2many:user_departments;association_autoupdate:false;association_autocreate:false"`
	Locations   []Location   `gorm:"many2many:user_locations;association_autoupdate:false;association_autocreate:false"`
	Roles       []Role       `gorm:"many2many:user_roles;association_autoupdate:false;association_autocreate:false"`
}

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (store *UserStore) Save(user *User) (err error) {
	if err = store.db.Create(user).Error; err != nil {
		logrus.Panic(err)
	}
	return
}

func (store *UserStore) Flush() (err error) {
	return store.db.Model(&User{}).Delete(&User{}).Error
}
