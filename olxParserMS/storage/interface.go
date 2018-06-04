package storage

import "github.com/IhorBondartsov/OLX_Parser/olxParserMS/entities"

type Storage interface {
	CreateOrder(entities.Order) error
	CreateAdvertisement(a entities.Advertisement) error
}
