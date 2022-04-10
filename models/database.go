package models

import (
	"apiserver/configs"
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

func ConnectDB(dbConfig configs.DatabaseConfig) (*sql.DB, error) {

	oralInfo := fmt.Sprintf("%s/%s@%s:%d/%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Sid)
	db, err := sql.Open("godror", oralInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}
