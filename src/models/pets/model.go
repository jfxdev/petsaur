package pets

import (
	"github.com/jfxdev/petsaur/src/entities"
	"github.com/jfxdev/petsaur/src/infra/db"
)

func CreatePet(pet entities.Pet) (err error) {
	err = db.DB().Create(&pet).Error
	return
}

func GetPets() (result []entities.Pet, err error) {
	tx := db.DB().Find(&result)
	err = tx.Error
	return
}
