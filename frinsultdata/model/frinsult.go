package model

import "github.com/jinzhu/gorm"

//Frinsult is an insult model
type Frinsult struct {
	gorm.Model
	Text  string
	Score int
}
