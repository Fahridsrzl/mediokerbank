package repository

import "medioker-bank/model"

type ProfileRepository interface {
    CreateProfile(payload model.Profile, profile *model.Profile) error
    GetProfileByID(payload model.Profile, id string) (*model.Profile, error)
    UpdateProfile(payload model.Profile, id string, profile *model.Profile) error
    DeleteProfile(payload model.Profile, id string) error
}