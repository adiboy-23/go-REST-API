package sqlite

import (
	"database/sql"

	"github.com/adiboy-23/go-REST-API/internal/config"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) {

}
