package pet

import (
	"github.com/jfxdev/petsaur/src/entities"
	"github.com/jfxdev/petsaur/src/models/pets"
)

func (s Service) Create(payload entities.PetHandler) (err error) {
	return nil
}

func (s Service) Get() (response []entities.PetHandler, err error) {
	result, err := pets.GetPets()
	if err != nil {
		return
	}

	for _, r := range result {
		response = append(response, entities.PetHandler{
			UUID:  r.UUID,
			Name:  r.Name,
			Kind:  r.Kind,
			Breed: r.Breed,
		})
	}
	return
}

func (s Service) Validate(payload entities.PetHandler) error {
	return nil
}
