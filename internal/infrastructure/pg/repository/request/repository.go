package request

import (
	trmgr "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db       *sqlx.DB
	txGetter *trmgr.CtxGetter
}

func New(db *sqlx.DB, txGetter *trmgr.CtxGetter) *Repository {
	return &Repository{
		db:       db,
		txGetter: txGetter,
	}
}
