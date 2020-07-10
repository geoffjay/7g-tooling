package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type User struct {
	gorm.Model
	SgID      int
	FirstName string  `mapstructure:"first-name"`
	LastName  string  `mapstructure:"last-name"`
	Email     *string `gorm:"unique;not null"`
	Phone     string
	// Creates empty record
	//Managers          []*User `gorm:"many2many:user_managers;association_jointable_foreignkey:manager_email"`
	State             string
	Title             string
	DepartmentsJoined string `gorm:"-"`
	LocationsJoined   string `gorm:"-"`
	RolesJoined       string `gorm:"-"`

	Departments []*Department `gorm:"many2many:user_departments;association_autoupdate:false;association_autocreate:false"`
	Locations   []*Location   `gorm:"many2many:user_locations;association_autoupdate:false;association_autocreate:false"`
	Roles       []*Role       `gorm:"many2many:user_roles;association_autoupdate:false;association_autocreate:false"`
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
	if store.HasUserWithEmail(*user.Email) {
		err = store.db.Update(user).Error
	} else {
		if err = store.db.Create(user).Error; err != nil {
			logrus.Panic(err)
		}
	}
	return
}

func (store *UserStore) Flush() (err error) {
	return store.db.Model(&User{}).Delete(&User{}).Error
}

func (store *UserStore) GetUserByEmail(email string) (user User, err error) {
	err = store.db.Model(&User{}).Where("email = ?", email).First(&user).Error
	return
}

func (store *UserStore) HasUserWithEmail(email string) bool {
	var users []User
	err := store.db.Model(&User{}).Where("email LIKE ?", email).Find(&users).Error
	if gorm.IsRecordNotFoundError(err) {
		return false
	}
	return true
}
