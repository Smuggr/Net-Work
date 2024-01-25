package models

import (
	"gorm.io/gorm"
)


type User struct {
    gorm.Model
    Username        string `json:"username" gorm:"unique;not null"`
    Password        string `json:"-" gorm:"not null"`
    Identifier      uint   `json:"identifier" gorm:"uniqueIndex;autoIncrement"`
    PermissionLevel int    `json:"permission_level" gorm:"not null;default:1"`
}