package models

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/config"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/db"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/entities"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/helpers"
)

func Fetch_log() (helpers.Response, error) {
	var obj entities.Model_log
	var arraobj []entities.Model_log
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()

	sql_select := `SELECT 
			idlog, datetimelog, username, pagelog,  
			tipelog, noteafter, idcompany  
			FROM ` + config.DB_tbl_trx_log + ` 
			ORDER BY datetimelog DESC  LIMIT 300
		`
	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idlog_db                                                                        int
			datetimelog_db, username_db, pagelog_db, tipelog_db, noteafter_db, idcompany_db string
		)

		err = row.Scan(&idlog_db, &datetimelog_db, &username_db, &pagelog_db, &tipelog_db, &noteafter_db, &idcompany_db)
		helpers.ErrorCheck(err)

		obj.Log_id = idlog_db
		obj.Log_datetime = datetimelog_db
		obj.Log_company = idcompany_db
		obj.Log_username = username_db
		obj.Log_page = pagelog_db
		obj.Log_tipe = tipelog_db
		obj.Log_note = noteafter_db
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(render_page).String()

	return res, nil
}
