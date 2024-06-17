package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
)

func InitDB() (*sql.DB, error) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	db, err := sql.Open("postgres", "postgres://pranav:pranav@localhost:5432/mydb")
	if err != nil {
		logger.Error("Could not open database", err)
		return nil, err
	}

	stmt, err := db.Prepare("create table web_url(id serial primary key, url text not null);")
	if err != nil {
		logger.Error("Could not prepare statement", err)
		return nil, err
	}

	res, err := stmt.Exec()
	if err != nil {
		logger.Error("Could not execute statement", err)
		return nil, err
	}
	logger.Info("Executed", res)
	return db, nil
}
