package entity

import "time"

type Model struct {
	ID           int64     `gorm:"column:id;primary_key;auto_increment;"`
	Creationdate time.Time `gorm:"autoCreateTime;column:creation_date;type:datetime;not null;"`
	Updateddate  time.Time `gorm:"autoUpdateTime;column:update_date;type:datetime;not null;"`
}
