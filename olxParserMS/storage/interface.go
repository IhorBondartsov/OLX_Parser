package storage

import "github.com/IhorBondartsov/OLX_Parser/olxParserMS/entities"

type Storage interface {
	CreateOrder(order entities.Order) (int, error)
    GetOrderByUserIDAndURL(uid int, url string) (entities.Order, error)

	CreateAdvertisement(a entities.Advertisement) error
	GetAdvertismentByOrderID(oid int) ([]entities.Advertisement, error)
}
