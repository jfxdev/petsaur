package entities

import "time"

type Pet struct {
	UUID        string    `gorm:"primaryKey"`
	Name        string    `gorm:"NOT NULL"`
	Kind        string    `gorm:"NOT NULL"`
	Breed       string    `gorm:"type:varchar(100);Column:breed"`
	Birthday    time.Time `gorm:"Column:birthday"`
	Adoptionday time.Time `gorm:"Column:adoptionday"`
	Picture     string
}
