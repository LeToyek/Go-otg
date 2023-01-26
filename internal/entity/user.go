package entity

import "database/sql"

type User struct {
	ID       string
	Username string
	Email    string
	Password string
	Address  string
	Age      sql.NullInt16
}
