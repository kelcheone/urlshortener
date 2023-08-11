package storage

import "database/sql"

type Storage interface {
	Connect()
	GetDb() *sql.DB
}
