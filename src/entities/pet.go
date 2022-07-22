package entities

type Pet struct {
	UUID        string `gorm:"primaryKey"`
	Name        string `gorm:"NOT NULL"`
	Kind        string `gorm:"NOT NULL"`
	Breed       string `gorm:"type:varchar(100);Column:breed"`
	Birthday    string `gorm:"type:varchar(10);Column:birthday"`
	Adoptionday string `gorm:"type:varchar(10);Column:adoptionday"`
	Picture     string
}
