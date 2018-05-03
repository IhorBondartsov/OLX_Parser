package mysql

import (
	"github.com/IhorBondartsov/OLX_Parser/userms/entities"
	"github.com/jmoiron/sqlx"
)

func NewMyClientMySQL(db *sqlx.DB) *myClientMySQL {
	return &myClientMySQL{db: db}
}

type myClientMySQL struct {
	db *sqlx.DB
}

const (
	createUserStmt = `
				INSERT INTO
					user
				SET
					login             = :login,
					password          = :password;
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

	getUserByID = `
				SELECT *
				FROM user WHERE
					id = ?
`
)

func (c *myClientMySQL) Create(user entities.User) error {

}
func (c *myClientMySQL) Update(user entities.User) error {

}
func (c *myClientMySQL) Delete(userID int) error {

}
func (c *myClientMySQL) GetUserByLogin(login string) (entities.User, error) {

}
