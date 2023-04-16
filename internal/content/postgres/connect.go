package postgres

import (
	"animeList/internal/content"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/jackc/pgx/v5"
	"github.com/jmoiron/sqlx"
	"log"
)

// Кароч, тут мы подключаем базу данных и реализуем CRUD админа и favorites

type Database struct {
	conn     *sqlx.DB
	content  content.RepositoryContent
	favorite content.RepositoryFavorites
}

func NewDB() content.Database {
	return &Database{}
}

func (D *Database) Connect(url string) error {
	conn, err := sqlx.Connect("pgx", url)
	if err != nil {
		return err
	}

	if err := conn.Ping(); err != nil {
		return err
	}

	D.conn = conn
	log.Println("DB has successful pinging")
	return nil
}

func (D *Database) Close() error {
	return D.conn.Close()
}
