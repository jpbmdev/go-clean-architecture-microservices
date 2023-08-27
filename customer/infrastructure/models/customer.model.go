package models

import "gorm.io/gorm"

type CustomerModel struct {
	gorm.Model
	Name string `gorm:"not null"`
	Id   uint   `gorm:"primaryKey;autoIncrement"`
}
