package entities

type User struct {
	ID       int    `db:"id"`
	Login    string `db:"login"`
	Password string `db:"password"`
}

type Token struct {
	UserID         int    `db:"user_id"`
	Token          string `db:"token"`
	ExpirationTime int    `db:"expiration_time"`
}
