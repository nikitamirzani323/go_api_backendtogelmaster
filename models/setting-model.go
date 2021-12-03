package models

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/config"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/db"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/helpers"
)

type setting struct {
	StartMaintenance string `json:"maintenance_start"`
	EndMaintenance   string `json:"maintenance_end"`
}

func Fetch_setting() (helpers.Response, error) {
	var obj setting
	var arraobj []setting
	var res helpers.Response
	msg := "Error"
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
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	msg := "Failed"
	flag := false

	rows_update, err_update := con.PrepareContext(ctx, `
				UPDATE 
				`+config.DB_tbl_mst_setting+`  
				SET startmaintenance=?, endmaintenance=? 
				WHERE idversion =? 
			`)
	helpers.ErrorCheck(err_update)
	rec_comp, err_comp := rows_update.ExecContext(ctx, start, end, "1")
	helpers.ErrorCheck(err_comp)
	update_comp, err_comp := rec_comp.RowsAffected()
	helpers.ErrorCheck(err_comp)
	defer rows_update.Close()
	if update_comp > 0 {
		flag = true
		msg = "Success"
		log.Printf("Update %s Success\n", config.DB_tbl_mst_setting)
	} else {
		log.Printf("Update %s Failed\n", config.DB_tbl_mst_setting)
	}
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
