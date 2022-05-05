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

func Fetch_dashboardwinlose(year string) (helpers.Response, error) {
	var obj entities.Model_dashboardwinlose_parent
	var arraobj []entities.Model_dashboardwinlose_parent
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()

	//COMPANY
	sql_select := `SELECT 
		idcompany, nmcompany
		FROM ` + config.DB_tbl_mst_company + ` 
		WHERE statuscompany = "ACTIVE" 
		ORDER BY nmcompany ASC 
	`
	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idcompany_db, nmcompany_db string
		)

		err = row.Scan(&idcompany_db, &nmcompany_db)
		helpers.ErrorCheck(err)
		var objdetail entities.Model_dashboardwinlose_child
		var arraobjdetail []entities.Model_dashboardwinlose_child
		for i := 1; i < 13; i++ {
			periode := ""
			switch i {
			case 1:
				periode = year + "-01"
			case 2:
				periode = year + "-02"
			case 3:
				periode = year + "-03"
			case 4:
				periode = year + "-04"
			case 5:
				periode = year + "-05"
			case 6:
				periode = year + "-06"
			case 7:
				periode = year + "-07"
			case 8:
				periode = year + "-08"
			case 9:
				periode = year + "-09"
			case 10:
				periode = year + "-10"
			case 11:
				periode = year + "-11"
			case 12:
				periode = year + "-12"
			}

			var winlose int = _invoicewinlose_id(idcompany_db, year, periode)
			objdetail.Dashboardwinlose_winlose = winlose
			arraobjdetail = append(arraobjdetail, objdetail)
			msg = "Success"
		}

		obj.Dashboardwinlose_nmcompany = nmcompany_db
		obj.Dashboardwinlose_detail = arraobjdetail
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
func Fetch_dashboardwinlosebycompany(company, year string) (helpers.Response, error) {
	var obj entities.Model_companypasaran_parent
	var arraobj []entities.Model_companypasaran_parent
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()

	//COMPANY
	sql_select := `SELECT 
		A.idcomppasaran, A.idcompany, C.nmpasarantogel 
		FROM ` + config.DB_tbl_mst_company_game_pasaran + ` as A 
		JOIN ` + config.DB_tbl_mst_company + ` as B ON B.idcompany = A.idcompany  
		JOIN ` + config.DB_tbl_mst_pasaran_togel + ` as C ON C.idpasarantogel  = A.idpasarantogel  
		WHERE B.statuscompany = "ACTIVE" 
		AND A.idcompany = ?  
		ORDER BY A.displaypasaran ASC 
	`
	row, err := con.QueryContext(ctx, sql_select, company)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idcomppasaran_db                int
			idcompany_db, nmpasarantogel_db string
		)

		err = row.Scan(&idcomppasaran_db, &idcompany_db, &nmpasarantogel_db)
		helpers.ErrorCheck(err)
		var objdetail entities.Model_companypasaran_child
		var arraobjdetail []entities.Model_companypasaran_child
		for i := 1; i < 13; i++ {
			periode := ""
			idinvoice := 0
			switch i {
			case 1:
				periode = year + "-01"
			case 2:
				periode = year + "-02"
			case 3:
				periode = year + "-03"
			case 4:
				periode = year + "-04"
			case 5:
				periode = year + "-05"
			case 6:
				periode = year + "-06"
			case 7:
				periode = year + "-07"
			case 8:
				periode = year + "-08"
			case 9:
				periode = year + "-09"
			case 10:
				periode = year + "-10"
			case 11:
				periode = year + "-11"
			case 12:
				periode = year + "-12"
			}
			idinvoice = _invoicewinlose_getidinvoice(idcompany_db, year, periode)
			var winlose int = _invoicewinlosepasaran_id(idinvoice, idcomppasaran_db)
			objdetail.Dashboardcompanypasaran_winlose = winlose
			arraobjdetail = append(arraobjdetail, objdetail)
			msg = "Success"
		}

		obj.Dashboardcompanypasaran_nmpasaran = nmpasarantogel_db
		obj.Dashboardcompanypasaran_detail = arraobjdetail
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
