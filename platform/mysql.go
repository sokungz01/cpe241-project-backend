package platform

import (
	_ "database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Mysql struct {
	*sqlx.DB
}

func NewSql(dataSource string) (*Mysql, error) {
	db, err := sqlx.Connect("mysql", dataSource)

	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &Mysql{db}, nil
}
