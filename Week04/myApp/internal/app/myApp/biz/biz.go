package biz

import "database/sql"

type Biz struct {
	DB *sql.DB
}

func NewBiz(db *sql.DB) Biz {
	return Biz{DB:db}
}