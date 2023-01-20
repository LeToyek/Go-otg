package db

import "database/sql"

type User struct {
	ID       string        `db:"id"`
	Username string        `db:"username"`
	Email    string        `db:"email"`
	Password string        `db:"password"`
	Address  string        `db:"addresss"`
	Age      sql.NullInt16 `db:"age"`
}
