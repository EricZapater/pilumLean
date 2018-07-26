package app

import(
	"database/sql"
	"github.com/gorilla/mux"
	_"github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB *sql.DB
}