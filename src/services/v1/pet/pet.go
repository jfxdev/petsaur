package pet

import (
	"time"

	"github.com/google/uuid"
	"github.com/jfxdev/petsaur/src/entities"
	"github.com/jfxdev/petsaur/src/models/pets"
	"gopkg.in/validator.v2"
)

func (s Service) Create(payload entities.PetHandler) (id string, err error) {
	if err = validator.Validate(payload); err != nil {
		return
	}

	birthDay, err := time.Parse("2006-01-02", payload.BirthDay)
	if err != nil {
		return
	}
	adoptionDay, err := time.Parse("2006-01-02", payload.AdoptionDay)
	if err != nil {
		return
	}

	p := entities.Pet{
		UUID:        uuid.New().String(),
		Name:        payload.Name,
		Kind:        payload.Kind,
		Breed:       payload.Breed,
		Birthday:    birthDay,
		Adoptionday: adoptionDay,
		Picture:     "",
	}

	id, err = pets.Create(p)
	return
}

func (s Service) Get() (response []entities.PetHandler, err error) {
	result, err := pets.List()
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
