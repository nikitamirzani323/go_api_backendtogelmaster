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
	"github.com/nikitamirzani323/go_api_backendtogelmaster/entities"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/helpers"
	"github.com/nleeper/goment"
)

func Fetch_invoice() (helpers.Response, error) {
	var obj entities.Model_invoicehome
	var arraobj []entities.Model_invoicehome
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

	if sData == "New" {
		year := tglnow.Format("YYYY")
		month := periode
		periodeinvoice := year + "-" + month
		endday, _, _, nmmonth := helpers.GetEndRangeDate(periode)
		nmcompinvoice := nmmonth + " " + year
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
				sql_insert := `
					INSERT INTO  
					` + config.DB_tbl_trx_company_invoice + ` (
						idcompinvoice , idcompany, datecompinvoice, periodeinvoice, nmcompinvoice, winlosecomp, 
						statuscompinvoice, createcompinvoice, createdatecompinvoice 
					)VALUES( 
						?,?,?,?,?,?,
						?,?,?
					) 
				`
				field_table := config.DB_tbl_trx_company_invoice + tglnow.Format("YYYY")
				idrecord_counter := Get_counter(field_table)
				idrecord := tglnow.Format("YY") + tglnow.Format("MM") + tglnow.Format("DD") + strconv.Itoa(idrecord_counter)
				flag_insert, msg_insert := Exec_SQL(sql_insert, config.DB_tbl_trx_company_invoice, "INSERT",
					idrecord, idcompany_db,
					tglnow.Format("YYYY-MM-DD"),
					periodeinvoice,
					nmcompinvoice,
					winlose,
					"PROGRESS",
					master, tglnow.Format("YYYY-MM-DD HH:mm:ss"))
				if flag_insert {
					msg = "Succes"
					log.Println(msg_insert)
				} else {
					log.Println(msg_insert)
				}
			} else {
				msg = "Duplicate Entry"
			}

		}
		defer row_select.Close()
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Save_invoicestatus(master, invoice, tipe string) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()
	msg := "Failed"

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
			endday, _, _, _ := helpers.GetEndRangeDate(month)
			start := year + "-" + month + "-" + "01"
			end := year + "-" + month + "-" + endday
			log.Printf("%s - %s", start, end)
			winlose = _winlose(idcompany_db, start, end, 0)

		}
		defer row_select.Close()
		sql_update := `
			UPDATE 
			` + config.DB_tbl_trx_company_invoice + `  
			SET winlosecomp=?,  
			updatecompinvoice=?, updatedatecompinvoice=? 
			WHERE idcompinvoice =? 
		`
		flag_update, msg_update := Exec_SQL(sql_update, config.DB_tbl_trx_company_invoice, "UPDATE",
			winlose, master, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idinvoice)

		if flag_update {
			msg = "Succes"
			log.Println(msg_update)
		} else {
			log.Println(msg_update)
		}
	case "UPDATE-STATUS":
		sql_update := `
			UPDATE 
			` + config.DB_tbl_trx_company_invoice + `  
			SET statuscompinvoice=?,  
			updatecompinvoice=?, updatedatecompinvoice=? 
			WHERE idcompinvoice =?  
		`
		flag_update, msg_update := Exec_SQL(sql_update, config.DB_tbl_trx_company_invoice, "UPDATE",
			"COMPLETED",
			master,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			idinvoice)

		if flag_update {
			msg = "Succes"
			log.Println(msg_update)
		} else {
			log.Println(msg_update)
		}
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
