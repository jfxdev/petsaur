package pets

import (
	"github.com/jfxdev/petsaur/src/entities"
	"github.com/jfxdev/petsaur/src/infra/db"
)

func Create(pet entities.Pet) (id string, err error) {
	err = db.DB().Create(&pet).Error
	return pet.UUID, err
}

func List() (result []entities.Pet, err error) {
	tx := db.DB().Find(&result)
	err = tx.Error
	return
}
