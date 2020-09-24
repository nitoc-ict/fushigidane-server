package rdbms

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var (
	Fushigidane *sql.DB
)

func InitGormClient() error {
	db, err := sql.Open("mysql", "root:root@/fushigidane")
	if err != nil {
		return errors.Wrap(err, "Failed open MySQL")
	}

	Fushigidane = db

	return nil
}
