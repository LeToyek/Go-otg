package db

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestRepository_GetCommonByID(t *testing.T) {
	tests := []struct {
		name    string
		args    int64
		want    Common
		wantErr error
		// mock    func(mock mockFields)
	}{
		{
			name: "Success get data",
			args: 1,
			want: Common{
				ID:      1,
				Message: "Halo gan!",
			},
			wantErr: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockDB, mockSQL, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				panic(err)
			}
			fmt.Println(mockSQL)
			repository := &Repository{
				db: sqlx.NewDb(mockDB, "pq"),
			}
			got, err := repository.GetCommonByID(test.args)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}
