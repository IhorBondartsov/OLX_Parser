package storage

import "github.com/IhorBondartsov/OLX_Parser/userms/entities"

type Storage interface {
	Create(user entities.User) error
	Update(user entities.User) error
	Delete(userID int) error
	GetUserByLogin(login string) (entities.User, error)
}

type RefreshToken interface {
	SetToken(token entities.Token) error
	GetTokenByToken(token string) (entities.Token, error)
	DeleteToken(token entities.Token) error
}
