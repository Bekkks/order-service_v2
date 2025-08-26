package http

import "crudl/internal/usecase"

type Handlers struct{
	profileService *usecase.Profile
}

func New(profile *usecase.Profile) *Handlers{
	return &Handlers{
		profileService: profile,
	}
}