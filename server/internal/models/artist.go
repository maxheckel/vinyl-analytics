package models

import "github.com/jinzhu/gorm"

type Artist struct {
	*gorm.Model
	Name string `gorm:"type:text"`
	Bio  string `gorm:"type:text"`
}
