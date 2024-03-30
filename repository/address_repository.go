package repository

import "medioker-bank/model"

type AddressRepository interface {
    CreateAddress(payload model.Address, address *model.Address) error
    GetAddressByID(payload model.Address, id string) (*model.Address, error)
    UpdateAddress(payload model.Address, id string, address *model.Address) error
    DeleteAddress(payload model.Address, id string) error
}
