package pet

import "github.com/jfxdev/petsaur/src/entities"

var _ *Service

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s Service) Routines() (response []entities.PetRoutineResponse) {
	return []entities.PetRoutineResponse{
		{
			Start: "0",
			End:   "1",
		},
		{
			Start: "2",
			End:   "3",
		},
	}
}
