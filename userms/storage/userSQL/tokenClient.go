package userSQL

import (
	"github.com/IhorBondartsov/OLX_Parser/userms/entities"
	"github.com/jmoiron/sqlx"
)

func NewTokenClientMySQL(db *sqlx.DB) *myTokenClientMySQL {
	return &myTokenClientMySQL{db: db}
}

type myTokenClientMySQL struct {
	db *sqlx.DB
}

const (
	createTokenStmt = `
				INSERT INTO
				refresh_token
				SET
					user_id             = :user_id,
					token               = :token,
					expiration_time     = :expiration_time;
`

	deleteTokenByUserIDStmt = `
				DELETE
				FROM
					refresh_token
				WHERE
					user_id = ?;
`

	getTokenByToken = `
				SELECT *
				FROM refresh_token WHERE
					token = ?
`

	getTokenByRange = `
				SELECT *
				FROM refresh_token WHERE
					expiration_time > ?
				AND
					expiration_time < ?				
`
)

func (c *myTokenClientMySQL) SetToken(token entities.Token) error {
	_, err := c.db.NamedExec(createTokenStmt, token)
	return err
}
func (c *myTokenClientMySQL) GetTokenByToken(token string) (entities.Token, error) {
	var tokenSt entities.Token
	err := c.db.Get(&tokenSt, getTokenByToken, token)
	return tokenSt, err
}
func (c *myTokenClientMySQL) DeleteToken(token entities.Token) error {
	_, err := c.db.Query(deleteTokenByUserIDStmt, token)
	return err
}
func (c *myTokenClientMySQL) GetTokenByRange(from, to int64) ([]entities.Token, error) {
	var tokenSt []entities.Token
	err := c.db.Get(&tokenSt, getTokenByRange, from, to)
	return tokenSt, err
}
