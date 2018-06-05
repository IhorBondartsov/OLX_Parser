package client

import (
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/entities"
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/olx_client/http_olx_client"
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/storage"
)

type OLXClient struct {
	HTTPClient http_olx_client.OlxHttpClient
	Storage    storage.Storage
}

func Start() {

}

func (c *OLXClient) FiltredAdvertisment(orderID int, advrtsmnts []entities.Advertisement) ([]entities.Advertisement, error) {
	savedAdvrtsmnts, err := c.Storage.GetAdvertismentByOrderID(orderID)
	if err != nil {
		return nil, err
	}
	var newAdvrtsmnts []entities.Advertisement
	for _, v := range advrtsmnts {
		for _, sv := range savedAdvrtsmnts {
			if sv.URL == v.URL {
				continue
			}
		}
		newAdvrtsmnts = append(newAdvrtsmnts, v)
	}
	return newAdvrtsmnts, nil
}
