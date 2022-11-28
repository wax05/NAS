package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/wax05/NAS/dict"

	"errors"
)

type rtnVal []interface{}

var (
	errDBConnectionError = errors.New("can't database connect")
)

func dbquery(Query string, db sql.DB) (*sql.Rows, error) {
	rows, err := db.Query(Query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return rows, nil
}

func DbInit() (*sql.DB, error) {
	db, DbErr := sql.Open("mysql", "root:0000@tcp(localhost:3306)/nas")
	if DbErr != nil {
		return nil, errDBConnectionError
	}
	defer db.Close()
	return db, DbErr
}

func AllUserData(db sql.DB) (rtnVal, error) {
	rows, err := dbquery("SELECT * FROM userdata", db)
	if err != nil {
		return nil, err
	}
	var (
		Userid, UserName, PwHash, Class, UserPath string
	)
	RV := rtnVal{}
	for rows.Next() {
		err := rows.Scan(&Userid, &UserName, &PwHash, &Class, &UserPath)
		if err != nil {
			return nil, err
		}
		d := dict.Dictionary{
			"Id":       Userid,
			"UserName": UserName,
			"PwHash":   PwHash,
			"Class":    Class,
			"UserPath": UserPath,
		}
		RV = append(RV, d)
	}
	return RV, nil
}
