package events

import "github.com/jmoiron/sqlx"

type EventDB struct {
	db *sqlx.DB
}

func InitEventRepo(db_conn *sqlx.DB) EventDB {
	return EventDB{
		db: db_conn,
	}
}
