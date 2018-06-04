package olxstorage

import (
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/entities"
	"github.com/jmoiron/sqlx"
)

func NewStorage(db *sqlx.DB) *parserStorage {
	return &parserStorage{db: db}
}

type parserStorage struct {
	db *sqlx.DB
}

const (
	createOrderStmt = `
				INSERT INTO
					order
				SET
					user_id = :user_id
					url = :url
					page_limit = :page_limit
					delivery_method = :delivery_method
					expiration_time = :expiration_time
					frequency = :frequency
`
	updateUserStmtByID = `
				UPDATE
					user
				SET
					login             = :login,
					password          = :password,
				WHERE
					id = :id;
`
	deleteUserStmtByID = `
				DELETE
				FROM
					user
				WHERE
					id = ?;
`

	getUserByLogin = `
				SELECT *
				FROM user WHERE
					login = ?;
`
	// table advertisements
	createAdvertisementsStmt = `
				INSERT INTO
					advertisements
				SET
					order_id = :order_id
					title = :title
					url = :url
					created_at = :created_at
`

	getAdvertismentByOrderIDStmt = `
				SELECT *
				FROM advertisements WHERE
					order_id = ?;
`
)

func (c *parserStorage) CreateOrder(order entities.Order) error {
	_, err := c.db.NamedExec(createOrderStmt, order)
	return err
}
func (c *parserStorage) CreateAdvertisement(a entities.Advertisement) error {
	_, err := c.db.NamedExec(createAdvertisementsStmt, a)
	return err
}
func (c *parserStorage) GetAdvertismentByOrderID(oid int) ([]entities.Advertisement, error) {
	var advs []entities.Advertisement
	err := c.db.Get(&advs, getAdvertismentByOrderIDStmt, oid)
	return advs, err
}
