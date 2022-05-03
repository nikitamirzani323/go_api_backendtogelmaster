package models

import (
	"context"
	"database/sql"
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
	msg := "Data Not Found"
	render_page := time.Now()
	con := db.CreateCon()
	ctx := context.Background()

	sql_periode := `SELECT 
			A.idcompinvoice, A.idcompany, B.nmcompany , A.datecompinvoice, A.nmcompinvoice, 
			A.winlosecomp, A.statuscompinvoice   
			FROM ` + config.DB_tbl_trx_company_invoice + ` as A 
			JOIN ` + config.DB_tbl_mst_company + ` as B ON B.idcompany = A.idcompany
			ORDER BY A.periodeinvoice DESC 
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
		obj.Total_pasaran = _invoicepasaran_count(strconv.Itoa(idcompinvoice_db))
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
func Fetch_invoicedetail(invoice string) (helpers.Response, error) {
	var obj entities.Model_invoicehomedetail
	var arraobj []entities.Model_invoicehomedetail
	var res helpers.Response
	msg := "Data Not Found"
	render_page := time.Now()
	con := db.CreateCon()
	ctx := context.Background()

	temp_invoice := strings.Split(invoice, "_")
	idinvoice := temp_invoice[1]

	sql_select := `SELECT 
			A.idcompinvoicedetail , A.winlosecomppasaran, C.nmpasarantogel, 
			A.createcompinvoicedetail, A.createdatecompinvoicedetail, A.updatecompinvoicedetail, A.updatedatecompinvoicedetail 
			FROM ` + config.DB_tbl_trx_company_invoice_detail + ` as A 
			JOIN ` + config.DB_tbl_mst_company_game_pasaran + ` as B ON B.idcomppasaran = A.idcomppasaran 
			JOIN ` + config.DB_tbl_mst_pasaran_togel + ` as C ON C.idpasarantogel = B.idpasarantogel 
			WHERE A.idcompinvoice = ? 
			ORDER BY A.winlosecomppasaran DESC    
		`
	row, err := con.QueryContext(ctx, sql_select, idinvoice)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idcompinvoicedetail_db, winlosecomppasaran_db                                                                          int
			nmpasarantogel_db                                                                                                      string
			createcompinvoicedetail_db, createdatecompinvoicedetail_db, updatecompinvoicedetail_db, updatedatecompinvoicedetail_db string
		)

		err = row.Scan(
			&idcompinvoicedetail_db, &winlosecomppasaran_db, &nmpasarantogel_db,
			&createcompinvoicedetail_db, &createdatecompinvoicedetail_db, &updatecompinvoicedetail_db, &updatedatecompinvoicedetail_db)
		helpers.ErrorCheck(err)

		create := createcompinvoicedetail_db + " , " + createdatecompinvoicedetail_db
		update := ""

		if updatecompinvoicedetail_db != "" {
			update = updatecompinvoicedetail_db + " , " + updatedatecompinvoicedetail_db
		}

		obj.Idinvoicedetail = strconv.Itoa(idcompinvoicedetail_db)
		obj.Pasaran = nmpasarantogel_db
		obj.Winlose = winlosecomppasaran_db
		obj.Create = create
		obj.Update = update
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
func Save_company_listpasaran(master, invoice string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()

	log.Println("Invoice URL:", invoice)
	temp_invoice := strings.Split(invoice, "_")
	idinvoice := temp_invoice[1]
	company, periode := _invoice_id(idinvoice)
	data_invoice := strings.Split(periode, "-")
	year_invoice := data_invoice[0]
	month_invoice := data_invoice[1]
	log.Println("Invoice DB:", idinvoice)
	log.Println("year_invoice :", year_invoice)
	log.Println("month_invoice :", month_invoice)

	tglnow, _ := goment.New(year_invoice + "-" + month_invoice + "-01")
	endday, _, _, _ := helpers.GetEndRangeDate(tglnow.Format("MM"))
	start := tglnow.Format("YYYY-MM") + "-" + "01"
	end := tglnow.Format("YYYY-MM") + "-" + endday

	sql_periode := `SELECT 
			idcomppasaran 
			FROM ` + config.DB_tbl_mst_company_game_pasaran + ` 
			WHERE idcompany = ? 
			AND statuspasaranactive = "Y"  
		`

	row, err := con.QueryContext(ctx, sql_periode, company)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idcomppasaran_db int
		)

		err = row.Scan(&idcomppasaran_db)
		helpers.ErrorCheck(err)

		flag_insert := CheckDBTwoField(config.DB_tbl_trx_company_invoice_detail, "idcompinvoice", idinvoice, "idcomppasaran", strconv.Itoa(idcomppasaran_db))
		if !flag_insert {
			sql_insert := `
				INSERT INTO  
				` + config.DB_tbl_trx_company_invoice_detail + ` (
					idcompinvoicedetail , idcompinvoice, idcomppasaran, winlosecomppasaran, 
					createcompinvoicedetail, createdatecompinvoicedetail
				)VALUES( 
					?,?,?,?,
					?,?
				) 
			`
			winlose := _winlose(company, start, end, idcomppasaran_db)
			field_table := config.DB_tbl_trx_company_invoice_detail + tglnow.Format("YYYY")
			idrecord_counter := Get_counter(field_table)
			idrecord := tglnow.Format("YY") + tglnow.Format("MM") + tglnow.Format("DD") + strconv.Itoa(idrecord_counter)
			flag_insert, msg_insert := Exec_SQL(sql_insert, config.DB_tbl_trx_company_invoice_detail, "INSERT",
				idrecord, idinvoice, idcomppasaran_db, winlose,
				master, tglnow.Format("YYYY-MM-DD HH:mm:ss"))
			if flag_insert {
				msg = "Succes"
				log.Println(msg_insert)
			} else {
				log.Println(msg_insert)
			}
		}
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Delete_company_listpasaran(master, invoice string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	render_page := time.Now()

	log.Println("Invoice URL:", invoice)
	temp_invoice := strings.Split(invoice, "_")
	idinvoice := temp_invoice[1]

	sql_delete := `
		DELETE FROM
		` + config.DB_tbl_trx_company_invoice_detail + ` 
		WHERE idcompinvoice=? 
	`

	flag_delete, msg_delete := Exec_SQL(sql_delete, config.DB_tbl_trx_company_invoice_detail, "DELETE", idinvoice)
	if flag_delete {
		msg = "Succes"
		log.Println(msg_delete)
	} else {
		log.Println(msg_delete)
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func _invoice_id(invoice string) (string, string) {
	con := db.CreateCon()
	ctx := context.Background()
	idcompany := ""
	periodeinvoice := ""
	sql_select := `SELECT 
		idcompany, periodeinvoice 
		FROM ` + config.DB_tbl_trx_company_invoice + `  
		WHERE idcompinvoice = ? 
	`
	var (
		idcompany_db, periodeinvoice_db string
	)
	rows := con.QueryRowContext(ctx, sql_select, invoice)
	switch err := rows.Scan(&idcompany_db, &periodeinvoice_db); err {
	case sql.ErrNoRows:

	case nil:
		idcompany = idcompany_db
		periodeinvoice = periodeinvoice_db
	default:
		helpers.ErrorCheck(err)
	}
	return idcompany, periodeinvoice
}
func _invoicepasaran_count(invoice string) int {
	con := db.CreateCon()
	ctx := context.Background()
	result := 0
	sql_select := `SELECT 
		COUNT(idcompinvoicedetail) as total
		FROM ` + config.DB_tbl_trx_company_invoice_detail + `  
		WHERE idcompinvoice = ? 
	`
	var (
		total_pasaran_db int
	)
	rows := con.QueryRowContext(ctx, sql_select, invoice)
	switch err := rows.Scan(&total_pasaran_db); err {
	case sql.ErrNoRows:

	case nil:
		result = total_pasaran_db
	default:
		helpers.ErrorCheck(err)
	}
	return result
}
