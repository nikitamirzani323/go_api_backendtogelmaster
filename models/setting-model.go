package models

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/config"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/db"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/entities"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/helpers"
)

func Fetch_setting() (helpers.Response, error) {
	var obj entities.Model_setting
	var arraobj []entities.Model_setting
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	flag := false
	var startmaintenance string = ""
	var endmaintenance string = ""
	sql_select := `SELECT 
			startmaintenance , endmaintenance 
			FROM ` + config.DB_tbl_mst_setting + ` 
			WHERE idversion = ? 
		`
	row := con.QueryRowContext(ctx, sql_select, "1")
	switch e := row.Scan(&startmaintenance, &endmaintenance); e {
	case sql.ErrNoRows:
		msg = "No Records"
	case nil:
		flag = true
	default:
		panic(e)
	}

	if startmaintenance == "00:00:00" {
		startmaintenance = ""
	}
	if endmaintenance == "00:00:00" {
		endmaintenance = ""
	}

	obj.StartMaintenance = startmaintenance
	obj.EndMaintenance = endmaintenance
	arraobj = append(arraobj, obj)

	if flag {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = arraobj
		res.Time = time.Since(render_page).String()
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = msg
		res.Record = nil
		res.Time = time.Since(render_page).String()
	}

	return res, nil
}
func Save_setting(start, end string) (helpers.Response, error) {
	var res helpers.Response
	render_page := time.Now()
	msg := "Failed"
	flag := false
	sql_update := `
			UPDATE 
			` + config.DB_tbl_mst_setting + `  
			SET startmaintenance=?, endmaintenance=? 
			WHERE idversion =? 
		`

	flag_update, msg_update := Exec_SQL(sql_update, config.DB_tbl_mst_setting, "UPDATE", start, end, "1")
	if flag_update {
		flag = true
		msg = "Succes"
		log.Println(msg_update)
	} else {
		log.Println(msg_update)
	}
	log.Printf("%s - %s - %t", start, end, flag_update)
	if flag {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = nil
		res.Time = time.Since(render_page).String()
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = msg
		res.Record = nil
		res.Time = time.Since(render_page).String()
	}

	return res, nil
}
