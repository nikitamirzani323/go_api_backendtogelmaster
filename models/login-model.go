package models

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/nikitamirzani323/go_api_backendtogelmaster/config"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/db"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/helpers"
)

func Login_Model(username, password string) (bool, error) {
	con := db.CreateCon()
	ctx := context.Background()

	var password_DB string

	sql_login := `
			SELECT 
			password  
			FROM ` + config.DB_tbl_mst_admin + ` 
			WHERE username  = ?
			AND statuslogin = "Y" 
	`

	rows := con.QueryRowContext(ctx, sql_login, username)
	switch e := rows.Scan(&password_DB); e {
	case sql.ErrNoRows:
		return false, errors.New("Username and Password Not Found")
	case nil:

	default:
		panic(e)
	}
	hashpass := helpers.HashPasswordMD5(password)
	log.Println("Password : " + hashpass)
	log.Println("Hash : " + password_DB)
	if hashpass != password_DB {
		return false, nil
	}

	return true, nil
}
