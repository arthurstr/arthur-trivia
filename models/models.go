package models

import "gorm.io/gorm"

type Fact struct {
	gorm.Model
	Question string 'json:"quastion" gorm:"text;not null;default:null'
	Answer string	'json:"answer" gorm:"text;not null;default:null'
}