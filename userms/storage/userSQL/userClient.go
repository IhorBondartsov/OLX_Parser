package mysql

import (
	"github.com/IhorBondartsov/OLX_Parser/userms/entities"
	"github.com/jmoiron/sqlx"
)

func NewUserMyClientMySQL(db *sqlx.DB) *myUserClientMySQL {
	return &myUserClientMySQL{db: db}
}

type myUserClientMySQL struct {
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
					id = ?;
`
)

func (c *myUserClientMySQL) Create(user entities.User) error {
	_, err := c.db.NamedExec(createUserStmt, user)
	return err
}

func (c *myUserClientMySQL) Update(user entities.User) error {
	_, err := c.db.NamedExec(updateUserStmtByID, user)
	return err
}

func (c *myUserClientMySQL) Delete(userID int) error {
	_, err := c.db.Query(deleteUserStmtByID, userID)
	return err
}

func (c *myUserClientMySQL) GetUserByLogin(login string) (entities.User, error) {
	var user entities.User
	err := c.db.Get(&user, getUserByID, login)
	return user, err
}
