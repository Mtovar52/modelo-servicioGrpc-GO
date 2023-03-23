package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Model
	Name        string `gorm:"column:name;type:varchar(255);not null"`
	NumDocument string `gorm:"column:num_document;type:varchar(45);not null"`
	Email       string `gorm:"uniqueIndex;column:email;type:varchar(255);not null"`
	Phone       string `gorm:"column:phone_contact;type:varchar(45);not null"`
	NickName    string `gorm:"column:nick_name;type:varchar(45);not null"`
	Conditions  bool   `gorm:"column:conditions;not null"`
	State       int8   `gorm:"column:state;type:tinyint(4);not null"`
	Password    string `gorm:"column:password;varchar(255);not null"`
}

func (m User) BeforeCreate(tx *gorm.DB) (err error) {
	m.Creationdate = time.Now()
	m.Updateddate = time.Now()
	return
}

func (m User) BeforeUpdate(tx *gorm.DB) (err error) {
	m.Updateddate = time.Now()
	return
}
