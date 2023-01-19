package db

import (
	// golang package
	"context"
	"database/sql"

	// external package
	"github.com/jmoiron/sqlx"
)

type SQLProvider interface {
	// Begin begins a transaction on master DB.
	Begin() (*sql.Tx, error)

	// BeginTx starts a transaction.
	//
	// The provided context is used until the transaction is committed or rolled back.
	// If the context is canceled, the sql package will roll back
	// the transaction. Tx.Commit will return an error if the context provided to
	// BeginTx is canceled.
	//
	// The provided TxOptions is optional and may be nil if defaults should be used.
	// If a non-default isolation level is used that the driver doesn't support,
	// an error will be returned.
	BeginTx(ctx context.Context, options *sql.TxOptions) (*sql.Tx, error)

	// BindNamed will do usual BindNamed by driverName param in db.
	// Please use this rather than BindNamed in GetMaster() or GetFollower() to make sure the rebind is correct, especially if you use newrelic
	BindNamed(query string, arg interface{}) (string, []interface{}, error)

	// ExecContext execs the query.
	// Any placeholder parameters are replaced with supplied args.
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)

	// GetContext does a QueryRow using the provided query and scans the resulting row to dest.
	// If dest is scannable, the result must only have one column. Otherwise, StructScan is used. Get will return sql.ErrNoRows like row.Scan would.
	// Any placeholder parameters are replaced with supplied args. An error is returned if the result set is empty.
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error

	// NamedExecContext executes a query using the provided query.
	//
	// Any named placeholder parameters are replaced with fields from arg.
	NamedExecContext(context context.Context, query string, arg interface{}) (sql.Result, error)

	// NamedQueryContext do named query on follower DB with context.
	// Any placeholder parameters are replaced with supplied args.
	NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error)

	// Rebind a query from the default bindtype (QUESTION) to the target bindtype.
	Rebind(query string) string

	// SelectContext executes a query using the provided query, and StructScans each row into dest, which must be a slice.
	// If the slice elements are scannable, then the result set must have only one column. Otherwise, StructScan is used. The *sql.Rows are closed automatically.
	// Any placeholder parameters are replaced with supplied args.
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type Repository struct {
	db SQLProvider
}

// New instantiates a new repository by given db pointer of SQLProvider.
//
// It returns pointer of Repository when successful.
// Otherwise, nil pointer of Repository will be returned.
func New(db SQLProvider) *Repository {
	return &Repository{
		db: db,
	}
}
