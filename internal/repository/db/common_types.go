package db

type Common struct {
	ID      int64  `db:"id"`
	Message string `db:"message"`
}
