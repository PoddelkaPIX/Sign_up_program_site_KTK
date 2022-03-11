package db

import (
	"advanced_training_site_KTK/ServerSettings"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var Link *sql.DB

func Connect(cfg *ServerSettings.Setting) error {
	var e error
	Link, e = sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PgHost,
		cfg.PgPort,
		cfg.PgUser,
		cfg.PgPass,
		cfg.PgName,
	))
	if e != nil {
		fmt.Println("Не получилось подключиться к базе данных!")
		return e
	}

	e = Link.Ping()
	if e != nil {
		fmt.Println("Нет подключения к базе данных!")
		return e
	}

	return nil
}
