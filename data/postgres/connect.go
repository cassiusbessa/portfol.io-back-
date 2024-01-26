package postgres

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db     *sql.DB
	dbOnce sync.Once
)

func Connect() (*sql.DB, error) {
	var err error
	dbOnce.Do(func() {
		db, err = sql.Open("postgres", "postgres://postgres:postgres@orange-portfolio-postgres:5432/orange-portfolio?sslmode=disable")
		if err != nil {
			fmt.Println("Erro ao conectar ao banco de dados:", err)
		}
	})
	return db, err
}
