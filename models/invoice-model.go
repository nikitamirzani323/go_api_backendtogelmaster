package models

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/config"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/db"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/helpers"
	"github.com/nleeper/goment"
)

type invoicehome struct {
	Idinvoice string `json:"invoice_id"`
	Company   string `json:"invoice_company"`
	Date      string `json:"invoice_date"`
	Name      string `json:"invoice_name"`
	Winlose   int    `json:"invoice_winlose"`
	Status    string `json:"invoice_status"`
	Statuscss string `json:"invoice_statuscss"`
}

func Fetch_invoice() (helpers.Response, error) {
	var obj invoicehome
	var arraobj []invoicehome
	var res helpers.Response
	msg := "Error"
	render_page := time.Now()
	con := db.CreateCon()
	ctx := context.Background()

	sql_periode := `SELECT 
			A.idcompinvoice, A.idcompany, B.nmcompany , A.datecompinvoice, A.nmcompinvoice, 
			A.winlosecomp, A.statuscompinvoice   
			FROM ` + config.DB_tbl_trx_company_invoice + ` as A 
			JOIN ` + config.DB_tbl_mst_company + ` as B ON B.idcompany = A.idcompany
			ORDER BY A.datecompinvoice DESC 
		`
	row, err := con.QueryContext(ctx, sql_periode)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idcompinvoice_db, winlosecomp_db                                                       int
			idcompany_db, datecompinvoice_db, nmcompany_db, nmcompinvoice_db, statuscompinvoice_db string
		)

		err = row.Scan(
			&idcompinvoice_db, &idcompany_db, &nmcompany_db, &datecompinvoice_db,
			&nmcompinvoice_db, &winlosecomp_db, &statuscompinvoice_db)
		helpers.ErrorCheck(err)

		status_css := ""
		switch statuscompinvoice_db {
		case "COMPLETED":
			status_css = config.STATUS_COMPLETE
		case "PROGRESS":
			status_css = config.STATUS_RUNNING
		}

		obj.Idinvoice = "INV_" + strconv.Itoa(idcompinvoice_db)
		obj.Company = nmcompany_db
		obj.Date = datecompinvoice_db
		obj.Name = nmcompinvoice_db
		obj.Winlose = winlosecomp_db
		obj.Status = statuscompinvoice_db
		obj.Statuscss = status_css
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
func Save_invoice(sData, master, periode string) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()
	msg := "Failed"
	flag := false

	if sData == "New" {
		year := tglnow.Format("YYYY")
		month := periode
		periodeinvoice := year + "-" + month
		nmcompinvoice := strings.ToUpper(tglnow.Format("MMMM")) + " " + year
		endday, _, _ := helpers.GetEndRangeDate(periode)
		start := tglnow.Format("YYYY") + "-" + month + "-" + "01"
		end := tglnow.Format("YYYY") + "-" + month + "-" + endday

		sql_select := `SELECT 
			idcompany    
			FROM ` + config.DB_tbl_mst_company + ` 
			WHERE statuscompany = 'ACTIVE'
		`
		row_select, err_select := con.QueryContext(ctx, sql_select)
		helpers.ErrorCheck(err_select)
		for row_select.Next() {
			var (
				idcompany_db string
			)

			err_select = row_select.Scan(&idcompany_db)
			helpers.ErrorCheck(err_select)
			winlose := _winlose(idcompany_db, start, end, 0)

			flag_insert := CheckDBTwoField(config.DB_tbl_trx_company_invoice, "idcompany", idcompany_db, "periodeinvoice", periodeinvoice)
			if !flag_insert {
				rows_insert, err_insert := con.PrepareContext(ctx, `
					INSERT INTO  
					`+config.DB_tbl_trx_company_invoice+` (
						idcompinvoice , idcompany, datecompinvoice, periodeinvoice, nmcompinvoice, winlosecomp, 
						statuscompinvoice, createcompinvoice, createdatecompinvoice 
					)VALUES( 
						?,?,?,?,?,?,
						?,?,?
					) 
				`)
				helpers.ErrorCheck(err_insert)

				field_table := config.DB_tbl_trx_company_invoice + tglnow.Format("YYYY")
				idrecord_counter := Get_counter(field_table)
				idrecord := tglnow.Format("YY") + tglnow.Format("MM") + tglnow.Format("DD") + strconv.Itoa(idrecord_counter)

				rec_comp, err_comp := rows_insert.ExecContext(ctx,
					idrecord,
					idcompany_db,
					tglnow.Format("YYYY-MM-DD"),
					periodeinvoice,
					nmcompinvoice,
					winlose,
					"PROGRESS",
					master,
					tglnow.Format("YYYY-MM-DD HH:mm:ss"))
				helpers.ErrorCheck(err_comp)
				insert, e := rec_comp.RowsAffected()
				helpers.ErrorCheck(e)
				defer rows_insert.Close()
				if insert > 0 {
					flag = true
					msg = "Success"
					log.Println("Data Berhasil di save")
				}
			} else {
				msg = "Duplicate Entry"
			}

		}
		defer row_select.Close()
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
func Save_invoicestatus(master, invoice, tipe string) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()
	msg := "Failed"
	flag := false

	temp_invoice := strings.Split(invoice, "_")
	idinvoice := temp_invoice[1]

	switch tipe {
	case "UPDATE-WINLOSE":
		winlose := 0
		sql_select := `SELECT 
			idcompany, periodeinvoice    
			FROM ` + config.DB_tbl_trx_company_invoice + ` 
			WHERE idcompinvoice =? 
		`
		row_select, err_select := con.QueryContext(ctx, sql_select, idinvoice)
		helpers.ErrorCheck(err_select)
		for row_select.Next() {
			var (
				idcompany_db, periodeinvoice_db string
			)

			err_select = row_select.Scan(&idcompany_db, &periodeinvoice_db)
			helpers.ErrorCheck(err_select)

			temp := strings.Split(periodeinvoice_db, "-")
			year := temp[0]
			month := temp[1]
			endday, _, _ := helpers.GetEndRangeDate(month)
			start := year + "-" + month + "-" + "01"
			end := year + "-" + month + "-" + endday
			winlose = _winlose(idcompany_db, start, end, 0)

		}
		defer row_select.Close()

		rows_update, err_update := con.PrepareContext(ctx, `
				UPDATE 
				`+config.DB_tbl_trx_company_invoice+`  
				SET winlosecomp=?,  
				updatecompinvoice=?, updatedatecompinvoice=? 
				WHERE idcompinvoice =? 
			`)
		helpers.ErrorCheck(err_update)

		rec_comp, err_comp := rows_update.ExecContext(ctx,
			winlose,
			master,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			idinvoice)
		helpers.ErrorCheck(err_comp)
		update_comp, err_comp := rec_comp.RowsAffected()
		helpers.ErrorCheck(err_comp)
		defer rows_update.Close()
		if update_comp > 0 {
			flag = true
			msg = "Success"
			log.Printf("Update %s WINLOSE Success : %s\n", config.DB_tbl_trx_company_invoice, idinvoice)
		} else {
			log.Printf("Update %s WINLOSE Failed \n", config.DB_tbl_trx_company_invoice)
		}
	case "UPDATE-STATUS":
		rows_update, err_update := con.PrepareContext(ctx, `
				UPDATE 
				`+config.DB_tbl_trx_company_invoice+`  
				SET statuscompinvoice=?,  
				updatecompinvoice=?, updatedatecompinvoice=? 
				WHERE idcompinvoice =? 
			`)
		helpers.ErrorCheck(err_update)

		rec_comp, err_comp := rows_update.ExecContext(ctx,
			"COMPLETED",
			master,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			idinvoice)
		helpers.ErrorCheck(err_comp)
		update_comp, err_comp := rec_comp.RowsAffected()
		helpers.ErrorCheck(err_comp)
		defer rows_update.Close()
		if update_comp > 0 {
			flag = true
			msg = "Success"
			log.Printf("Update %s STATUS Success : %s\n", config.DB_tbl_trx_company_invoice, idinvoice)
		} else {
			log.Printf("Update %s STATUS Failed \n", config.DB_tbl_trx_company_invoice)
		}
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
