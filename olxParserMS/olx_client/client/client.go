package client


import (
	"github.com/Sirupsen/logrus"

	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/entities"
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/olx_client/http_olx_client"
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/storage"
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/alarm_clock"
	"database/sql"
	"errors"
	"time"
	"github.com/alecthomas/chroma/lexers/q"
)

var log = logrus.New()


type OLXClient struct {
	HTTPClient http_olx_client.OlxHttpClient
	Storage    storage.Storage
	List alarm_clock.AlarmClock

	ResendQuit      chan struct{}
	AddToAlarmClock chan alarm_clock.Item
	RequestElem     chan int
}

func Start() {

}

// AddNewOrder - add new order to storage, and make first request to OLX, collect all
// advertisements from OLX and saved them to database, also added to AlarmClock
func (c *OLXClient) AddNewOrder(order entities.Order) error{
	var err error
	if !c.HasOrder(order){
		log.Error("[OLXClient][AddNewOrder]Duplicate order")
		return errors.New("Duplicate order")
	}

	order.ID, err = c.Storage.CreateOrder(order)
	if err != nil {
		return err
	}

	advrtsmnts := c.HTTPClient.GetHTMLPages(order.URL, order.PageLimit)
	for _, v := range advrtsmnts{
		v.OrderID = order.ID
		err = c.Storage.CreateAdvertisement(v)
		if err != nil {
			log.Errorf("[OLXClient][AddNewOrder] Database error %v", err)
		}
	}
	go c.sendItemToAlarmClock(time.Now().Add(time.Duration(order.Frequency) * time.Second).Unix(),order.ID)
	return nil
}

func (c *OLXClient) sendItemToAlarmClock(time int64, id int) {
	item := alarm_clock.Item{
		Time: time,
		Id:   id,
	}
	c.AddToAlarmClock <- item
}




func (c *OLXClient) HasOrder(order entities.Order) bool{
	_, err := c.Storage.GetOrderByUserIDAndURL(order.UserID, order.URL)
	if err == sql.ErrNoRows{
		return true
	}
	return false
}


func (c *OLXClient) FiltredAdvertisment(orderID int, advrtsmnts []entities.Advertisement) ([]entities.Advertisement, error) {
	savedAdvrtsmnts, err := c.Storage.GetAdvertismentByOrderID(orderID)
	if err == sql.ErrNoRows{
		log.Errorf("[OLXClient][FiltredAdvertisment] Database error %v", err)
		return  advrtsmnts, nil
	}
	if err != nil {
		log.Errorf("[OLXClient][FiltredAdvertisment] Database error %v", err)
		return nil, err
	}

	var newAdvrtsmnts []entities.Advertisement
	for _, v := range advrtsmnts {
		has := false
		for _, sv := range savedAdvrtsmnts {
			if sv.URL == v.URL {
				has = true
				break
			}
		}
		if !has {
			newAdvrtsmnts = append(newAdvrtsmnts, v)
		}
	}
	return newAdvrtsmnts, nil
}
