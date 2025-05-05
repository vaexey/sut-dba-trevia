package routes

import (
	"back/db"
)

type Api struct {
	Db *db.Database
}

func NewApi(db *db.Database) Api {
	return Api { Db : db }
}