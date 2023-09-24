package models

import "gorm.io/gorm"

type CustomerModel struct {
	gorm.Model
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"not null"`
	Age       int    `gorm:"not null"`
	Score     int    `gorm:"not null"`
	Id        uint   `gorm:"primaryKey;autoIncrement"`
}
