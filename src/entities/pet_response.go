package entities

type PetResponse struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Kind        string `json:"kind"`
	Breed       string `json:"breed"`
	BirthDay    string `json:"birth_day"`
	AdoptionDay string `json:"adoption_day"`
	PictureURL  string `json:"picture_url"`
}

type PetRoutineResponse struct {
	Kind  string `json:"kind"`
	Start string `json:"start"`
	End   string `json:"end"`
}
