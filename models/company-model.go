package models

import (
	"context"
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/config"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/db"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/entities"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/helpers"
	"github.com/nleeper/goment"
)

type sCompanyListPasaranOnline struct {
	Pasaran_onlineid int    `json:"company_pasaranonline_id"`
	Pasaran_harian   string `json:"company_pasaranonline_hari"`
}
type scompanylistkeluaran struct {
	No                int     `json:"pasaran_no"`
	Idtrxkeluaran     int     `json:"pasaran_invoice"`
	Idcomppasaran     int     `json:"pasaran_idcompp"`
	Pasarancode       string  `json:"pasaran_code"`
	Keluaranperiode   string  `json:"pasaran_periode"`
	Nmpasaran         string  `json:"pasaran_name"`
	Tanggalperiode    string  `json:"pasaran_tanggal"`
	Keluarantogel     string  `json:"pasaran_keluaran"`
	Status            string  `json:"pasaran_status"`
	Status_css        string  `json:"pasaran_status_css"`
	Total_Member      float32 `json:"pasaran_totalmember"`
	Total_bet         float32 `json:"pasaran_totalbet"`
	Total_outstanding float32 `json:"pasaran_totaloutstanding"`
	Total_cancelbet   float32 `json:"pasaran_totalcancelbet"`
	Winlose           float32 `json:"pasaran_winlose"`
	Winlosetemp       int     `json:"pasaran_winlosetemp"`
	Revisi            int     `json:"pasaran_revisi"`
	Msgrevisi         string  `json:"pasaran_msgrevisi"`
}
type invoicelistMember struct {
	Member         string `json:"member"`
	Totalbet       int    `json:"totalbet"`
	Totalbayar     int    `json:"totalbayar"`
	Totalcancelbet int    `json:"totalcancelbet"`
	Totalwin       int    `json:"totalwin"`
}
type invoicelistGroupPermainan struct {
	Permainan string `json:"permainan"`
}
type invoicelistpermainan struct {
	Bet_id           int     `json:"bet_id"`
	Bet_datetime     string  `json:"bet_datetime"`
	Bet_ipaddress    string  `json:"bet_ipaddress"`
	Bet_device       string  `json:"bet_device"`
	Bet_timezone     string  `json:"bet_timezone"`
	Bet_username     string  `json:"bet_username"`
	Bet_typegame     string  `json:"bet_typegame"`
	Bet_nomortogel   string  `json:"bet_nomortogel"`
	Bet_bet          int     `json:"bet_bet"`
	Bet_diskon       int     `json:"bet_diskon"`
	Bet_diskonpercen int     `json:"bet_diskonpercen"`
	Bet_kei          int     `json:"bet_kei"`
	Bet_keipercen    int     `json:"bet_keipercen"`
	Bet_win          float32 `json:"bet_win"`
	Bet_totalwin     int     `json:"bet_totalwin"`
	Bet_bayar        int     `json:"bet_bayar"`
	Bet_status       string  `json:"bet_status"`
	Bet_statuscss    string  `json:"bet_statuscss"`
	Bet_create       string  `json:"bet_create"`
	Bet_createDate   string  `json:"bet_createdate"`
	Bet_update       string  `json:"bet_update"`
	Bet_updateDate   string  `json:"bet_updatedate"`
}
type sCompanyDetailPasaranConf struct {
	Pasaran_diundi                    string  `json:"pasaran_diundi"`
	Pasaran_url                       string  `json:"pasaran_url"`
	Pasaran_jamtutup                  string  `json:"pasaran_jamtutup"`
	Pasaran_jamjadwal                 string  `json:"pasaran_jamjadwal"`
	Pasaran_jamopen                   string  `json:"pasaran_jamopen"`
	Pasaran_statusactive              string  `json:"pasaran_statusactive"`
	Limitline4d                       int     `json:"limitline_4d"`
	Limitline3d                       int     `json:"limitline_3d"`
	Limitline2d                       int     `json:"limitline_2d"`
	Limitline2dd                      int     `json:"limitline_2dd"`
	Limitline2dt                      int     `json:"limitline_2dt"`
	Bbfs                              int     `json:"bbfs"`
	Minbet_432d                       float32 `json:"minbet_432d"`
	Maxbet4d_432d                     float32 `json:"maxbet4d_432d"`
	Maxbet3d_432d                     float32 `json:"maxbet3d_432d"`
	Maxbet2d_432d                     float32 `json:"maxbet2d_432d"`
	Maxbet2dd_432d                    float32 `json:"maxbet2dd_432d"`
	Maxbet2dt_432d                    float32 `json:"maxbet2dt_432d"`
	Limitotal4d_432d                  float32 `json:"limitotal4d_432d"`
	Limitotal3d_432d                  float32 `json:"limitotal3d_432d"`
	Limitotal2d_432d                  float32 `json:"limitotal2d_432d"`
	Limitotal2dd_432d                 float32 `json:"limitotal2dd_432d"`
	Limitotal2dt_432d                 float32 `json:"limitotal2dt_432d"`
	Limitglobal4d_432d                float32 `json:"limitglobal4d_432d"`
	Limitglobal3d_432d                float32 `json:"limitglobal3d_432d"`
	Limitglobal2d_432d                float32 `json:"limitglobal2d_432d"`
	Limitglobal2dd_432d               float32 `json:"limitglobal2dd_432d"`
	Limitglobal2dt_432d               float32 `json:"limitglobal2dt_432d"`
	Disc4d_432d                       float32 `json:"disc4d_432d"`
	Disc3d_432d                       float32 `json:"disc3d_432d"`
	Disc2d_432d                       float32 `json:"disc2d_432d"`
	Disc2dd_432d                      float32 `json:"disc2dd_432d"`
	Disc2dt_432d                      float32 `json:"disc2dt_432d"`
	Win4d_432d                        float32 `json:"win4d_432d"`
	Win3d_432d                        float32 `json:"win3d_432d"`
	Win2d_432d                        float32 `json:"win2d_432d"`
	Win2dd_432d                       float32 `json:"win2dd_432d"`
	Win2dt_432d                       float32 `json:"win2dt_432d"`
	Minbet_cbebas                     float32 `json:"minbet_cbebas"`
	Maxbet_cbebas                     float32 `json:"maxbet_cbebas"`
	Win_cbebas                        float32 `json:"win_cbebas"`
	Disc_cbebas                       float32 `json:"disc_cbebas"`
	Limitglobal_cbebas                float32 `json:"limitglobal_cbebas"`
	Limittotal_cbebas                 float32 `json:"limittotal_cbebas"`
	Minbet_cmacau                     float32 `json:"minbet_cmacau"`
	Maxbet_cmacau                     float32 `json:"maxbet_cmacau"`
	Win2d_cmacau                      float32 `json:"win2d_cmacau"`
	Win3d_cmacau                      float32 `json:"win3d_cmacau"`
	Win4d_cmacau                      float32 `json:"win4d_cmacau"`
	Disc_cmacau                       float32 `json:"disc_cmacau"`
	Limitglobal_cmacau                float32 `json:"limitglobal_cmacau"`
	Limitotal_cmacau                  float32 `json:"limitotal_cmacau"`
	Minbet_cnaga                      float32 `json:"minbet_cnaga"`
	Maxbet_cnaga                      float32 `json:"maxbet_cnaga"`
	Win3_cnaga                        float32 `json:"win3_cnaga"`
	Win4_cnaga                        float32 `json:"win4_cnaga"`
	Disc_cnaga                        float32 `json:"disc_cnaga"`
	Limitglobal_cnaga                 float32 `json:"limitglobal_cnaga"`
	Limittotal_cnaga                  float32 `json:"limittotal_cnaga"`
	Minbet_cjitu                      float32 `json:"minbet_cjitu"`
	Maxbet_cjitu                      float32 `json:"maxbet_cjitu"`
	Winas_cjitu                       float32 `json:"winas_cjitu"`
	Winkop_cjitu                      float32 `json:"winkop_cjitu"`
	Winkepala_cjitu                   float32 `json:"winkepala_cjitu"`
	Winekor_cjitu                     float32 `json:"winekor_cjitu"`
	Desc_cjitu                        float32 `json:"desc_cjitu"`
	Limitglobal_cjitu                 float32 `json:"limitglobal_cjitu"`
	Limittotal_cjitu                  float32 `json:"limittotal_cjitu"`
	Minbet_5050umum                   float32 `json:"minbet_5050umum"`
	Maxbet_5050umum                   float32 `json:"maxbet_5050umum"`
	Keibesar_5050umum                 float32 `json:"keibesar_5050umum"`
	Keikecil_5050umum                 float32 `json:"keikecil_5050umum"`
	Keigenap_5050umum                 float32 `json:"keigenap_5050umum"`
	Keiganjil_5050umum                float32 `json:"keiganjil_5050umum"`
	Keitengah_5050umum                float32 `json:"keitengah_5050umum"`
	Keitepi_5050umum                  float32 `json:"keitepi_5050umum"`
	Discbesar_5050umum                float32 `json:"discbesar_5050umum"`
	Disckecil_5050umum                float32 `json:"disckecil_5050umum"`
	Discgenap_5050umum                float32 `json:"discgenap_5050umum"`
	Discganjil_5050umum               float32 `json:"discganjil_5050umum"`
	Disctengah_5050umum               float32 `json:"disctengah_5050umum"`
	Disctepi_5050umum                 float32 `json:"disctepi_5050umum"`
	Limitglobal_5050umum              float32 `json:"limitglobal_5050umum"`
	Limittotal_5050umum               float32 `json:"limittotal_5050umum"`
	Minbet_5050special                float32 `json:"minbet_5050special"`
	Maxbet_5050special                float32 `json:"maxbet_5050special"`
	Keiasganjil_5050special           float32 `json:"keiasganjil_5050special"`
	Keiasgenap_5050special            float32 `json:"keiasgenap_5050special"`
	Keiasbesar_5050special            float32 `json:"keiasbesar_5050special"`
	Keiaskecil_5050special            float32 `json:"keiaskecil_5050special"`
	Keikopganjil_5050special          float32 `json:"keikopganjil_5050special"`
	Keikopgenap_5050special           float32 `json:"keikopgenap_5050special"`
	Keikopbesar_5050special           float32 `json:"keikopbesar_5050special"`
	Keikopkecil_5050special           float32 `json:"keikopkecil_5050special"`
	Keikepalaganjil_5050special       float32 `json:"keikepalaganjil_5050special"`
	Keikepalagenap_5050special        float32 `json:"keikepalagenap_5050special"`
	Keikepalabesar_5050special        float32 `json:"keikepalabesar_5050special"`
	Keikepalakecil_5050special        float32 `json:"keikepalakecil_5050special"`
	Keiekorganjil_5050special         float32 `json:"keiekorganjil_5050special"`
	Keiekorgenap_5050special          float32 `json:"keiekorgenap_5050special"`
	Keiekorbesar_5050special          float32 `json:"keiekorbesar_5050special"`
	Keiekorkecil_5050special          float32 `json:"keiekorkecil_5050special"`
	Discasganjil_5050special          float32 `json:"discasganjil_5050special"`
	Discasgenap_5050special           float32 `json:"discasgenap_5050special"`
	Discasbesar_5050special           float32 `json:"discasbesar_5050special"`
	Discaskecil_5050special           float32 `json:"discaskecil_5050special"`
	Disckopganjil_5050special         float32 `json:"disckopganjil_5050special"`
	Disckopgenap_5050special          float32 `json:"disckopgenap_5050special"`
	Disckopbesar_5050special          float32 `json:"disckopbesar_5050special"`
	Disckopkecil_5050special          float32 `json:"disckopkecil_5050special"`
	Disckepalaganjil_5050special      float32 `json:"disckepalaganjil_5050special"`
	Disckepalagenap_5050special       float32 `json:"disckepalagenap_5050special"`
	Disckepalabesar_5050special       float32 `json:"disckepalabesar_5050special"`
	Disckepalakecil_5050special       float32 `json:"disckepalakecil_5050special"`
	Discekorganjil_5050special        float32 `json:"discekorganjil_5050special"`
	Discekorgenap_5050special         float32 `json:"discekorgenap_5050special"`
	Discekorbesar_5050special         float32 `json:"discekorbesar_5050special"`
	Discekorkecil_5050special         float32 `json:"discekorkecil_5050special"`
	Limitglobal_5050special           float32 `json:"limitglobal_5050special"`
	Limittotal_5050special            float32 `json:"limittotal_5050special"`
	Minbet_5050kombinasi              float32 `json:"minbet_5050kombinasi"`
	Maxbet_5050kombinasi              float32 `json:"maxbet_5050kombinasi"`
	Belakangkeimono_5050kombinasi     float32 `json:"belakangkeimono_5050kombinasi"`
	Belakangkeistereo_5050kombinasi   float32 `json:"belakangkeistereo_5050kombinasi"`
	Belakangkeikembang_5050kombinasi  float32 `json:"belakangkeikembang_5050kombinasi"`
	Belakangkeikempis_5050kombinasi   float32 `json:"belakangkeikempis_5050kombinasi"`
	Belakangkeikembar_5050kombinasi   float32 `json:"belakangkeikembar_5050kombinasi"`
	Tengahkeimono_5050kombinasi       float32 `json:"tengahkeimono_5050kombinasi"`
	Tengahkeistereo_5050kombinasi     float32 `json:"tengahkeistereo_5050kombinasi"`
	Tengahkeikembang_5050kombinasi    float32 `json:"tengahkeikembang_5050kombinasi"`
	Tengahkeikempis_5050kombinasi     float32 `json:"tengahkeikempis_5050kombinasi"`
	Tengahkeikembar_5050kombinasi     float32 `json:"tengahkeikembar_5050kombinasi"`
	Depankeimono_5050kombinasi        float32 `json:"depankeimono_5050kombinasi"`
	Depankeistereo_5050kombinasi      float32 `json:"depankeistereo_5050kombinasi"`
	Depankeikembang_5050kombinasi     float32 `json:"depankeikembang_5050kombinasi"`
	Depankeikempis_5050kombinasi      float32 `json:"depankeikempis_5050kombinasi"`
	Depankeikembar_5050kombinasi      float32 `json:"depankeikembar_5050kombinasi"`
	Belakangdiscmono_5050kombinasi    float32 `json:"belakangdiscmono_5050kombinasi"`
	Belakangdiscstereo_5050kombinasi  float32 `json:"belakangdiscstereo_5050kombinasi"`
	Belakangdisckembang_5050kombinasi float32 `json:"belakangdisckembang_5050kombinasi"`
	Belakangdisckempis_5050kombinasi  float32 `json:"belakangdisckempis_5050kombinasi"`
	Belakangdisckembar_5050kombinasi  float32 `json:"belakangdisckembar_5050kombinasi"`
	Tengahdiscmono_5050kombinasi      float32 `json:"tengahdiscmono_5050kombinasi"`
	Tengahdiscstereo_5050kombinasi    float32 `json:"tengahdiscstereo_5050kombinasi"`
	Tengahdisckembang_5050kombinasi   float32 `json:"tengahdisckembang_5050kombinasi"`
	Tengahdisckempis_5050kombinasi    float32 `json:"tengahdisckempis_5050kombinasi"`
	Tengahdisckembar_5050kombinasi    float32 `json:"tengahdisckembar_5050kombinasi"`
	Depandiscmono_5050kombinasi       float32 `json:"depandiscmono_5050kombinasi"`
	Depandiscstereo_5050kombinasi     float32 `json:"depandiscstereo_5050kombinasi"`
	Depandisckembang_5050kombinasi    float32 `json:"depandisckembang_5050kombinasi"`
	Depandisckempis_5050kombinasi     float32 `json:"depandisckempis_5050kombinasi"`
	Depandisckembar_5050kombinasi     float32 `json:"depandisckembar_5050kombinasi"`
	Limitglobal_5050kombinasi         float32 `json:"limitglobal_5050kombinasi"`
	Limittotal_5050kombinasi          float32 `json:"limittotal_5050kombinasi"`
	Minbet_kombinasi                  float32 `json:"minbet_kombinasi"`
	Maxbet_kombinasi                  float32 `json:"maxbet_kombinasi"`
	Win_kombinasi                     float32 `json:"win_kombinasi"`
	Disc_kombinasi                    float32 `json:"disc_kombinasi"`
	Limitglobal_kombinasi             float32 `json:"limitglobal_kombinasi"`
	Limittotal_kombinasi              float32 `json:"limittotal_kombinasi"`
	Minbet_dasar                      float32 `json:"minbet_dasar"`
	Maxbet_dasar                      float32 `json:"maxbet_dasar"`
	Keibesar_dasar                    float32 `json:"keibesar_dasar"`
	Keikecil_dasar                    float32 `json:"keikecil_dasar"`
	Keigenap_dasar                    float32 `json:"keigenap_dasar"`
	Keiganjil_dasar                   float32 `json:"keiganjil_dasar"`
	Discbesar_dasar                   float32 `json:"discbesar_dasar"`
	Disckecil_dasar                   float32 `json:"disckecil_dasar"`
	Discgenap_dasar                   float32 `json:"discgenap_dasar"`
	Discganjil_dasar                  float32 `json:"discganjil_dasar"`
	Limitglobal_dasar                 float32 `json:"limitglobal_dasar"`
	Limittotal_dasar                  float32 `json:"limittotal_dasar"`
	Minbet_shio                       float32 `json:"minbet_shio"`
	Maxbet_shio                       float32 `json:"maxbet_shio"`
	Win_shio                          float32 `json:"win_shio"`
	Disc_shio                         float32 `json:"disc_shio"`
	Shioyear_shio                     string  `json:"shioyear_shio"`
	Limitglobal_shio                  float32 `json:"limitglobal_shio"`
	Limittotal_shio                   float32 `json:"limittotal_shio"`
}

func Fetch_company() (helpers.Response, error) {
	var obj entities.Model_company
	var arraobj []entities.Model_company
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	bulanskrg := tglnow.Format("MMMM") + "-" + tglnow.Format("YYYY")
	render_page := time.Now()
	endday, _, _ := helpers.GetEndRangeDate(tglnow.Format("MM"))
	start := tglnow.Format("YYYY-MM") + "-" + "01"
	end := tglnow.Format("YYYY-MM") + "-" + endday

	var no int = 0
	sql_periode := `SELECT 
			idcompany , startjoincompany, COALESCE(endjoincompany,""), idcurr, 
			nmcompany, statuscompany   
			FROM ` + config.DB_tbl_mst_company + ` 
			ORDER BY idcompany ASC 
		`
	row, err := con.QueryContext(ctx, sql_periode)
	helpers.ErrorCheck(err)
	for row.Next() {
		no++
		var (
			idcompany_db, startjoincompany_db, endjoincompany_db, idcurr_db string
			nmcompany_db, statuscompany_db                                  string
		)

		err = row.Scan(
			&idcompany_db, &startjoincompany_db, &endjoincompany_db,
			&idcurr_db, &nmcompany_db, &statuscompany_db)
		helpers.ErrorCheck(err)

		status_css := ""
		switch statuscompany_db {
		case "ACTIVE":
			status_css = config.STATUS_COMPLETE
		case "DEACTIVE":
			status_css = config.STATUS_CANCEL
		}
		if startjoincompany_db == "" {
			startjoincompany_db = ""
		}
		if endjoincompany_db == "" {
			endjoincompany_db = ""
		}

		obj.Company_no = no
		obj.Company_idcompany = idcompany_db
		obj.Company_startjoin = startjoincompany_db
		obj.Company_endjoin = endjoincompany_db
		obj.Company_curr = idcurr_db
		obj.Company_name = nmcompany_db
		obj.Company_periode = bulanskrg
		obj.Company_winlose = _winlose(idcompany_db, start, end, 0)
		obj.Company_winlosetemp = _winlosetemp(idcompany_db, start, end, 0)
		obj.Company_status = statuscompany_db
		obj.Company_statuscss = status_css
		arraobj = append(arraobj, obj)
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Fetch_companyDetail(company string) (helpers.Response, error) {
	var obj entities.Model_companydetail
	var arraobj []entities.Model_companydetail
	var res helpers.Response
	msg := "Data Not Found"
	flag := true
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()

	sql_detail := `SELECT 
		nmcompany, companyurl,  
		statuscompany, createcompany, createdatecompany, updatecompany, updatedatecompany  
		FROM ` + config.DB_tbl_mst_company + `
		WHERE idcompany = ? 
	`
	var (
		nmcompany_db, companyurl_db                                                                      string
		statuscompany_db, createcompany_db, createdatecompany_db, updatecompany_db, updatedatecompany_db string
	)
	rows := con.QueryRowContext(ctx, sql_detail, company)
	switch err := rows.Scan(
		&nmcompany_db, &companyurl_db, &statuscompany_db,
		&createcompany_db, &createdatecompany_db, &updatecompany_db, &updatedatecompany_db); err {
	case sql.ErrNoRows:
		flag = false
	case nil:
		create := ""
		update := ""
		if createcompany_db != "" {
			create = createcompany_db + ", " + createdatecompany_db
		}
		if updatecompany_db != "" {
			update = updatecompany_db + ", " + updatedatecompany_db
		}
		obj.Company_name = nmcompany_db
		obj.Company_url = companyurl_db
		obj.Company_status = statuscompany_db
		obj.Company_create = create
		obj.Company_update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	default:
		flag = false
		helpers.ErrorCheck(err)
	}
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
func Fetch_company_listadmin(company string) (helpers.Response, error) {
	var obj entities.Model_companylistadmin
	var arraobj []entities.Model_companylistadmin
	var res helpers.Response
	msg := "Error"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	var no int = 0

	sql_periode := `SELECT 
			username_comp, typeadmin, nama_comp, status_comp, lastlogin_comp, 
			lastipaddres_comp, createcomp_admin, createdatecomp_admin, updatecomp_admin, updatedatecomp_admin    
			FROM ` + config.DB_tbl_mst_company_admin + ` 
			WHERE idcompany = ? 
			ORDER BY lastlogin_comp  
		`

	row, err := con.QueryContext(ctx, sql_periode, company)
	helpers.ErrorCheck(err)
	for row.Next() {
		no++
		var (
			username_comp_db, typeadmin_db, nama_comp_db, status_comp_db, lastlogin_comp_db                                  string
			lastipaddres_comp_db, createcomp_admin_db, createdatecomp_admin_db, updatecomp_admin_db, updatedatecomp_admin_db string
		)

		err = row.Scan(
			&username_comp_db, &typeadmin_db, &nama_comp_db, &status_comp_db, &lastlogin_comp_db,
			&lastipaddres_comp_db, &createcomp_admin_db, &createdatecomp_admin_db, &updatecomp_admin_db, &updatedatecomp_admin_db)
		helpers.ErrorCheck(err)

		status_css := ""
		switch status_comp_db {
		case "ACTIVE":
			status_css = config.STATUS_COMPLETE
		case "DEACTIVE":
			status_css = config.STATUS_CANCEL
		}
		if lastlogin_comp_db == "0000-00-00 00:00:00" {
			lastlogin_comp_db = ""
		}
		create := ""
		update := ""
		if createcomp_admin_db != "" {
			create = createcomp_admin_db + ", " + createdatecomp_admin_db
		}
		if updatecomp_admin_db != "" {
			update = updatecomp_admin_db + ", " + updatedatecomp_admin_db
		}
		obj.Company_admin_username = username_comp_db
		obj.Company_admin_typeadmin = typeadmin_db
		obj.Company_admin_name = nama_comp_db
		obj.Company_admin_status = status_comp_db
		obj.Company_admin_statuscss = status_css
		obj.Company_admin_lastlogin = lastlogin_comp_db
		obj.Company_admin_lastippadress = lastipaddres_comp_db
		obj.Company_admin_create = create
		obj.Company_admin_update = update
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
func Fetch_company_listpasaran(company string) (helpers.Response, error) {
	var obj entities.Model_companylistpasaran
	var arraobj []entities.Model_companylistpasaran
	var res helpers.Response
	msg := "Error"
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	bulanskrg := tglnow.Format("MMMM") + "-" + tglnow.Format("YYYY")
	render_page := time.Now()
	endday, _, _ := helpers.GetEndRangeDate(tglnow.Format("MM"))
	start := tglnow.Format("YYYY-MM") + "-" + "01"
	end := tglnow.Format("YYYY-MM") + "-" + endday
	var no int = 0

	sql_periode := `SELECT 
			A.idcomppasaran , A.idpasarantogel, A.statuspasaran, A.statuspasaranactive, A.displaypasaran, B.nmpasarantogel 
			FROM ` + config.DB_tbl_mst_company_game_pasaran + ` as A 
			JOIN ` + config.DB_tbl_mst_pasaran_togel + ` as B ON B.idpasarantogel = A.idpasarantogel 
			WHERE A.idcompany = ? 
			ORDER BY A.displaypasaran ASC   
		`

	row, err := con.QueryContext(ctx, sql_periode, company)
	helpers.ErrorCheck(err)
	for row.Next() {
		no++
		var (
			idcomppasaran_db, displaypasaran_db                                            int
			idpasarantogel_db, statuspasaran_db, statuspasaranactive_db, nmpasarantogel_db string
		)

		err = row.Scan(
			&idcomppasaran_db, &idpasarantogel_db, &statuspasaran_db, &statuspasaranactive_db, &displaypasaran_db, &nmpasarantogel_db)
		helpers.ErrorCheck(err)

		status_css := ""
		status_cssactive := ""
		switch statuspasaran_db {
		case "ONLINE":
			status_css = config.STATUS_RUNNING
		case "OFFLINE":
			status_css = config.STATUS_CANCEL
		}
		switch statuspasaranactive_db {
		case "Y":
			status_cssactive = config.STATUS_COMPLETE
		case "N":
			status_cssactive = config.STATUS_CANCEL
		}

		obj.Company_pasaran_idcomppasaran = idcomppasaran_db
		obj.Company_pasaran_idpasarantogel = idpasarantogel_db
		obj.Company_pasaran_nmpasarantogel = nmpasarantogel_db
		obj.Company_pasaran_periode = bulanskrg
		obj.Company_pasaran_winlose = _winlose(company, start, end, idcomppasaran_db)
		obj.Company_pasaran_displaypasaran = displaypasaran_db
		obj.Company_pasaran_status = statuspasaran_db
		obj.Company_pasaran_statuscss = status_css
		obj.Company_pasaran_statusactive = statuspasaranactive_db
		obj.Company_pasaran_statusactivecss = status_cssactive

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
func Fetch_company_listpasaranConf(company string, idcomppasaran int) (helpers.Response, error) {
	var obj entities.Model_companypasaranconf
	var arraobj []entities.Model_companypasaranconf
	var res helpers.Response
	msg := "Error"
	flag := true
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()

	sql_detail := `SELECT 
		pasarandiundi, pasaranurl, jamtutup, jamjadwal, jamopen, statuspasaranactive,  
		limitline_4d, limitline_3d, limitline_2d, limitline_2dd, limitline_2dt, bbfs, 
		1_minbet as minbet_432d, 1_maxbet4d as maxbet4d_432d, 1_maxbet3d as maxbet3d_432d,
		1_maxbet2d as maxbet2d_432d, 1_maxbet2dd as maxbet2dd_432d, 1_maxbet2dt as maxbet2dt_432d, 
		1_limittotal4d as limitotal4d_432d, 1_limittotal3d as limitotal3d_432d, 1_limittotal2d as limitotal2d_432d, 
		1_limittotal2dd as limitotal2dd_432d, 1_limittotal2dt as limitotal2dt_432d, 
		1_limitbuang4d as limitglobal4d_432d, 1_limitbuang3d as limitglobal3d_432d, 1_limitbuang2d as limitglobal2d_432d, 
		1_limitbuang2dd as limitglobal2dd_432d, 1_limitbuang2dt as limitglobal2dt_432d, 
		1_disc4d as disc4d_432d, 1_disc3d as disc3d_432d, 1_disc2d as disc2d_432d, 1_disc2dd as disc2dd_432d, 1_disc2dt as disc2dt_432d, 
		1_win4d as win4d_432d, 1_win3d as win3d_432d, 1_win2d as win2d_432d, 1_win2dd as win2dd_432d, 1_win2dt as win2dt_432d, 
		2_minbet as minbet_cbebas, 2_maxbet as maxbet_cbebas, 
		2_win as win_cbebas, 2_disc as disc_cbebas, 
		2_limitbuang as limitglobal_cbebas, 2_limitotal as limittotal_cbebas, 
		3_minbet as minbet_cmacau, 3_maxbet as maxbet_cmacau, 
		3_win2digit as win2d_cmacau, 3_win3digit as win3d_cmacau, 3_win4digit as win4d_cmacau, 
		3_disc as disc_cmacau, 3_limitbuang as limitglobal_cmacau, 3_limittotal as limitotal_cmacau, 
		4_minbet as minbet_cnaga, 4_maxbet as maxbet_cnaga, 
		4_win3digit as win3_cnaga, 4_win4digit as win4_cnaga, 
		4_disc as disc_cnaga, 4_limitbuang as limitglobal_cnaga, 4_limittotal as limittotal_cnaga, 
		5_minbet as minbet_cjitu, 5_maxbet as maxbet_cjitu, 
		5_winas as winas_cjitu, 5_winkop as winkop_cjitu, 5_winkepala as winkepala_cjitu, 5_winekor as winekor_cjitu, 
		5_desic as desc_cjitu, 5_limitbuang as limitglobal_cjitu, 5_limitotal as limittotal_cjitu, 
		6_minbet as minbet_5050umum, 6_maxbet as maxbet_5050umum, 
		6_keibesar as keibesar_5050umum, 6_keikecil as keikecil_5050umum, 6_keigenap as keigenap_5050umum, 
		6_keiganjil as keiganjil_5050umum, 6_keitengah as keitengah_5050umum, 6_keitepi as keitepi_5050umum, 
		6_discbesar as discbesar_5050umum, 6_disckecil as disckecil_5050umum, 6_discgenap as discgenap_5050umum, 
		6_discganjil as discganjil_5050umum, 6_disctengah as disctengah_5050umum, 6_disctepi as disctepi_5050umum, 
		6_limitbuang as limitglobal_5050umum, 6_limittotal as limittotal_5050umum, 
		7_minbet as minbet_5050special, 7_maxbet as maxbet_5050special,
		7_keiasganjil as keiasganjil_5050special, 7_keiasgenap as keiasgenap_5050special, 7_keiasbesar as keiasbesar_5050special, 
		7_keiaskecil as keiaskecil_5050special, 7_keikopganjil as keikopganjil_5050special, 7_keikopgenap as keikopgenap_5050special, 
		7_keikopbesar as keikopbesar_5050special, 7_keikopkecil as keikopkecil_5050special, 7_keikepalaganjil as keikepalaganjil_5050special, 
		7_keikepalagenap as keikepalagenap_5050special, 7_keikepalabesar as keikepalabesar_5050special, 7_keikepalakecil as keikepalakecil_5050special, 
		7_keiekorganjil as keiekorganjil_5050special, 7_keiekorgenap as keiekorgenap_5050special, 7_keiekorbesar as keiekorbesar_5050special, 
		7_keiekorkecil as keiekorkecil_5050special, 
		7_discasganjil as discasganjil_5050special, 7_discasgenap as discasgenap_5050special, 7_discasbesar as discasbesar_5050special, 
		7_discaskecil as discaskecil_5050special, 7_disckopganjil as disckopganjil_5050special, 7_disckopgenap as disckopgenap_5050special, 
		7_disckopbesar as disckopbesar_5050special, 7_disckopkecil as disckopkecil_5050special, 7_disckepalaganjil as disckepalaganjil_5050special, 
		7_disckepalagenap as disckepalagenap_5050special, 7_disckepalabesar as disckepalabesar_5050special, 7_disckepalakecil as disckepalakecil_5050special, 
		7_discekorganjil as discekorganjil_5050special, 7_discekorgenap as discekorgenap_5050special, 7_discekorbesar as discekorbesar_5050special, 
		7_discekorkecil as discekorkecil_5050special, 7_limitbuang as limitglobal_5050special, 7_limittotal as limittotal_5050special, 
		8_minbet as minbet_5050kombinasi, 8_maxbet as maxbet_5050kombinasi, 
		8_belakangkeimono as belakangkeimono_5050kombinasi, 8_belakangkeistereo as belakangkeistereo_5050kombinasi, 8_belakangkeikembang as belakangkeikembang_5050kombinasi, 8_belakangkeikempis as belakangkeikempis_5050kombinasi, 8_belakangkeikembar as belakangkeikembar_5050kombinasi, 
		8_tengahkeimono as tengahkeimono_5050kombinasi, 8_tengahkeistereo as tengahkeistereo_5050kombinasi, 8_tengahkeikembang as tengahkeikembang_5050kombinasi, 8_tengahkeikempis as tengahkeikempis_5050kombinasi, 8_tengahkeikembar as tengahkeikembar_5050kombinasi, 
		8_depankeimono as depankeimono_5050kombinasi, 8_depankeistereo as depankeistereo_5050kombinasi, 8_depankeikembang as depankeikembang_5050kombinasi, 8_depankeikempis as depankeikempis_5050kombinasi, 8_depankeikembar as depankeikembar_5050kombinasi, 
		8_belakangdiscmono as belakangdiscmono_5050kombinasi, 8_belakangdiscstereo as belakangdiscstereo_5050kombinasi, 8_belakangdisckembang as belakangdisckembang_5050kombinasi, 8_belakangdisckempis as belakangdisckempis_5050kombinasi, 8_belakangdisckembar as belakangdisckembar_5050kombinasi, 
		8_tengahdiscmono as tengahdiscmono_5050kombinasi, 8_tengahdiscstereo as tengahdiscstereo_5050kombinasi, 8_tengahdisckembang as tengahdisckembang_5050kombinasi, 8_tengahdisckempis as tengahdisckempis_5050kombinasi, 8_tengahdisckembar as tengahdisckembar_5050kombinasi, 
		8_depandiscmono as depandiscmono_5050kombinasi, 8_depandiscstereo as depandiscstereo_5050kombinasi, 8_depandisckembang as depandisckembang_5050kombinasi, 8_depandisckempis as depandisckempis_5050kombinasi, 8_depandisckembar as depandisckembar_5050kombinasi, 
		8_limitbuang as limitglobal_5050kombinasi, 8_limittotal as limittotal_5050kombinasi, 
		9_minbet as minbet_kombinasi, 9_maxbet as maxbet_kombinasi, 9_win as win_kombinasi, 9_discount as disc_kombinasi, 9_limitbuang as limitglobal_kombinasi, 9_limittotal as limittotal_kombinasi, 
		10_minbet as minbet_dasar, 10_maxbet as maxbet_dasar, 
		10_keibesar as keibesar_dasar, 10_keikecil as keikecil_dasar, 10_keigenap as keigenap_dasar, 10_keiganjil as keiganjil_dasar, 
		10_discbesar as discbesar_dasar, 10_disckecil as disckecil_dasar, 10_discigenap as discgenap_dasar, 10_discganjil as discganjil_dasar, 
		10_limitbuang as limitglobal_dasar, 10_limittotal as limittotal_dasar, 
		11_minbet as minbet_shio, 11_maxbet as maxbet_shio, 11_win as win_shio, 11_disc as disc_shio, 11_limitbuang as limitglobal_shio, 11_limittotal as limittotal_shio, 
		11_shiotahunini as shioyear_shio  
		FROM ` + config.DB_tbl_mst_company_game_pasaran + ` 
		WHERE idcomppasaran = ?  AND idcompany=? 
	`
	var (
		pasarandiundi_db, pasaranurl_db, jamtutup_db, jamjadwal_db, jamopen_db, statuspasaranactive_db                                                                                                                                                                                                           string
		limitline_4d_db, limitline_3d_db, limitline_2d_db, limitline_2dd_db, limitline_2dt_db, bbfs_db                                                                                                                                                                                                           int
		minbet_432d_db, maxbet4d_432d_db, maxbet3d_432d_db, maxbet2d_432d_db, maxbet2dd_432d_db, maxbet2dt_432d_db                                                                                                                                                                                               float32
		limitotal4d_432d_db, limitotal3d_432d_db, limitotal2d_432d_db, limitotal2dd_432d_db, limitotal2dt_432d_db                                                                                                                                                                                                float32
		limitglobal4d_432d_db, limitglobal3d_432d_db, limitglobal2d_432d_db, limitglobal2dd_432d_db, limitglobal2dt_432d_db                                                                                                                                                                                      float32
		disc4d_432d_db, disc3d_432d_db, disc2d_432d_db, disc2dd_432d_db, disc2dt_432d_db                                                                                                                                                                                                                         float32
		win4d_432d_db, win3d_432d_db, win2d_432d_db, win2dd_432d_db, win2dt_432d_db                                                                                                                                                                                                                              float32
		minbet_cbebas_db, maxbet_cbebas_db, win_cbebas_db, disc_cbebas_db, limitglobal_cbebas_db, limittotal_cbebas_db                                                                                                                                                                                           float32
		minbet_cmacau_db, maxbet_cmacau_db, win2d_cmacau_db, win3d_cmacau_db, win4d_cmacau_db, disc_cmacau_db, limitglobal_cmacau_db, limitotal_cmacau_db                                                                                                                                                        float32
		minbet_cnaga_db, maxbet_cnaga_db, win3_cnaga_db, win4_cnaga_db, disc_cnaga_db, limitglobal_cnaga_db, limittotal_cnaga_db                                                                                                                                                                                 float32
		minbet_cjitu_db, maxbet_cjitu_db, winas_cjitu_db, winkop_cjitu_db, winkepala_cjitu_db, winekor_cjitu_db, desc_cjitu_db, limitglobal_cjitu_db, limittotal_cjitu_db                                                                                                                                        float32
		minbet_5050umum_db, maxbet_5050umum_db, keibesar_5050umum_db, keikecil_5050umum_db, keigenap_5050umum_db, keiganjil_5050umum_db, keitengah_5050umum_db, keitepi_5050umum_db                                                                                                                              float32
		discbesar_5050umum_db, disckecil_5050umum_db, discgenap_5050umum_db, discganjil_5050umum_db, disctengah_5050umum_db, disctepi_5050umum_db, limitglobal_5050umum_db, limittotal_5050umum_db                                                                                                               float32
		minbet_5050special_db, maxbet_5050special_db, keiasganjil_5050special_db, keiasgenap_5050special_db, keiasbesar_5050special_db, keiaskecil_5050special_db, keikopganjil_5050special_db, keikopgenap_5050special_db                                                                                       float32
		keikopbesar_5050special_db, keikopkecil_5050special_db, keikepalaganjil_5050special_db, keikepalagenap_5050special_db, keikepalabesar_5050special_db, keikepalakecil_5050special_db, keiekorganjil_5050special_db, keiekorgenap_5050special_db, keiekorbesar_5050special_db, keiekorkecil_5050special_db float32
		discasganjil_5050special_db, discasgenap_5050special_db, discasbesar_5050special_db, discaskecil_5050special_db, disckopganjil_5050special_db, disckopgenap_5050special_db, disckopbesar_5050special_db, disckopkecil_5050special_db, disckepalaganjil_5050special_db, disckepalagenap_5050special_db    float32
		disckepalabesar_5050special_db, disckepalakecil_5050special_db, discekorganjil_5050special_db, discekorgenap_5050special_db, discekorbesar_5050special_db, discekorkecil_5050special_db                                                                                                                  float32
		limitglobal_5050special_db, limittotal_5050special_db                                                                                                                                                                                                                                                    float32
		minbet_5050kombinasi_db, maxbet_5050kombinasi_db                                                                                                                                                                                                                                                         float32
		belakangkeimono_5050kombinasi_db, belakangkeistereo_5050kombinasi_db, belakangkeikembang_5050kombinasi_db, belakangkeikempis_5050kombinasi_db, belakangkeikembar_5050kombinasi_db                                                                                                                        float32
		tengahkeimono_5050kombinasi_db, tengahkeistereo_5050kombinasi_db, tengahkeikembang_5050kombinasi_db, tengahkeikempis_5050kombinasi_db, tengahkeikembar_5050kombinasi_db                                                                                                                                  float32
		depankeimono_5050kombinasi_db, depankeistereo_5050kombinasi_db, depankeikembang_5050kombinasi_db, depankeikempis_5050kombinasi_db, depankeikembar_5050kombinasi_db                                                                                                                                       float32
		belakangdiscmono_5050kombinasi_db, belakangdiscstereo_5050kombinasi_db, belakangdisckembang_5050kombinasi_db, belakangdisckempis_5050kombinasi_db, belakangdisckembar_5050kombinasi_db                                                                                                                   float32
		tengahdiscmono_5050kombinasi_db, tengahdiscstereo_5050kombinasi_db, tengahdisckembang_5050kombinasi_db, tengahdisckempis_5050kombinasi_db, tengahdisckembar_5050kombinasi_db                                                                                                                             float32
		depandiscmono_5050kombinasi_db, depandiscstereo_5050kombinasi_db, depandisckembang_5050kombinasi_db, depandisckempis_5050kombinasi_db, depandisckembar_5050kombinasi_db                                                                                                                                  float32
		limitglobal_5050kombinasi_db, limittotal_5050kombinasi_db                                                                                                                                                                                                                                                float32
		minbet_kombinasi_db, maxbet_kombinasi_db, win_kombinasi_db, disc_kombinasi_db, limitglobal_kombinasi_db, limittotal_kombinasi_db                                                                                                                                                                         float32
		minbet_dasar_db, maxbet_dasar_db, keibesar_dasar_db, keikecil_dasar_db, keigenap_dasar_db, keiganjil_dasar_db, discbesar_dasar_db, disckecil_dasar_db, discgenap_dasar_db, discganjil_dasar_db, limitglobal_dasar_db, limittotal_dasar_db                                                                float32
		minbet_shio_db, maxbet_shio_db, win_shio_db, disc_shio_db, limitglobal_shio_db, limittotal_shio_db                                                                                                                                                                                                       float32
		shioyear_shio_db                                                                                                                                                                                                                                                                                         string
	)
	rows := con.QueryRowContext(ctx, sql_detail, idcomppasaran, company)
	switch err := rows.Scan(
		&pasarandiundi_db, &pasaranurl_db, &jamtutup_db, &jamjadwal_db, &jamopen_db, &statuspasaranactive_db,
		&limitline_4d_db, &limitline_3d_db, &limitline_2d_db, &limitline_2dd_db, &limitline_2dt_db, &bbfs_db,
		&minbet_432d_db, &maxbet4d_432d_db, &maxbet3d_432d_db, &maxbet2d_432d_db, &maxbet2dd_432d_db, &maxbet2dt_432d_db,
		&limitotal4d_432d_db, &limitotal3d_432d_db, &limitotal2d_432d_db, &limitotal2dd_432d_db, &limitotal2dt_432d_db,
		&limitglobal4d_432d_db, &limitglobal3d_432d_db, &limitglobal2d_432d_db, &limitglobal2dd_432d_db, &limitglobal2dt_432d_db,
		&disc4d_432d_db, &disc3d_432d_db, &disc2d_432d_db, &disc2dd_432d_db, &disc2dt_432d_db,
		&win4d_432d_db, &win3d_432d_db, &win2d_432d_db, &win2dd_432d_db, &win2dt_432d_db,
		&minbet_cbebas_db, &maxbet_cbebas_db, &win_cbebas_db, &disc_cbebas_db, &limitglobal_cbebas_db, &limittotal_cbebas_db,
		&minbet_cmacau_db, &maxbet_cmacau_db, &win2d_cmacau_db, &win3d_cmacau_db, &win4d_cmacau_db, &disc_cmacau_db, &limitglobal_cmacau_db, &limitotal_cmacau_db,
		&minbet_cnaga_db, &maxbet_cnaga_db, &win3_cnaga_db, &win4_cnaga_db, &disc_cnaga_db, &limitglobal_cnaga_db, &limittotal_cnaga_db,
		&minbet_cjitu_db, &maxbet_cjitu_db,
		&winas_cjitu_db, &winkop_cjitu_db, &winkepala_cjitu_db, &winekor_cjitu_db,
		&desc_cjitu_db, &limitglobal_cjitu_db, &limittotal_cjitu_db,
		&minbet_5050umum_db, &maxbet_5050umum_db,
		&keibesar_5050umum_db, &keikecil_5050umum_db, &keigenap_5050umum_db, &keiganjil_5050umum_db, &keitengah_5050umum_db, &keitepi_5050umum_db,
		&discbesar_5050umum_db, &disckecil_5050umum_db, &discgenap_5050umum_db, &discganjil_5050umum_db, &disctengah_5050umum_db, &disctepi_5050umum_db, &limitglobal_5050umum_db, &limittotal_5050umum_db,
		&minbet_5050special_db, &maxbet_5050special_db, &keiasganjil_5050special_db, &keiasgenap_5050special_db, &keiasbesar_5050special_db, &keiaskecil_5050special_db, &keikopganjil_5050special_db, &keikopgenap_5050special_db,
		&keikopbesar_5050special_db, &keikopkecil_5050special_db, &keikepalaganjil_5050special_db, &keikepalagenap_5050special_db, &keikepalabesar_5050special_db, &keikepalakecil_5050special_db, &keiekorganjil_5050special_db, &keiekorgenap_5050special_db, &keiekorbesar_5050special_db, &keiekorkecil_5050special_db,
		&discasganjil_5050special_db, &discasgenap_5050special_db, &discasbesar_5050special_db, &discaskecil_5050special_db, &disckopganjil_5050special_db, &disckopgenap_5050special_db, &disckopbesar_5050special_db, &disckopkecil_5050special_db, &disckepalaganjil_5050special_db, &disckepalagenap_5050special_db,
		&disckepalabesar_5050special_db, &disckepalakecil_5050special_db, &discekorganjil_5050special_db, &discekorgenap_5050special_db, &discekorbesar_5050special_db, &discekorkecil_5050special_db, &limitglobal_5050special_db, &limittotal_5050special_db,
		&minbet_5050kombinasi_db, &maxbet_5050kombinasi_db,
		&belakangkeimono_5050kombinasi_db, &belakangkeistereo_5050kombinasi_db, &belakangkeikembang_5050kombinasi_db, &belakangkeikempis_5050kombinasi_db, &belakangkeikembar_5050kombinasi_db,
		&tengahkeimono_5050kombinasi_db, &tengahkeistereo_5050kombinasi_db, &tengahkeikembang_5050kombinasi_db, &tengahkeikempis_5050kombinasi_db, &tengahkeikembar_5050kombinasi_db,
		&depankeimono_5050kombinasi_db, &depankeistereo_5050kombinasi_db, &depankeikembang_5050kombinasi_db, &depankeikempis_5050kombinasi_db, &depankeikembar_5050kombinasi_db,
		&belakangdiscmono_5050kombinasi_db, &belakangdiscstereo_5050kombinasi_db, &belakangdisckembang_5050kombinasi_db, &belakangdisckempis_5050kombinasi_db, &belakangdisckembar_5050kombinasi_db,
		&tengahdiscmono_5050kombinasi_db, &tengahdiscstereo_5050kombinasi_db, &tengahdisckembang_5050kombinasi_db, &tengahdisckempis_5050kombinasi_db, &tengahdisckembar_5050kombinasi_db,
		&depandiscmono_5050kombinasi_db, &depandiscstereo_5050kombinasi_db, &depandisckembang_5050kombinasi_db, &depandisckempis_5050kombinasi_db, &depandisckembar_5050kombinasi_db,
		&limitglobal_5050kombinasi_db, &limittotal_5050kombinasi_db,
		&minbet_kombinasi_db, &maxbet_kombinasi_db, &win_kombinasi_db, &disc_kombinasi_db, &limitglobal_kombinasi_db, &limittotal_kombinasi_db,
		&minbet_dasar_db, &maxbet_dasar_db, &keibesar_dasar_db, &keikecil_dasar_db, &keigenap_dasar_db, &keiganjil_dasar_db, &discbesar_dasar_db, &disckecil_dasar_db, &discgenap_dasar_db, &discganjil_dasar_db, &limitglobal_dasar_db, &limittotal_dasar_db,
		&minbet_shio_db, &maxbet_shio_db, &win_shio_db, &disc_shio_db, &limitglobal_shio_db, &limittotal_shio_db, &shioyear_shio_db); err {
	case sql.ErrNoRows:
		flag = false
	case nil:
		obj.Company_Pasaran_diundi = pasarandiundi_db
		obj.Company_Pasaran_url = pasaranurl_db
		obj.Company_Pasaran_jamtutup = jamtutup_db
		obj.Company_Pasaran_jamjadwal = jamjadwal_db
		obj.Company_Pasaran_jamopen = jamopen_db
		obj.Company_Pasaran_statusactive = statuspasaranactive_db
		obj.Company_Limitline4d = limitline_4d_db
		obj.Company_Limitline3d = limitline_3d_db
		obj.Company_Limitline2d = limitline_2d_db
		obj.Company_Limitline2dd = limitline_2dd_db
		obj.Company_Limitline2dt = limitline_2dt_db
		obj.Company_Bbfs = bbfs_db
		obj.Company_Minbet_432d = minbet_432d_db
		obj.Company_Maxbet4d_432d = maxbet4d_432d_db
		obj.Company_Maxbet3d_432d = maxbet3d_432d_db
		obj.Company_Maxbet2d_432d = maxbet2d_432d_db
		obj.Company_Maxbet2dd_432d = maxbet2dd_432d_db
		obj.Company_Maxbet2dt_432d = maxbet2dt_432d_db
		obj.Company_Limitotal4d_432d = limitotal4d_432d_db
		obj.Company_Limitotal3d_432d = limitotal3d_432d_db
		obj.Company_Limitotal2d_432d = limitotal2d_432d_db
		obj.Company_Limitotal2dd_432d = limitotal2dd_432d_db
		obj.Company_Limitotal2dt_432d = limitotal2dt_432d_db
		obj.Company_Limitglobal4d_432d = limitglobal4d_432d_db
		obj.Company_Limitglobal3d_432d = limitglobal3d_432d_db
		obj.Company_Limitglobal2d_432d = limitglobal2d_432d_db
		obj.Company_Limitglobal2dd_432d = limitglobal2dd_432d_db
		obj.Company_Limitglobal2dt_432d = limitglobal2dt_432d_db
		obj.Company_Disc4d_432d = disc4d_432d_db
		obj.Company_Disc3d_432d = disc3d_432d_db
		obj.Company_Disc2d_432d = disc2d_432d_db
		obj.Company_Disc2dd_432d = disc2dd_432d_db
		obj.Company_Disc2dt_432d = disc2dt_432d_db
		obj.Company_Win4d_432d = win4d_432d_db
		obj.Company_Win3d_432d = win3d_432d_db
		obj.Company_Win2d_432d = win2d_432d_db
		obj.Company_Win2dd_432d = win2dd_432d_db
		obj.Company_Win2dt_432d = win2dt_432d_db
		obj.Company_Minbet_cbebas = minbet_cbebas_db
		obj.Company_Maxbet_cbebas = maxbet_cbebas_db
		obj.Company_Win_cbebas = win_cbebas_db
		obj.Company_Disc_cbebas = disc_cbebas_db
		obj.Company_Limitglobal_cbebas = limitglobal_cbebas_db
		obj.Company_Limittotal_cbebas = limittotal_cbebas_db
		obj.Company_Minbet_cmacau = minbet_cmacau_db
		obj.Company_Maxbet_cmacau = maxbet_cmacau_db
		obj.Company_Win2d_cmacau = win2d_cmacau_db
		obj.Company_Win3d_cmacau = win3d_cmacau_db
		obj.Company_Win4d_cmacau = win4d_cmacau_db
		obj.Company_Disc_cmacau = disc_cmacau_db
		obj.Company_Limitglobal_cmacau = limitglobal_cmacau_db
		obj.Company_Limitotal_cmacau = limitotal_cmacau_db
		obj.Company_Minbet_cnaga = minbet_cnaga_db
		obj.Company_Maxbet_cnaga = maxbet_cnaga_db
		obj.Company_Win3_cnaga = win3_cnaga_db
		obj.Company_Win4_cnaga = win4_cnaga_db
		obj.Company_Disc_cnaga = disc_cnaga_db
		obj.Company_Limitglobal_cnaga = limitglobal_cnaga_db
		obj.Company_Limittotal_cnaga = limittotal_cnaga_db
		obj.Company_Minbet_cjitu = minbet_cjitu_db
		obj.Company_Maxbet_cjitu = maxbet_cjitu_db
		obj.Company_Winas_cjitu = winas_cjitu_db
		obj.Company_Winkop_cjitu = winkop_cjitu_db
		obj.Company_Winkepala_cjitu = winkepala_cjitu_db
		obj.Company_Winekor_cjitu = winekor_cjitu_db
		obj.Company_Desc_cjitu = desc_cjitu_db
		obj.Company_Limitglobal_cjitu = limitglobal_cjitu_db
		obj.Company_Limittotal_cjitu = limittotal_cjitu_db
		obj.Company_Minbet_5050umum = minbet_5050umum_db
		obj.Company_Maxbet_5050umum = maxbet_5050umum_db
		obj.Company_Keibesar_5050umum = keibesar_5050umum_db
		obj.Company_Keikecil_5050umum = keikecil_5050umum_db
		obj.Company_Keigenap_5050umum = keigenap_5050umum_db
		obj.Company_Keiganjil_5050umum = keiganjil_5050umum_db
		obj.Company_Keitengah_5050umum = keitengah_5050umum_db
		obj.Company_Keitepi_5050umum = keitengah_5050umum_db
		obj.Company_Discbesar_5050umum = discbesar_5050umum_db
		obj.Company_Disckecil_5050umum = disckecil_5050umum_db
		obj.Company_Discgenap_5050umum = discgenap_5050umum_db
		obj.Company_Discganjil_5050umum = discganjil_5050umum_db
		obj.Company_Disctengah_5050umum = disctengah_5050umum_db
		obj.Company_Disctepi_5050umum = disctepi_5050umum_db
		obj.Company_Limitglobal_5050umum = limitglobal_5050umum_db
		obj.Company_Limittotal_5050umum = limittotal_5050umum_db
		obj.Company_Minbet_5050special = minbet_5050special_db
		obj.Company_Maxbet_5050special = maxbet_5050special_db
		obj.Company_Keiasganjil_5050special = keiasganjil_5050special_db
		obj.Company_Keiasgenap_5050special = keiasgenap_5050special_db
		obj.Company_Keiasbesar_5050special = keiasbesar_5050special_db
		obj.Company_Keiaskecil_5050special = keiaskecil_5050special_db
		obj.Company_Keikopganjil_5050special = keikopganjil_5050special_db
		obj.Company_Keikopgenap_5050special = keikopgenap_5050special_db
		obj.Company_Keikopbesar_5050special = keikopbesar_5050special_db
		obj.Company_Keikopkecil_5050special = keikopkecil_5050special_db
		obj.Company_Keikepalaganjil_5050special = keikepalaganjil_5050special_db
		obj.Company_Keikepalagenap_5050special = keikepalagenap_5050special_db
		obj.Company_Keikepalabesar_5050special = keikepalabesar_5050special_db
		obj.Company_Keikepalakecil_5050special = keikepalakecil_5050special_db
		obj.Company_Keiekorganjil_5050special = keiekorganjil_5050special_db
		obj.Company_Keiekorgenap_5050special = keiekorgenap_5050special_db
		obj.Company_Keiekorbesar_5050special = keiekorbesar_5050special_db
		obj.Company_Keiekorkecil_5050special = keiekorkecil_5050special_db
		obj.Company_Discasganjil_5050special = discasganjil_5050special_db
		obj.Company_Discasgenap_5050special = discasgenap_5050special_db
		obj.Company_Discasbesar_5050special = discasbesar_5050special_db
		obj.Company_Discaskecil_5050special = discaskecil_5050special_db
		obj.Company_Disckopganjil_5050special = disckopganjil_5050special_db
		obj.Company_Disckopgenap_5050special = disckopgenap_5050special_db
		obj.Company_Disckopbesar_5050special = disckopbesar_5050special_db
		obj.Company_Disckopkecil_5050special = disckopkecil_5050special_db
		obj.Company_Disckepalaganjil_5050special = disckepalaganjil_5050special_db
		obj.Company_Disckepalagenap_5050special = disckepalagenap_5050special_db
		obj.Company_Disckepalabesar_5050special = disckepalabesar_5050special_db
		obj.Company_Disckepalakecil_5050special = disckepalakecil_5050special_db
		obj.Company_Discekorganjil_5050special = discekorganjil_5050special_db
		obj.Company_Discekorgenap_5050special = discekorgenap_5050special_db
		obj.Company_Discekorbesar_5050special = discekorbesar_5050special_db
		obj.Company_Discekorkecil_5050special = discekorkecil_5050special_db
		obj.Company_Limitglobal_5050special = limitglobal_5050special_db
		obj.Company_Limittotal_5050special = limittotal_5050special_db
		obj.Company_Minbet_5050kombinasi = minbet_5050kombinasi_db
		obj.Company_Maxbet_5050kombinasi = maxbet_5050kombinasi_db
		obj.Company_Belakangkeimono_5050kombinasi = belakangkeimono_5050kombinasi_db
		obj.Company_Belakangkeistereo_5050kombinasi = belakangkeistereo_5050kombinasi_db
		obj.Company_Belakangkeikembang_5050kombinasi = belakangkeikembang_5050kombinasi_db
		obj.Company_Belakangkeikempis_5050kombinasi = belakangkeikempis_5050kombinasi_db
		obj.Company_Belakangkeikembar_5050kombinasi = belakangkeikembang_5050kombinasi_db
		obj.Company_Tengahkeimono_5050kombinasi = tengahkeimono_5050kombinasi_db
		obj.Company_Tengahkeistereo_5050kombinasi = tengahkeistereo_5050kombinasi_db
		obj.Company_Tengahkeikembang_5050kombinasi = tengahkeikembang_5050kombinasi_db
		obj.Company_Tengahkeikempis_5050kombinasi = tengahkeikempis_5050kombinasi_db
		obj.Company_Tengahkeikembar_5050kombinasi = tengahkeikembar_5050kombinasi_db
		obj.Company_Depankeimono_5050kombinasi = depankeimono_5050kombinasi_db
		obj.Company_Depankeistereo_5050kombinasi = depankeistereo_5050kombinasi_db
		obj.Company_Depankeikembang_5050kombinasi = depankeikembang_5050kombinasi_db
		obj.Company_Depankeikempis_5050kombinasi = depankeikempis_5050kombinasi_db
		obj.Company_Depankeikembar_5050kombinasi = depankeikembar_5050kombinasi_db
		obj.Company_Belakangdiscmono_5050kombinasi = belakangdiscmono_5050kombinasi_db
		obj.Company_Belakangdiscstereo_5050kombinasi = belakangdiscstereo_5050kombinasi_db
		obj.Company_Belakangdisckembang_5050kombinasi = belakangdisckembang_5050kombinasi_db
		obj.Company_Belakangdisckempis_5050kombinasi = belakangdisckempis_5050kombinasi_db
		obj.Company_Belakangdisckembar_5050kombinasi = belakangdisckembang_5050kombinasi_db
		obj.Company_Tengahdiscmono_5050kombinasi = tengahdiscmono_5050kombinasi_db
		obj.Company_Tengahdiscstereo_5050kombinasi = tengahdiscstereo_5050kombinasi_db
		obj.Company_Tengahdisckembang_5050kombinasi = tengahdisckembang_5050kombinasi_db
		obj.Company_Tengahdisckempis_5050kombinasi = tengahdisckempis_5050kombinasi_db
		obj.Company_Tengahdisckembar_5050kombinasi = tengahdisckembar_5050kombinasi_db
		obj.Company_Depandiscmono_5050kombinasi = depandiscstereo_5050kombinasi_db
		obj.Company_Depandiscstereo_5050kombinasi = depandiscstereo_5050kombinasi_db
		obj.Company_Depandisckembang_5050kombinasi = depandisckembang_5050kombinasi_db
		obj.Company_Depandisckempis_5050kombinasi = depandisckempis_5050kombinasi_db
		obj.Company_Depandisckembar_5050kombinasi = depandisckembang_5050kombinasi_db
		obj.Company_Limitglobal_5050kombinasi = limitglobal_5050kombinasi_db
		obj.Company_Limittotal_5050kombinasi = limittotal_5050kombinasi_db
		obj.Company_Minbet_kombinasi = minbet_kombinasi_db
		obj.Company_Maxbet_kombinasi = maxbet_kombinasi_db
		obj.Company_Win_kombinasi = win_kombinasi_db
		obj.Company_Disc_kombinasi = disc_kombinasi_db
		obj.Company_Limitglobal_kombinasi = limitglobal_kombinasi_db
		obj.Company_Limittotal_kombinasi = limittotal_kombinasi_db
		obj.Company_Minbet_dasar = minbet_dasar_db
		obj.Company_Maxbet_dasar = maxbet_dasar_db
		obj.Company_Keibesar_dasar = keibesar_dasar_db
		obj.Company_Keikecil_dasar = keikecil_dasar_db
		obj.Company_Keigenap_dasar = keigenap_dasar_db
		obj.Company_Keiganjil_dasar = keiganjil_dasar_db
		obj.Company_Discbesar_dasar = discbesar_dasar_db
		obj.Company_Disckecil_dasar = disckecil_dasar_db
		obj.Company_Discgenap_dasar = discgenap_dasar_db
		obj.Company_Discganjil_dasar = discganjil_dasar_db
		obj.Company_Limitglobal_dasar = limitglobal_dasar_db
		obj.Company_Limittotal_dasar = limittotal_dasar_db
		obj.Company_Minbet_shio = minbet_shio_db
		obj.Company_Maxbet_shio = maxbet_shio_db
		obj.Company_Win_shio = win_shio_db
		obj.Company_Disc_shio = disc_shio_db
		obj.Company_Shioyear_shio = shioyear_shio_db
		obj.Company_Limitglobal_shio = limitglobal_shio_db
		obj.Company_Limittotal_shio = limittotal_shio_db

		arraobj = append(arraobj, obj)
		msg = "Success"
		flag = true
	default:
		flag = false
		helpers.ErrorCheck(err)
	}
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
func Fetch_company_listpasaranonline(company string, idcomppasaran int) (helpers.Response, error) {
	var obj sCompanyListPasaranOnline
	var arraobj []sCompanyListPasaranOnline
	var res helpers.Response
	msg := "Error"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	var no int = 0

	sql_periode := `SELECT 
			idcomppasaranoff  , haripasaran 
			FROM ` + config.DB_tbl_mst_company_game_pasaran_offline + ` 
			WHERE idcompany=? AND idcomppasaran=? 
		`

	row, err := con.QueryContext(ctx, sql_periode, company, idcomppasaran)
	helpers.ErrorCheck(err)
	for row.Next() {
		no++
		var (
			idcomppasaranoff int
			haripasaran      string
		)

		err = row.Scan(&idcomppasaranoff, &haripasaran)
		helpers.ErrorCheck(err)

		obj.Pasaran_onlineid = idcomppasaranoff
		obj.Pasaran_harian = haripasaran

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
func Fetch_company_listkeluaran(company, periode string, pasaran int) (helpers.ResponseListKeluaran, error) {
	var obj scompanylistkeluaran
	var arraobj []scompanylistkeluaran
	var res helpers.ResponseListKeluaran
	msg := "Error"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	_, startmonthyear, endmonthyear := helpers.GetEndRangeDate(periode)
	var no int = 0
	subtotalwinlose := 0
	tbl_trx_keluaran, _, _, _ := Get_mappingdatabase(company)

	sql_periode := `SELECT 
			A.idtrxkeluaran, A.idcomppasaran, A.keluaranperiode, A.datekeluaran, A.keluarantogel, 
			A.total_member, A.total_bet, A.total_outstanding, A.winlose, A.total_cancel, 
			C.nmpasarantogel, B.idpasarantogel, A.revisi, A.noterevisi
			FROM ` + tbl_trx_keluaran + ` as A 
			JOIN ` + config.DB_tbl_mst_company_game_pasaran + ` as B ON B.idcomppasaran = A.idcomppasaran  
			JOIN ` + config.DB_tbl_mst_pasaran_togel + ` as C ON C.idpasarantogel  = B.idpasarantogel  
			WHERE A.idcompany = ? 
			AND A.idcomppasaran = ? 
			AND A.datekeluaran >= ? 
			AND A.datekeluaran <= ? 
			ORDER BY A.datekeluaran DESC 
		`

	row, err := con.QueryContext(ctx, sql_periode, company, pasaran, startmonthyear, endmonthyear)
	helpers.ErrorCheck(err)
	for row.Next() {
		no++
		var (
			idtrxkeluaran_db, idcomppasaran_db, revisi_db                                                              int
			datekeluaran_db, keluarantogel_db, nmpasarantogel_db, idpasarantogel_db, keluaranperiode_db, noterevisi_db string
			total_member_db, total_bet_db, total_outstanding_db, winlose_db, total_cancel_db                           float32
		)

		err = row.Scan(
			&idtrxkeluaran_db, &idcomppasaran_db, &keluaranperiode_db,
			&datekeluaran_db, &keluarantogel_db, &total_member_db,
			&total_bet_db, &total_outstanding_db, &winlose_db, &total_cancel_db,
			&nmpasarantogel_db, &idpasarantogel_db, &revisi_db, &noterevisi_db)
		helpers.ErrorCheck(err)
		status := "DONE"
		status_css := config.STATUS_COMPLETE
		if keluarantogel_db == "" {
			status = "RUNNING"
			status_css = config.STATUS_RUNNING
		}
		totalwinlose := total_outstanding_db - total_cancel_db - winlose_db
		subtotalwinlose = subtotalwinlose + int(totalwinlose)
		obj.No = no
		obj.Idtrxkeluaran = idtrxkeluaran_db
		obj.Idcomppasaran = idcomppasaran_db
		obj.Pasarancode = idpasarantogel_db
		obj.Nmpasaran = nmpasarantogel_db
		obj.Keluaranperiode = keluaranperiode_db + "-" + idpasarantogel_db
		obj.Tanggalperiode = datekeluaran_db
		obj.Keluarantogel = keluarantogel_db
		obj.Total_Member = total_member_db
		obj.Total_bet = total_bet_db
		obj.Total_outstanding = total_outstanding_db
		obj.Total_cancelbet = total_cancel_db
		obj.Winlose = totalwinlose
		obj.Winlosetemp = _winlosetemp(company, "", "", idtrxkeluaran_db)
		obj.Revisi = revisi_db
		obj.Msgrevisi = noterevisi_db
		obj.Status = status
		obj.Status_css = status_css
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Totalwinlose = subtotalwinlose
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Fetch_company_invoice_member(company string, invoice int) (helpers.Response, error) {
	var obj invoicelistMember
	var arraobj []invoicelistMember
	var res helpers.Response
	msg := "Error"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	_, tbl_trx_keluarandetail, _, _ := Get_mappingdatabase(company)

	sql := `SELECT 
		username, 
		count(username) as totalbet, 
		sum(bet-(bet*diskon)-(bet*kei)) as totalbayar,
		sum(cancelbet) as totalcancel,  
		sum(winhasil) as totalwin 
		FROM ` + tbl_trx_keluarandetail + ` 
		WHERE idcompany = ? 
		AND idtrxkeluaran = ? 
		GROUP BY username 
	`
	row, err := con.QueryContext(ctx, sql, company, invoice)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			totalbet_db, totalbayar_db, totalwin_db, totalcancel_db float64
			username_db                                             string
		)

		err = row.Scan(
			&username_db,
			&totalbet_db,
			&totalbayar_db,
			&totalcancel_db,
			&totalwin_db)

		helpers.ErrorCheck(err)

		obj.Member = username_db
		obj.Totalbet = int(totalbet_db)
		obj.Totalbayar = int(totalbayar_db)
		obj.Totalcancelbet = int(totalcancel_db)
		obj.Totalwin = int(totalwin_db)
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
func Fetch_company_invoice_membertemp(company string, invoice int) (helpers.Response, error) {
	var obj invoicelistMember
	var arraobj []invoicelistMember
	var res helpers.Response
	msg := "Error"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	_, _, tbl_trx_keluaranmember, _ := Get_mappingdatabase(company)

	sql := `SELECT 
		username, 
		totalbet, 
		totalbayar,
		totalcancel,  
		totalwin 
		FROM ` + tbl_trx_keluaranmember + ` 
		WHERE idcompany = ? 
		AND idtrxkeluaran = ? 
	`
	row, err := con.QueryContext(ctx, sql, company, invoice)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			totalbet_db, totalbayar_db, totalwin_db, totalcancel_db float64
			username_db                                             string
		)

		err = row.Scan(
			&username_db,
			&totalbet_db,
			&totalbayar_db,
			&totalcancel_db,
			&totalwin_db)

		helpers.ErrorCheck(err)

		obj.Member = username_db
		obj.Totalbet = int(totalbet_db)
		obj.Totalbayar = int(totalbayar_db)
		obj.Totalcancelbet = int(totalcancel_db)
		obj.Totalwin = int(totalwin_db)
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
func Fetch_company_invoice_membersync(company string, invoice int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	year := tglnow.Format("YYYY")
	month := tglnow.Format("MM")
	render_page := time.Now()
	_, tbl_trx_keluarandetail, tbl_trx_keluaranmember, _ := Get_mappingdatabase(company)
	flag := false
	//DELETE MEMBER
	stmt_keluarantogelmember, e_member := con.PrepareContext(ctx, `
		DELETE FROM  
		`+tbl_trx_keluaranmember+`   
		WHERE idtrxkeluaran=? AND idcompany=? 
	`)

	helpers.ErrorCheck(e_member)
	rec_keluarantogelmember, e_keluarantogelmember := stmt_keluarantogelmember.ExecContext(ctx, invoice, company)
	helpers.ErrorCheck(e_keluarantogelmember)
	affect_keluarantogelmember, err_affer_keluarantogelmember := rec_keluarantogelmember.RowsAffected()
	helpers.ErrorCheck(err_affer_keluarantogelmember)

	defer stmt_keluarantogelmember.Close()
	if affect_keluarantogelmember > 0 {
		flag = true
		msg = "Success"
		log.Printf("Delete tbl_trx_keluarantogel_member : %d\n", invoice)
	} else {
		flag = false
		log.Println("Delete tbl_trx_keluarantogel_member failed")
	}
	if flag {
		sql := `SELECT 
			username, 
			count(username) as totalbet, 
			sum(bet-(bet*diskon)-(bet*kei)) as totalbayar,
			sum(cancelbet) as totalcancel,  
			sum(winhasil) as totalwin 
			FROM ` + tbl_trx_keluarandetail + ` 
			WHERE idcompany = ? 
			AND idtrxkeluaran = ? 
			GROUP BY username 
		`
		row, err := con.QueryContext(ctx, sql, company, invoice)
		helpers.ErrorCheck(err)
		for row.Next() {
			var (
				totalbet_db, totalbayar_db, totalwin_db, totalcancel_db float64
				username_db                                             string
			)

			err = row.Scan(
				&username_db,
				&totalbet_db,
				&totalbayar_db,
				&totalcancel_db,
				&totalwin_db)

			helpers.ErrorCheck(err)
			field_col2 := tbl_trx_keluaranmember + year + month
			idkeluaranmember_counter := Get_counter(field_col2)
			idkeluaranmember := year + month + strconv.Itoa(idkeluaranmember_counter)
			stmt_keluaranmember, e_keluaranmember := con.PrepareContext(ctx, `
				insert into
				`+tbl_trx_keluaranmember+` (
					idkeluaranmember, idtrxkeluaran, idcompany,
					username, totalbet, totalbayar, totalwin, totalcancel, 
					createkeluaranmember, createdatekeluaranmember
				) values (
					?, ?, ?,
					?, ?, ?, ?, ?, 
					?, ?
				)
			`)
			helpers.ErrorCheck(e_keluaranmember)
			defer stmt_keluaranmember.Close()
			res_keluaranmember, e_keluaranmember := stmt_keluaranmember.ExecContext(ctx,
				idkeluaranmember,
				invoice,
				company,
				username_db,
				totalbet_db,
				totalbayar_db,
				totalwin_db,
				totalcancel_db,
				"SYSTEM",
				tglnow.Format("YYYY-MM-DD HH:mm:ss"))
			helpers.ErrorCheck(e_keluaranmember)
			insert_keluaranmember2, e_keluaranmember2 := res_keluaranmember.RowsAffected()
			helpers.ErrorCheck(e_keluaranmember2)
			log.Println("Affected :", insert_keluaranmember2)
			if insert_keluaranmember2 > 0 {
				msg = "Success"
				log.Println("Data Member tbl_trx_keluarantogel_member Berhasil di save ")
			} else {

				log.Println("Data Member tbl_trx_keluarantogel_member failed ")
			}
		}
		defer row.Close()
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Fetch_company_invoice_grouppermainan(company, username string, invoice int) (helpers.Response, error) {
	var obj invoicelistGroupPermainan
	var arraobj []invoicelistGroupPermainan
	var res helpers.Response
	msg := "Error"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	_, tbl_trx_keluarandetail, _, _ := Get_mappingdatabase(company)

	if username == "" {
		sql := `SELECT 
			typegame 
			FROM ` + tbl_trx_keluarandetail + `  
			WHERE idcompany = ? 
			AND idtrxkeluaran = ? 
			GROUP BY typegame 
		`
		row, err := con.QueryContext(ctx, sql, company, invoice)
		helpers.ErrorCheck(err)
		for row.Next() {
			var typegame_db string

			err = row.Scan(&typegame_db)
			helpers.ErrorCheck(err)
			obj.Permainan = typegame_db
			arraobj = append(arraobj, obj)
			msg = "Success"
		}
		defer row.Close()

		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = arraobj
		res.Time = time.Since(render_page).String()

		return res, nil
	} else {
		sql := `SELECT 
			typegame 
			FROM ` + tbl_trx_keluarandetail + `  
			WHERE idcompany = ? 
			AND idtrxkeluaran = ? 
			AND username = ? 
			GROUP BY typegame 
		`
		row, err := con.QueryContext(ctx, sql, company, invoice, username)
		helpers.ErrorCheck(err)
		for row.Next() {
			var typegame_db string

			err = row.Scan(&typegame_db)
			helpers.ErrorCheck(err)
			obj.Permainan = typegame_db
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
}
func Fetch_company_invoice_listpermainan(company, permainan string, invoice int) (helpers.ResponseListPermainan, error) {
	var obj invoicelistpermainan
	var arraobj []invoicelistpermainan
	var res helpers.ResponseListPermainan
	msg := "Error"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	totalbet := 0
	subtotalbayar := 0
	subtotalwin := 0
	_, tbl_trx_keluarandetail, _, _ := Get_mappingdatabase(company)

	sqldetail := `SELECT
					idtrxkeluarandetail , datetimedetail, ipaddress, browsertogel, devicetogel,  username, typegame, nomortogel, 
					bet, diskon, win, kei, statuskeluarandetail , createkeluarandetail, 
					createdatekeluarandetail, updatekeluarandetail, updatedatekeluarandetail 
					FROM ` + tbl_trx_keluarandetail + ` 
					WHERE idcompany = ? 
					AND idtrxkeluaran = ? 
					AND typegame = ? 
					ORDER BY datetimedetail DESC 
				`
	row, err := con.QueryContext(ctx, sqldetail, company, invoice, permainan)

	helpers.ErrorCheck(err)
	for row.Next() {
		totalbet += 1
		var (
			idtrxkeluarandetail_db, bet_db                                                                                                      int
			datetimedetail_db, ipaddresss_db, username_db, typegame_db, nomortogel_db, browsertogel_db, devicetogel_db                          string
			statuskeluarandetail_db, createkeluarandetail_db, createdatekeluarandetail_db, updatekeluarandetail_db, updatedatekeluarandetail_db string
			diskon_db, win_db, kei_db                                                                                                           float32
		)

		err = row.Scan(
			&idtrxkeluarandetail_db,
			&datetimedetail_db, &ipaddresss_db, &browsertogel_db, &devicetogel_db, &username_db, &typegame_db, &nomortogel_db,
			&bet_db, &diskon_db, &win_db, &kei_db, &statuskeluarandetail_db, &createkeluarandetail_db,
			&createdatekeluarandetail_db, &updatekeluarandetail_db, &updatedatekeluarandetail_db)

		helpers.ErrorCheck(err)

		diskonpercen := diskon_db * 100
		diskonbet := int(float32(bet_db) * diskon_db)
		keipercen := kei_db * 100
		keibet := int(float32(bet_db) * kei_db)
		bayar := bet_db - int(float32(bet_db)*diskon_db) - int(float32(bet_db)*kei_db)
		subtotalbayar = subtotalbayar + bayar
		winhasil := _rumuswinhasil(typegame_db, bayar, bet_db, win_db)
		totalwin := 0

		status_css := ""
		switch statuskeluarandetail_db {
		case "RUNNING":
			totalwin = 0
			status_css = config.STATUS_RUNNING
		case "WINNER":
			totalwin = winhasil
			subtotalwin = subtotalwin + winhasil
			status_css = config.STATUS_COMPLETE
		case "LOSE":
			totalwin = 0
			status_css = config.STATUS_CANCEL
		case "CANCEL":
			totalwin = 0
			status_css = config.STATUS_CANCELBET
		}

		obj.Bet_id = idtrxkeluarandetail_db
		obj.Bet_datetime = datetimedetail_db
		obj.Bet_ipaddress = ipaddresss_db
		obj.Bet_device = devicetogel_db
		obj.Bet_timezone = browsertogel_db
		obj.Bet_username = username_db
		obj.Bet_typegame = typegame_db
		obj.Bet_nomortogel = nomortogel_db
		obj.Bet_bet = bet_db
		obj.Bet_diskon = diskonbet
		obj.Bet_diskonpercen = int(diskonpercen)
		obj.Bet_kei = keibet
		obj.Bet_keipercen = int(keipercen)
		obj.Bet_bayar = bayar
		obj.Bet_win = win_db
		obj.Bet_totalwin = totalwin
		obj.Bet_status = statuskeluarandetail_db
		obj.Bet_statuscss = status_css
		obj.Bet_create = createkeluarandetail_db
		obj.Bet_createDate = createdatekeluarandetail_db
		obj.Bet_update = updatekeluarandetail_db
		obj.Bet_updateDate = updatedatekeluarandetail_db
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Totalbet = totalbet
	res.Subtotal = subtotalbayar
	res.Subtotalwin = subtotalwin
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Fetch_company_invoice_listpermainanbystatus(company, status string, invoice int) (helpers.ResponseListPermainan, error) {
	var obj invoicelistpermainan
	var arraobj []invoicelistpermainan
	var res helpers.ResponseListPermainan
	msg := "Error"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	totalbet := 0
	subtotalbayar := 0
	subtotalwin := 0
	_, tbl_trx_keluarandetail, _, _ := Get_mappingdatabase(company)

	sqldetail := `SELECT
					idtrxkeluarandetail , datetimedetail, ipaddress, browsertogel, devicetogel,  username, typegame, nomortogel, 
					bet, diskon, win, kei, statuskeluarandetail , createkeluarandetail, 
					createdatekeluarandetail, updatekeluarandetail, updatedatekeluarandetail 
					FROM ` + tbl_trx_keluarandetail + ` 
					WHERE idcompany = ? 
					AND idtrxkeluaran = ? 
					AND statuskeluarandetail = ? 
					ORDER BY datetimedetail DESC 
				`
	row, err := con.QueryContext(ctx, sqldetail, company, invoice, status)

	helpers.ErrorCheck(err)

	for row.Next() {
		totalbet += 1
		var (
			idtrxkeluarandetail_db, bet_db                                                                                                      int
			datetimedetail_db, ipaddresss_db, username_db, typegame_db, nomortogel_db, browsertogel_db, devicetogel_db                          string
			statuskeluarandetail_db, createkeluarandetail_db, createdatekeluarandetail_db, updatekeluarandetail_db, updatedatekeluarandetail_db string
			diskon_db, win_db, kei_db                                                                                                           float32
		)

		err = row.Scan(
			&idtrxkeluarandetail_db,
			&datetimedetail_db, &ipaddresss_db, &browsertogel_db, &devicetogel_db, &username_db, &typegame_db, &nomortogel_db,
			&bet_db, &diskon_db, &win_db, &kei_db, &statuskeluarandetail_db, &createkeluarandetail_db,
			&createdatekeluarandetail_db, &updatekeluarandetail_db, &updatedatekeluarandetail_db)

		helpers.ErrorCheck(err)

		diskonpercen := diskon_db * 100
		diskonbet := int(float32(bet_db) * diskon_db)
		keipercen := kei_db * 100
		keibet := int(float32(bet_db) * kei_db)
		bayar := bet_db - int(float32(bet_db)*diskon_db) - int(float32(bet_db)*kei_db)
		subtotalbayar = subtotalbayar + bayar
		winhasil := _rumuswinhasil(typegame_db, bayar, bet_db, win_db)
		totalwin := 0

		status_css := ""
		switch statuskeluarandetail_db {
		case "RUNNING":
			totalwin = 0
			status_css = config.STATUS_RUNNING
		case "WINNER":
			totalwin = winhasil
			subtotalwin = subtotalwin + winhasil
			status_css = config.STATUS_COMPLETE
		case "LOSE":
			totalwin = 0
			status_css = config.STATUS_CANCEL
		case "CANCEL":
			totalwin = 0
			status_css = config.STATUS_CANCELBET
		}

		obj.Bet_id = idtrxkeluarandetail_db
		obj.Bet_datetime = datetimedetail_db
		obj.Bet_ipaddress = ipaddresss_db
		obj.Bet_device = devicetogel_db
		obj.Bet_timezone = browsertogel_db
		obj.Bet_username = username_db
		obj.Bet_typegame = typegame_db
		obj.Bet_nomortogel = nomortogel_db
		obj.Bet_bet = bet_db
		obj.Bet_diskon = diskonbet
		obj.Bet_diskonpercen = int(diskonpercen)
		obj.Bet_kei = keibet
		obj.Bet_keipercen = int(keipercen)
		obj.Bet_bayar = bayar
		obj.Bet_win = win_db
		obj.Bet_totalwin = totalwin
		obj.Bet_status = statuskeluarandetail_db
		obj.Bet_statuscss = status_css
		obj.Bet_create = createkeluarandetail_db
		obj.Bet_createDate = createdatekeluarandetail_db
		obj.Bet_update = updatekeluarandetail_db
		obj.Bet_updateDate = updatedatekeluarandetail_db
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()
	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(render_page).String()
	res.Totalbet = totalbet
	res.Subtotal = subtotalbayar
	res.Subtotalwin = subtotalwin
	return res, nil
}
func Fetch_company_invoice_listpermainanbyusername(company, username, permainan string, invoice int) (helpers.ResponseListPermainan, error) {
	var obj invoicelistpermainan
	var arraobj []invoicelistpermainan
	var res helpers.ResponseListPermainan
	msg := "Error"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	totalbet := 0
	subtotalbayar := 0
	subtotalcancel := 0
	subtotalwin := 0
	_, tbl_trx_keluarandetail, _, _ := Get_mappingdatabase(company)

	sqldetail := `SELECT
					idtrxkeluarandetail , datetimedetail, ipaddress, browsertogel, devicetogel,  username, typegame, nomortogel, 
					bet, diskon, win, kei, statuskeluarandetail , createkeluarandetail, 
					createdatekeluarandetail, updatekeluarandetail, updatedatekeluarandetail 
					FROM ` + tbl_trx_keluarandetail + ` 
					WHERE idcompany = ? 
					AND idtrxkeluaran = ? 
					AND username = ? 
					AND typegame = ? 
					ORDER BY datetimedetail DESC 
				`
	row, err := con.QueryContext(ctx, sqldetail, company, invoice, username, permainan)

	helpers.ErrorCheck(err)
	for row.Next() {
		totalbet += 1
		var (
			idtrxkeluarandetail_db, bet_db                                                                                                      int
			datetimedetail_db, ipaddresss_db, username_db, typegame_db, nomortogel_db, browsertogel_db, devicetogel_db                          string
			statuskeluarandetail_db, createkeluarandetail_db, createdatekeluarandetail_db, updatekeluarandetail_db, updatedatekeluarandetail_db string
			diskon_db, win_db, kei_db                                                                                                           float32
		)

		err = row.Scan(
			&idtrxkeluarandetail_db,
			&datetimedetail_db, &ipaddresss_db, &browsertogel_db, &devicetogel_db, &username_db, &typegame_db, &nomortogel_db,
			&bet_db, &diskon_db, &win_db, &kei_db, &statuskeluarandetail_db, &createkeluarandetail_db,
			&createdatekeluarandetail_db, &updatekeluarandetail_db, &updatedatekeluarandetail_db)

		helpers.ErrorCheck(err)

		diskonpercen := diskon_db * 100
		diskonbet := int(float32(bet_db) * diskon_db)
		keipercen := kei_db * 100
		keibet := int(float32(bet_db) * kei_db)
		bayar := bet_db - int(float32(bet_db)*diskon_db) - int(float32(bet_db)*kei_db)
		subtotalbayar = subtotalbayar + bayar
		winhasil := _rumuswinhasil(typegame_db, bayar, bet_db, win_db)
		totalwin := 0

		status_css := ""
		switch statuskeluarandetail_db {
		case "RUNNING":
			totalwin = 0
			status_css = config.STATUS_RUNNING
		case "WINNER":
			totalwin = winhasil
			subtotalwin = subtotalwin + winhasil
			status_css = config.STATUS_COMPLETE
		case "LOSE":
			totalwin = 0
			status_css = config.STATUS_CANCEL
		case "CANCEL":
			totalwin = 0
			subtotalcancel = subtotalcancel + bayar
			status_css = config.STATUS_CANCELBET
		}

		obj.Bet_id = idtrxkeluarandetail_db
		obj.Bet_datetime = datetimedetail_db
		obj.Bet_ipaddress = ipaddresss_db
		obj.Bet_device = devicetogel_db
		obj.Bet_timezone = browsertogel_db
		obj.Bet_username = username_db
		obj.Bet_typegame = typegame_db
		obj.Bet_nomortogel = nomortogel_db
		obj.Bet_bet = bet_db
		obj.Bet_diskon = diskonbet
		obj.Bet_diskonpercen = int(diskonpercen)
		obj.Bet_kei = keibet
		obj.Bet_keipercen = int(keipercen)
		obj.Bet_bayar = bayar
		obj.Bet_win = win_db
		obj.Bet_totalwin = totalwin
		obj.Bet_status = statuskeluarandetail_db
		obj.Bet_statuscss = status_css
		obj.Bet_create = createkeluarandetail_db
		obj.Bet_createDate = createdatekeluarandetail_db
		obj.Bet_update = updatekeluarandetail_db
		obj.Bet_updateDate = updatedatekeluarandetail_db
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Totalbet = totalbet
	res.Subtotal = subtotalbayar
	res.Subtotalcancel = subtotalcancel
	res.Subtotalwin = subtotalwin
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Save_company(sData, master, company, name, urldomain, status string) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	msg := "Failed"
	flag := false
	_, _, _, tbl_mst_company := Get_mappingdatabase(company)

	log.Println(tbl_mst_company)
	if sData == "New" {
		flag_insert := CheckDB(config.DB_tbl_mst_company, "idcompany", company)

		if !flag_insert {
			rows_insert, err_insert := con.PrepareContext(ctx, `
				INSERT INTO  
				`+config.DB_tbl_mst_company+` (
					idcompany, startjoincompany, idcurr, nmcompany, companyurl, statuscompany, createcompany, createdatecompany 
				)VALUES( 
					?,?,?,?,?,?,?,?
				) 
			`)
			helpers.ErrorCheck(err_insert)
			rec_comp, err_comp := rows_insert.ExecContext(ctx,
				company,
				tglnow.Format("YYYY-MM-DD HH:mm:ss"),
				"IDR",
				name,
				urldomain,
				status,
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

			newDB, err := CreateNewCompanyDB(tbl_mst_company, company, con)
			helpers.ErrorCheck(err)

			if newDB == "ok" {
				flag = true
				msg = "Success"
				log.Println("Database Berhasil di buat")
			}
		} else {
			msg = "Duplicate Entry"
		}
	} else {
		rows_update, err_update := con.PrepareContext(ctx, `
				UPDATE 
				`+config.DB_tbl_mst_company+`  
				SET nmcompany=?, companyurl=?, statuscompany=?,  
				updatecompany=?, updatedatecompany=? 
				WHERE idcompany=? 
			`)
		helpers.ErrorCheck(err_update)
		rec_comp, err_comp := rows_update.ExecContext(ctx,
			name,
			urldomain,
			status,
			master,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			company)
		helpers.ErrorCheck(err_comp)
		update_comp, err_comp := rec_comp.RowsAffected()
		helpers.ErrorCheck(err_comp)
		defer rows_update.Close()
		if update_comp > 0 {
			flag = true
			msg = "Success"
			log.Printf("Update %s Success : %s\n", config.DB_tbl_mst_company, name)
		} else {
			log.Printf("Update %s Failed \n", config.DB_tbl_mst_company)
		}
	}
	if flag {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	}

	return res, nil
}
func Save_companyNewAdmin(sData, master, company, username, password, name, status string) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	msg := "Failed"
	flag := false
	if sData == "New" {
		flag_insert := CheckDBTwoField(config.DB_tbl_mst_company_admin, "username_comp", username, "idcompany", company)

		if !flag_insert {
			hashpass := helpers.HashPasswordMD5(password)
			rows_insert, err_insert := con.PrepareContext(ctx, `
				INSERT INTO  
				`+config.DB_tbl_mst_company_admin+` (
					username_comp, password_comp, idcompany, typeadmin, nama_comp, status_comp, createcomp_admin, createdatecomp_admin 
				)VALUES( 
					?,?,?,?,?,?,?,? 
				) 
			`)
			helpers.ErrorCheck(err_insert)
			rec_comp, err_comp := rows_insert.ExecContext(ctx,
				username,
				hashpass,
				company,
				"MASTER",
				name,
				status,
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
	} else {
		if password == "" {
			rows_update, err_update := con.PrepareContext(ctx, `
				UPDATE 
				`+config.DB_tbl_mst_company_admin+`  
				SET nama_comp=?, status_comp=?, 
				updatecomp_admin=?, updatedatecomp_admin=? 
				WHERE idcompany=? AND username_comp=? 
			`)
			helpers.ErrorCheck(err_update)
			rec_comp, err_comp := rows_update.ExecContext(ctx,
				name,
				status,
				master,
				tglnow.Format("YYYY-MM-DD HH:mm:ss"),
				company, username)
			helpers.ErrorCheck(err_comp)
			update_comp, err_comp := rec_comp.RowsAffected()
			helpers.ErrorCheck(err_comp)
			defer rows_update.Close()
			if update_comp > 0 {
				flag = true
				msg = "Success"
				log.Printf("Update %s Success : %s\n", config.DB_tbl_mst_company_admin, name)
			} else {
				log.Printf("Update %s Failed \n", config.DB_tbl_mst_company_admin)
			}
		} else {
			hashpass := helpers.HashPasswordMD5(password)
			rows_update, err_update := con.PrepareContext(ctx, `
				UPDATE 
				`+config.DB_tbl_mst_company_admin+`  
				SET nama_comp=?, password_comp=?, status_comp=?,   
				updatecomp_admin=?, updatedatecomp_admin=? 
				WHERE idcompany=? AND username_comp=? 
			`)
			helpers.ErrorCheck(err_update)
			rec_comp, err_comp := rows_update.ExecContext(ctx,
				name,
				hashpass,
				status,
				master,
				tglnow.Format("YYYY-MM-DD HH:mm:ss"),
				company, username)
			helpers.ErrorCheck(err_comp)
			update_comp, err_comp := rec_comp.RowsAffected()
			helpers.ErrorCheck(err_comp)
			defer rows_update.Close()
			if update_comp > 0 {
				flag = true
				msg = "Success"
				log.Printf("Update %s Success : %s\n", config.DB_tbl_mst_company_admin, name)
			} else {
				log.Printf("Update %s Failed \n", config.DB_tbl_mst_company_admin)
			}
		}

	}
	if flag {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	}

	return res, nil
}
func Save_companyNewPasaran(master, company, pasarancode string) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	msg := "Failed"
	flag := false

	flag_insert := CheckDBTwoField(config.DB_tbl_mst_company_game_pasaran, "idcompany", company, "idpasarantogel", pasarancode)

	if !flag_insert {
		rows_insert, err_insert := con.PrepareContext(ctx, `
				INSERT INTO  
				`+config.DB_tbl_mst_company_game_pasaran+` (
					idcomppasaran, idcompany, idpasarantogel, 
					pasarandiundi, pasaranurl, jamtutup, jamjadwal, jamopen, statuspasaran, statuspasaranactive, 
					createcomppas, createdatecomppas 
				)VALUES( 
					?,?,?,?,?,?,?,?,?,?,?,?
				) 
			`)
		helpers.ErrorCheck(err_insert)

		sql_pasaran := `SELECT 
			urlpasaran, pasarandiundi, 
			jamtutup, jamjadwal, jamopen 
			FROM ` + config.DB_tbl_mst_pasaran_togel + ` 
			WHERE idpasarantogel = ? 
		`
		var (
			urlpasaran_db, pasarandiundi_db, jamtutup_db, jamjadwal_db, jamopen_db string
		)
		rows_select := con.QueryRowContext(ctx, sql_pasaran, pasarancode)
		switch err := rows_select.Scan(&urlpasaran_db, &pasarandiundi_db, &jamtutup_db, &jamjadwal_db, &jamopen_db); err {
		case sql.ErrNoRows:
			flag = false
		case nil:
			flag = true
		default:
			flag = false
			helpers.ErrorCheck(err)
		}
		if flag {
			year := tglnow.Format("YYYY")
			field_col := config.DB_tbl_mst_company_game_pasaran + year
			idcomppasaran_counter := Get_counter(field_col)
			idcomppasaran := tglnow.Format("YY") + strconv.Itoa(idcomppasaran_counter)
			rec_comp, err_comp := rows_insert.ExecContext(ctx,
				idcomppasaran,
				company,
				pasarancode,
				pasarandiundi_db,
				urlpasaran_db,
				jamtutup_db,
				jamjadwal_db,
				jamopen_db,
				"OFFLINE",
				"Y",
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
		}
	} else {
		msg = "Duplicate Entry"
	}
	if flag {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	}

	return res, nil
}
func Fetch_companyPasaranlimitline(master, company, pasarancode string, idcomppasaran int) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	msg := "Failed"
	flag := false

	sql_pasaran := `SELECT 
			limitline_4d, limitline_3d, limitline_2d, limitline_2dd, limitline_2dt, bbfs
			FROM ` + config.DB_tbl_mst_pasaran_togel + ` 
			WHERE idpasarantogel = ? 
		`
	var (
		limitline_4d_db, limitline_3d_db, limitline_2d_db, limitline_2dd_db, limitline_2dt_db, bbfs_db int
	)
	rows_select := con.QueryRowContext(ctx, sql_pasaran, pasarancode)
	switch err := rows_select.Scan(
		&limitline_4d_db, &limitline_3d_db, &limitline_2d_db, &limitline_2dd_db, &limitline_2dt_db, &bbfs_db); err {
	case sql.ErrNoRows:
		flag = false
	case nil:
		flag = true
	default:
		flag = false
		helpers.ErrorCheck(err)
	}

	if flag {
		sql_update := `
			UPDATE   
			` + config.DB_tbl_mst_company_game_pasaran + ` 
			SET limitline_4d=?, limitline_3d=?, limitline_2d=?, limitline_2dd=?, 
			limitline_2dt=?, bbfs=?  
			WHERE idcomppasaran=? AND idcompany=? 
		`
		rows_update, err_update := con.PrepareContext(ctx, sql_update)
		helpers.ErrorCheck(err_update)
		rec_comp, err_comp := rows_update.ExecContext(ctx,
			limitline_4d_db, limitline_3d_db, limitline_2d_db, limitline_2dd_db, limitline_2dt_db, bbfs_db,
			idcomppasaran, company)
		helpers.ErrorCheck(err_comp)
		update_comp, err_comp := rec_comp.RowsAffected()

		helpers.ErrorCheck(err_comp)
		defer rows_update.Close()
		if update_comp > 0 {
			flag = true
			msg = "Success"
			log.Printf("Update %s Success : %d\n", config.DB_tbl_mst_company_game_pasaran, idcomppasaran)
		} else {
			flag = true
			msg = "Failed"
			log.Printf("Update %s Failed \n", config.DB_tbl_mst_company_game_pasaran)
		}
	}

	if flag {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	}

	return res, nil
}
func Fetch_companyPasaran432(master, company, pasarancode string, idcomppasaran int) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	msg := "Failed"
	flag := false

	sql_pasaran := `SELECT 
			1_minbet as minbet_432d, 1_maxbet4d as maxbet4d_432d, 1_maxbet3d as maxbet3d_432d,
			1_maxbet2d as maxbet2d_432d, 1_maxbet2dd as maxbet2dd_432d, 1_maxbet2dt as maxbet2dt_432d, 
			1_limittotal4d as limitotal4d_432d, 1_limittotal3d as limitotal3d_432d, 1_limittotal2d as limitotal2d_432d, 
			1_limittotal2dd as limitotal2dd_432d, 1_limittotal2dt as limitotal2dt_432d, 
			1_limitbuang4d as limitglobal4d_432d, 1_limitbuang3d as limitglobal3d_432d, 1_limitbuang2d as limitglobal2d_432d, 
			1_limitbuang2dd as limitglobal2dd_432d, 1_limitbuang2dt as limitglobal2dt_432d, 
			1_disc4d as disc4d_432d, 1_disc3d as disc3d_432d, 1_disc2d as disc2d_432d, 1_disc2dd as disc2dd_432d, 1_disc2dt as disc2dt_432d, 
			1_win4d as win4d_432d, 1_win3d as win3d_432d, 1_win2d as win2d_432d, 1_win2dd as win2dd_432d, 1_win2dt as win2dt_432d 
			FROM ` + config.DB_tbl_mst_pasaran_togel + ` 
			WHERE idpasarantogel = ? 
		`
	var (
		minbet_432d_db, maxbet4d_432d_db, maxbet3d_432d_db, maxbet2d_432d_db, maxbet2dd_432d_db, maxbet2dt_432d_db          float32
		limitotal4d_432d_db, limitotal3d_432d_db, limitotal2d_432d_db, limitotal2dd_432d_db, limitotal2dt_432d_db           float32
		limitglobal4d_432d_db, limitglobal3d_432d_db, limitglobal2d_432d_db, limitglobal2dd_432d_db, limitglobal2dt_432d_db float32
		disc4d_432d_db, disc3d_432d_db, disc2d_432d_db, disc2dd_432d_db, disc2dt_432d_db                                    float32
		win4d_432d_db, win3d_432d_db, win2d_432d_db, win2dd_432d_db, win2dt_432d_db                                         float32
	)
	rows_select := con.QueryRowContext(ctx, sql_pasaran, pasarancode)
	switch err := rows_select.Scan(
		&minbet_432d_db, &maxbet4d_432d_db, &maxbet3d_432d_db, &maxbet2d_432d_db, &maxbet2dd_432d_db, &maxbet2dt_432d_db,
		&limitotal4d_432d_db, &limitotal3d_432d_db, &limitotal2d_432d_db, &limitotal2dd_432d_db, &limitotal2dt_432d_db,
		&limitglobal4d_432d_db, &limitglobal3d_432d_db, &limitglobal2d_432d_db, &limitglobal2dd_432d_db, &limitglobal2dt_432d_db,
		&disc4d_432d_db, &disc3d_432d_db, &disc2d_432d_db, &disc2dd_432d_db, &disc2dt_432d_db,
		&win4d_432d_db, &win3d_432d_db, &win2d_432d_db, &win2dd_432d_db, &win2dt_432d_db); err {
	case sql.ErrNoRows:
		flag = false
	case nil:
		flag = true
	default:
		flag = false
		helpers.ErrorCheck(err)
	}

	if flag {
		sql_update := `
			UPDATE   
			` + config.DB_tbl_mst_company_game_pasaran + ` 
			SET 1_minbet=?, 1_maxbet4d=?, 1_maxbet3d=?, 1_maxbet2d=?, 1_maxbet2dd=?, 1_maxbet2dt=?, 
			1_win4d=? , 1_win3d=? , 1_win2d=? , 1_win2dd=? , 1_win2dt=?, 
			1_disc4d=?, 1_disc3d=?, 1_disc2d=?, 1_disc2dd=?, 1_disc2dt=?, 
			1_limitbuang4d=?,1_limitbuang3d=?,1_limitbuang2d=?,1_limitbuang2dd=?,1_limitbuang2dt=?, 
			1_limittotal4d=?,1_limittotal3d=?,1_limittotal2d=?,1_limittotal2dd=?,1_limittotal2dt=?  
			WHERE idcomppasaran=? AND idcompany=? 
		`
		rows_update, err_update := con.PrepareContext(ctx, sql_update)
		helpers.ErrorCheck(err_update)
		log.Println(minbet_432d_db)
		rec_comp, err_comp := rows_update.ExecContext(ctx,
			minbet_432d_db, maxbet4d_432d_db, maxbet3d_432d_db, maxbet2d_432d_db, maxbet2dd_432d_db, maxbet2dt_432d_db,
			win4d_432d_db, win3d_432d_db, win2d_432d_db, win2dd_432d_db, win2dt_432d_db,
			disc4d_432d_db, disc3d_432d_db, disc2d_432d_db, disc2dd_432d_db, disc2dt_432d_db,
			limitglobal4d_432d_db, limitglobal3d_432d_db, limitglobal2d_432d_db, limitglobal2dd_432d_db, limitglobal2dt_432d_db,
			limitotal4d_432d_db, limitotal3d_432d_db, limitotal2d_432d_db, limitotal2dd_432d_db, limitotal2dt_432d_db,
			idcomppasaran, company)
		helpers.ErrorCheck(err_comp)
		update_comp, err_comp := rec_comp.RowsAffected()

		helpers.ErrorCheck(err_comp)
		defer rows_update.Close()
		if update_comp > 0 {
			flag = true
			msg = "Success"
			log.Printf("Update %s Success : %d\n", config.DB_tbl_mst_company_game_pasaran, idcomppasaran)
		} else {
			flag = true
			msg = "Failed"
			log.Printf("Update %s Failed \n", config.DB_tbl_mst_company_game_pasaran)
		}
	}

	if flag {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	}

	return res, nil
}
func Fetch_companyPasarancolokbebas(master, company, pasarancode string, idcomppasaran int) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	msg := "Failed"
	flag := false

	sql_pasaran := `SELECT 
			2_minbet as minbet_cbebas, 2_maxbet as maxbet_cbebas, 
			2_win as win_cbebas, 2_disc as disc_cbebas, 
			2_limitbuang as limitglobal_cbebas, 2_limitotal as limittotal_cbebas 
			FROM ` + config.DB_tbl_mst_pasaran_togel + ` 
			WHERE idpasarantogel = ? 
		`
	var (
		minbet_cbebas_db, maxbet_cbebas_db, win_cbebas_db, disc_cbebas_db, limitglobal_cbebas_db, limittotal_cbebas_db float32
	)
	rows_select := con.QueryRowContext(ctx, sql_pasaran, pasarancode)
	switch err := rows_select.Scan(
		&minbet_cbebas_db, &maxbet_cbebas_db, &win_cbebas_db, &disc_cbebas_db, &limitglobal_cbebas_db, &limittotal_cbebas_db); err {
	case sql.ErrNoRows:
		flag = false
	case nil:
		flag = true
	default:
		flag = false
		helpers.ErrorCheck(err)
	}
	if flag {
		sql_update := `
			UPDATE   
			` + config.DB_tbl_mst_company_game_pasaran + ` 
			SET 2_minbet=? , 2_maxbet=?, 2_win=?, 2_disc=?, 
			2_limitbuang=?, 2_limitotal=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`
		rows_update, err_update := con.PrepareContext(ctx, sql_update)
		helpers.ErrorCheck(err_update)

		rec_comp, err_comp := rows_update.ExecContext(ctx,
			minbet_cbebas_db, maxbet_cbebas_db, win_cbebas_db, disc_cbebas_db,
			limitglobal_cbebas_db, limittotal_cbebas_db, idcomppasaran, company)
		helpers.ErrorCheck(err_comp)
		update_comp, err_comp := rec_comp.RowsAffected()
		helpers.ErrorCheck(err_comp)
		defer rows_update.Close()
		if update_comp > 0 {
			flag = true
			msg = "Success"
			log.Printf("Update %s Success : %d\n", config.DB_tbl_mst_company_game_pasaran, idcomppasaran)
		} else {
			flag = false
			msg = "Failed"
			log.Printf("Update %s Failed \n", config.DB_tbl_mst_company_game_pasaran)
		}
	}

	if flag {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	}

	return res, nil
}
func Fetch_companyPasarancolokmacau(master, company, pasarancode string, idcomppasaran int) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	msg := "Failed"
	flag := false

	sql_pasaran := `SELECT 
			3_minbet as minbet_cmacau, 3_maxbet as maxbet_cmacau, 
			3_win2digit as win2d_cmacau, 3_win3digit as win3d_cmacau, 3_win4digit as win4d_cmacau, 
			3_disc as disc_cmacau, 3_limitbuang as limitglobal_cmacau, 3_limittotal as limitotal_cmacau 
			FROM ` + config.DB_tbl_mst_pasaran_togel + ` 
			WHERE idpasarantogel = ? 
		`
	var (
		minbet_cmacau_db, maxbet_cmacau_db, win2d_cmacau_db, win3d_cmacau_db, win4d_cmacau_db, disc_cmacau_db, limitglobal_cmacau_db, limitotal_cmacau_db float32
	)
	rows_select := con.QueryRowContext(ctx, sql_pasaran, pasarancode)
	switch err := rows_select.Scan(
		&minbet_cmacau_db, &maxbet_cmacau_db, &win2d_cmacau_db, &win3d_cmacau_db, &win4d_cmacau_db, &disc_cmacau_db, &limitglobal_cmacau_db, &limitotal_cmacau_db); err {
	case sql.ErrNoRows:
		flag = false
	case nil:
		flag = true
	default:
		flag = false
		helpers.ErrorCheck(err)
	}

	if flag {
		sql_update := `
			UPDATE   
			` + config.DB_tbl_mst_company_game_pasaran + ` 
			SET 3_minbet=? , 3_maxbet=?, 3_win2digit=?, 3_win3digit=?, 3_win4digit=?, 
			3_disc=?, 3_limitbuang=?, 3_limittotal=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`

		rows_update, err_update := con.PrepareContext(ctx, sql_update)
		helpers.ErrorCheck(err_update)

		rec_comp, err_comp := rows_update.ExecContext(ctx,
			minbet_cmacau_db, maxbet_cmacau_db,
			win2d_cmacau_db, win3d_cmacau_db, win4d_cmacau_db, disc_cmacau_db, limitglobal_cmacau_db, limitotal_cmacau_db, idcomppasaran, company)
		helpers.ErrorCheck(err_comp)
		update_comp, err_comp := rec_comp.RowsAffected()
		helpers.ErrorCheck(err_comp)
		defer rows_update.Close()
		if update_comp > 0 {
			flag = true
			msg = "Success"
			log.Printf("Update %s Success : %d\n", config.DB_tbl_mst_company_game_pasaran, idcomppasaran)
		} else {
			flag = false
			msg = "Failed"
			log.Printf("Update %s Failed \n", config.DB_tbl_mst_company_game_pasaran)
		}

	}

	if flag {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	}

	return res, nil
}
func Fetch_companyPasarancoloknaga(master, company, pasarancode string, idcomppasaran int) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	msg := "Failed"
	flag := false

	sql_pasaran := `SELECT 
			4_minbet as minbet_cnaga, 4_maxbet as maxbet_cnaga, 
			4_win3digit as win3_cnaga, 4_win4digit as win4_cnaga, 
			4_disc as disc_cnaga, 4_limitbuang as limitglobal_cnaga, 4_limittotal as limittotal_cnaga
			FROM ` + config.DB_tbl_mst_pasaran_togel + ` 
			WHERE idpasarantogel = ? 
		`
	var (
		minbet_cnaga_db, maxbet_cnaga_db, win3_cnaga_db, win4_cnaga_db, disc_cnaga_db, limitglobal_cnaga_db, limittotal_cnaga_db float32
	)
	rows_select := con.QueryRowContext(ctx, sql_pasaran, pasarancode)
	switch err := rows_select.Scan(
		&minbet_cnaga_db, &maxbet_cnaga_db, &win3_cnaga_db, &win4_cnaga_db, &disc_cnaga_db, &limitglobal_cnaga_db, &limittotal_cnaga_db); err {
	case sql.ErrNoRows:
		flag = false
	case nil:
		flag = true
	default:
		flag = false
		helpers.ErrorCheck(err)
	}
	if flag {
		sql_update := `
			UPDATE   
			` + config.DB_tbl_mst_company_game_pasaran + ` 
			SET 4_minbet=? , 4_maxbet=?, 4_win3digit=?, 4_win4digit=?,  
			4_disc=?, 4_limitbuang=?, 4_limittotal=? 
			WHERE idcomppasaran=? AND idcompany=?  
		`

		rows_update, err_update := con.PrepareContext(ctx, sql_update)
		helpers.ErrorCheck(err_update)

		rec_comp, err_comp := rows_update.ExecContext(ctx,
			minbet_cnaga_db, maxbet_cnaga_db,
			win3_cnaga_db, win4_cnaga_db, disc_cnaga_db, limitglobal_cnaga_db, limittotal_cnaga_db, idcomppasaran, company)
		helpers.ErrorCheck(err_comp)
		update_comp, err_comp := rec_comp.RowsAffected()
		helpers.ErrorCheck(err_comp)
		defer rows_update.Close()
		if update_comp > 0 {
			flag = true
			msg = "Success"
			log.Printf("Update %s Success : %d\n", config.DB_tbl_mst_company_game_pasaran, idcomppasaran)
		} else {
			flag = false
			msg = "Failed"
			log.Printf("Update %s Failed \n", config.DB_tbl_mst_company_game_pasaran)
		}
	}

	if flag {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	}

	return res, nil
}
func Fetch_companyPasarancolokjitu(master, company, pasarancode string, idcomppasaran int) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	msg := "Failed"
	flag := false

	sql_pasaran := `SELECT 
			5_minbet as minbet_cjitu, 5_maxbet as maxbet_cjitu, 
			5_winas as winas_cjitu, 5_winkop as winkop_cjitu, 5_winkepala as winkepala_cjitu, 5_winekor as winekor_cjitu, 
			5_desic as desc_cjitu, 5_limitbuang as limitglobal_cjitu, 5_limitotal as limittotal_cjitu 
			FROM ` + config.DB_tbl_mst_pasaran_togel + ` 
			WHERE idpasarantogel = ? 
		`
	var (
		minbet_cjitu_db, maxbet_cjitu_db, winas_cjitu_db, winkop_cjitu_db, winkepala_cjitu_db, winekor_cjitu_db, desc_cjitu_db, limitglobal_cjitu_db, limittotal_cjitu_db float32
	)
	rows_select := con.QueryRowContext(ctx, sql_pasaran, pasarancode)
	switch err := rows_select.Scan(
		&minbet_cjitu_db, &maxbet_cjitu_db,
		&winas_cjitu_db, &winkop_cjitu_db, &winkepala_cjitu_db, &winekor_cjitu_db,
		&desc_cjitu_db, &limitglobal_cjitu_db, &limittotal_cjitu_db); err {
	case sql.ErrNoRows:
		flag = false
	case nil:
		flag = true
	default:
		flag = false
		helpers.ErrorCheck(err)
	}
	if flag {
		sql_update := `
			UPDATE   
			` + config.DB_tbl_mst_company_game_pasaran + ` 
			SET 5_minbet=? , 5_maxbet=?, 
			5_winas=?, 5_winkop=?, 5_winkepala=?, 5_winekor=?, 
			5_desic=?, 5_limitbuang=?, 5_limitotal=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`
		rows_update, err_update := con.PrepareContext(ctx, sql_update)
		helpers.ErrorCheck(err_update)

		rec_comp, err_comp := rows_update.ExecContext(ctx,
			minbet_cjitu_db, maxbet_cjitu_db,
			winas_cjitu_db, winkop_cjitu_db, winkepala_cjitu_db, winekor_cjitu_db,
			desc_cjitu_db, limitglobal_cjitu_db, limittotal_cjitu_db, idcomppasaran, company)
		helpers.ErrorCheck(err_comp)
		update_comp, err_comp := rec_comp.RowsAffected()
		helpers.ErrorCheck(err_comp)
		defer rows_update.Close()
		if update_comp > 0 {
			flag = true
			msg = "Success"
			log.Printf("Update %s Success : %d\n", config.DB_tbl_mst_company_game_pasaran, idcomppasaran)
		} else {
			flag = false
			msg = "Failed"
			log.Printf("Update %s Failed \n", config.DB_tbl_mst_company_game_pasaran)
		}
	}

	if flag {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	}

	return res, nil
}
func Fetch_companyPasaran5050umum(master, company, pasarancode string, idcomppasaran int) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	msg := "Failed"
	flag := false

	sql_pasaran := `SELECT 
			6_minbet as minbet_5050umum, 6_maxbet as maxbet_5050umum, 
			6_keibesar as keibesar_5050umum, 6_keikecil as keikecil_5050umum, 6_keigenap as keigenap_5050umum, 
			6_keiganjil as keiganjil_5050umum, 6_keitengah as keitengah_5050umum, 6_keitepi as keitepi_5050umum, 
			6_discbesar as discbesar_5050umum, 6_disckecil as disckecil_5050umum, 6_discgenap as discgenap_5050umum, 
			6_discganjil as discganjil_5050umum, 6_disctengah as disctengah_5050umum, 6_disctepi as disctepi_5050umum, 
			6_limitbuang as limitglobal_5050umum, 6_limittotal as limittotal_5050umum 
			FROM ` + config.DB_tbl_mst_pasaran_togel + ` 
			WHERE idpasarantogel = ? 
		`
	var (
		minbet_5050umum_db, maxbet_5050umum_db, keibesar_5050umum_db, keikecil_5050umum_db, keigenap_5050umum_db, keiganjil_5050umum_db, keitengah_5050umum_db, keitepi_5050umum_db                float32
		discbesar_5050umum_db, disckecil_5050umum_db, discgenap_5050umum_db, discganjil_5050umum_db, disctengah_5050umum_db, disctepi_5050umum_db, limitglobal_5050umum_db, limittotal_5050umum_db float64
	)
	rows_select := con.QueryRowContext(ctx, sql_pasaran, pasarancode)
	switch err := rows_select.Scan(
		&minbet_5050umum_db, &maxbet_5050umum_db,
		&keibesar_5050umum_db, &keikecil_5050umum_db, &keigenap_5050umum_db, &keiganjil_5050umum_db,
		&keitengah_5050umum_db, &keitepi_5050umum_db,
		&discbesar_5050umum_db, &disckecil_5050umum_db, &discgenap_5050umum_db, &discganjil_5050umum_db,
		&disctengah_5050umum_db, &disctepi_5050umum_db,
		&limitglobal_5050umum_db, &limittotal_5050umum_db); err {
	case sql.ErrNoRows:
		flag = false
	case nil:
		flag = true
	default:
		flag = false
		helpers.ErrorCheck(err)
	}
	if flag {
		sql_update := `
			UPDATE   
			` + config.DB_tbl_mst_company_game_pasaran + ` 
			SET 6_minbet=? , 6_maxbet=?, 
			6_keibesar=?, 6_keikecil=?, 6_keigenap=?, 6_keiganjil=?, 6_keitengah=?, 6_keitepi=?, 
			6_discbesar=?, 6_disckecil=?, 6_discgenap=?, 6_discganjil=?, 6_disctengah=?, 6_disctepi=?,  
			6_limitbuang=?, 6_limittotal=?   
			WHERE idcomppasaran=? AND idcompany=? 
		`
		rows_update, err_update := con.PrepareContext(ctx, sql_update)
		helpers.ErrorCheck(err_update)

		rec_comp, err_comp := rows_update.ExecContext(ctx,
			minbet_5050umum_db, maxbet_5050umum_db,
			keibesar_5050umum_db, keikecil_5050umum_db, keigenap_5050umum_db, keiganjil_5050umum_db,
			keitengah_5050umum_db, keitepi_5050umum_db,
			discbesar_5050umum_db, disckecil_5050umum_db, discgenap_5050umum_db, discganjil_5050umum_db,
			disctengah_5050umum_db, disctepi_5050umum_db,
			limitglobal_5050umum_db, limittotal_5050umum_db, idcomppasaran, company)
		helpers.ErrorCheck(err_comp)
		update_comp, err_comp := rec_comp.RowsAffected()
		helpers.ErrorCheck(err_comp)
		defer rows_update.Close()
		if update_comp > 0 {
			flag = true
			msg = "Success"
			log.Printf("Update %s Success : %d\n", config.DB_tbl_mst_company_game_pasaran, idcomppasaran)
		} else {
			flag = false
			msg = "Success"
			log.Printf("Update %s Failed \n", config.DB_tbl_mst_company_game_pasaran)
		}
	}

	if flag {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	}

	return res, nil
}
func Fetch_companyPasaran5050special(master, company, pasarancode string, idcomppasaran int) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	msg := "Failed"
	flag := false

	sql_pasaran := `SELECT 
			7_minbet as minbet_5050special, 7_maxbet as maxbet_5050special,
			7_keiasganjil as keiasganjil_5050special, 7_keiasgenap as keiasgenap_5050special, 7_keiasbesar as keiasbesar_5050special, 
			7_keiaskecil as keiaskecil_5050special, 7_keikopganjil as keikopganjil_5050special, 7_keikopgenap as keikopgenap_5050special, 
			7_keikopbesar as keikopbesar_5050special, 7_keikopkecil as keikopkecil_5050special, 7_keikepalaganjil as keikepalaganjil_5050special, 
			7_keikepalagenap as keikepalagenap_5050special, 7_keikepalabesar as keikepalabesar_5050special, 7_keikepalakecil as keikepalakecil_5050special, 
			7_keiekorganjil as keiekorganjil_5050special, 7_keiekorgenap as keiekorgenap_5050special, 7_keiekorbesar as keiekorbesar_5050special, 
			7_keiekorkecil as keiekorkecil_5050special, 
			7_discasganjil as discasganjil_5050special, 7_discasgenap as discasgenap_5050special, 7_discasbesar as discasbesar_5050special, 
			7_discaskecil as discaskecil_5050special, 7_disckopganjil as disckopganjil_5050special, 7_disckopgenap as disckopgenap_5050special, 
			7_disckopbesar as disckopbesar_5050special, 7_disckopkecil as disckopkecil_5050special, 7_disckepalaganjil as disckepalaganjil_5050special, 
			7_disckepalagenap as disckepalagenap_5050special, 7_disckepalabesar as disckepalabesar_5050special, 7_disckepalakecil as disckepalakecil_5050special, 
			7_discekorganjil as discekorganjil_5050special, 7_discekorgenap as discekorgenap_5050special, 7_discekorbesar as discekorbesar_5050special, 
			7_discekorkecil as discekorkecil_5050special, 7_limitbuang as limitglobal_5050special, 7_limittotal as limittotal_5050special 
			FROM ` + config.DB_tbl_mst_pasaran_togel + ` 
			WHERE idpasarantogel = ? 
		`
	var (
		minbet_5050special_db, maxbet_5050special_db, keiasganjil_5050special_db, keiasgenap_5050special_db, keiasbesar_5050special_db, keiaskecil_5050special_db, keikopganjil_5050special_db, keikopgenap_5050special_db                                                                                       float32
		keikopbesar_5050special_db, keikopkecil_5050special_db, keikepalaganjil_5050special_db, keikepalagenap_5050special_db, keikepalabesar_5050special_db, keikepalakecil_5050special_db, keiekorganjil_5050special_db, keiekorgenap_5050special_db, keiekorbesar_5050special_db, keiekorkecil_5050special_db float32
		discasganjil_5050special_db, discasgenap_5050special_db, discasbesar_5050special_db, discaskecil_5050special_db, disckopganjil_5050special_db, disckopgenap_5050special_db, disckopbesar_5050special_db, disckopkecil_5050special_db, disckepalaganjil_5050special_db, disckepalagenap_5050special_db    float32
		disckepalabesar_5050special_db, disckepalakecil_5050special_db, discekorganjil_5050special_db, discekorgenap_5050special_db, discekorbesar_5050special_db, discekorkecil_5050special_db                                                                                                                  float32
		limitglobal_5050special_db, limittotal_5050special_db                                                                                                                                                                                                                                                    float32
	)
	rows_select := con.QueryRowContext(ctx, sql_pasaran, pasarancode)
	switch err := rows_select.Scan(
		&minbet_5050special_db, &maxbet_5050special_db,
		&keiasganjil_5050special_db, &keiasgenap_5050special_db, &keiasbesar_5050special_db, &keiaskecil_5050special_db,
		&keikopganjil_5050special_db, &keikopgenap_5050special_db, &keikopbesar_5050special_db, &keikopkecil_5050special_db,
		&keikepalaganjil_5050special_db, &keikepalagenap_5050special_db, &keikepalabesar_5050special_db, &keikepalakecil_5050special_db,
		&keiekorganjil_5050special_db, &keiekorgenap_5050special_db, &keiekorbesar_5050special_db, &keiekorkecil_5050special_db,
		&discasganjil_5050special_db, &discasgenap_5050special_db, &discasbesar_5050special_db, &discaskecil_5050special_db,
		&disckopganjil_5050special_db, &disckopgenap_5050special_db, &disckopbesar_5050special_db, &disckopkecil_5050special_db,
		&disckepalaganjil_5050special_db, &disckepalagenap_5050special_db, &disckepalabesar_5050special_db, &disckepalakecil_5050special_db,
		&discekorganjil_5050special_db, &discekorgenap_5050special_db, &discekorbesar_5050special_db, &discekorkecil_5050special_db,
		&limitglobal_5050special_db, &limittotal_5050special_db); err {
	case sql.ErrNoRows:
		flag = false
	case nil:
		flag = true

	default:
		flag = false
		helpers.ErrorCheck(err)
	}
	if flag {
		sql_update := `
			UPDATE   
			` + config.DB_tbl_mst_company_game_pasaran + ` 
			SET 7_minbet=? , 7_maxbet=?, 
			7_keiasganjil=?, 7_keiasgenap=?, 7_keiasbesar=?, 7_keiaskecil=?, 
			7_keikopganjil=?, 7_keikopgenap=?, 7_keikopbesar=?, 7_keikopkecil=?, 
			7_keikepalaganjil=?, 7_keikepalagenap=?, 7_keikepalabesar=?, 7_keikepalakecil=?,  
			7_keiekorganjil=?, 7_keiekorgenap=?, 7_keiekorbesar=?, 7_keiekorkecil=?, 
			7_discasganjil=?, 7_discasgenap=?, 7_discasbesar=?, 7_discaskecil=?, 
			7_disckopganjil=?, 7_disckopgenap=?, 7_disckopbesar=?, 7_disckopkecil=?, 
			7_disckepalaganjil=?, 7_disckepalagenap=?, 7_disckepalabesar=?, 7_disckepalakecil=?, 
			7_discekorganjil=?, 7_discekorgenap=?, 7_discekorbesar=?, 7_discekorkecil=?, 
			7_limitbuang=?, 7_limittotal=?  
			WHERE idcomppasaran=? AND idcompany=? 
		`
		log.Println(sql_update)
		rows_update, err_update := con.PrepareContext(ctx, sql_update)

		helpers.ErrorCheck(err_update)

		rec_comp, err_comp := rows_update.ExecContext(ctx,
			minbet_5050special_db, maxbet_5050special_db,
			keiasganjil_5050special_db, keiasgenap_5050special_db, keiasbesar_5050special_db, keiaskecil_5050special_db,
			keikopganjil_5050special_db, keikopgenap_5050special_db, keikopbesar_5050special_db, keikopkecil_5050special_db,
			keikepalaganjil_5050special_db, keikepalagenap_5050special_db, keikepalabesar_5050special_db, keikepalakecil_5050special_db,
			keiekorganjil_5050special_db, keiekorgenap_5050special_db, keiekorbesar_5050special_db, keiekorkecil_5050special_db,
			discasganjil_5050special_db, discasgenap_5050special_db, discasbesar_5050special_db, discaskecil_5050special_db,
			disckopganjil_5050special_db, disckopgenap_5050special_db, disckopbesar_5050special_db, disckopkecil_5050special_db,
			disckepalaganjil_5050special_db, disckepalagenap_5050special_db, disckepalabesar_5050special_db, disckepalakecil_5050special_db,
			discekorganjil_5050special_db, discekorgenap_5050special_db, discekorbesar_5050special_db, discekorkecil_5050special_db,
			limitglobal_5050special_db, limittotal_5050special_db, idcomppasaran, company)
		helpers.ErrorCheck(err_comp)
		update_comp, err_comp := rec_comp.RowsAffected()
		helpers.ErrorCheck(err_comp)
		defer rows_update.Close()
		log.Println(update_comp)
		if update_comp > 0 {
			flag = true
			msg = "Success"
			log.Printf("Update %s Success : %d\n", config.DB_tbl_mst_company_game_pasaran, idcomppasaran)
		} else {
			flag = false
			msg = "Failed"
			log.Printf("Update %s Failed \n", config.DB_tbl_mst_company_game_pasaran)
		}
	}

	if flag {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	}

	return res, nil
}
func Fetch_companyPasaran5050kombinasi(master, company, pasarancode string, idcomppasaran int) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	msg := "Failed"
	flag := false

	sql_pasaran := `SELECT 
			8_minbet as minbet_5050kombinasi, 8_maxbet as maxbet_5050kombinasi, 
			8_belakangkeimono as belakangkeimono_5050kombinasi, 8_belakangkeistereo as belakangkeistereo_5050kombinasi, 8_belakangkeikembang as belakangkeikembang_5050kombinasi, 8_belakangkeikempis as belakangkeikempis_5050kombinasi, 8_belakangkeikembar as belakangkeikembar_5050kombinasi, 
			8_tengahkeimono as tengahkeimono_5050kombinasi, 8_tengahkeistereo as tengahkeistereo_5050kombinasi, 8_tengahkeikembang as tengahkeikembang_5050kombinasi, 8_tengahkeikempis as tengahkeikempis_5050kombinasi, 8_tengahkeikembar as tengahkeikembar_5050kombinasi, 
			8_depankeimono as depankeimono_5050kombinasi, 8_depankeistereo as depankeistereo_5050kombinasi, 8_depankeikembang as depankeikembang_5050kombinasi, 8_depankeikempis as depankeikempis_5050kombinasi, 8_depankeikembar as depankeikembar_5050kombinasi, 
			8_belakangdiscmono as belakangdiscmono_5050kombinasi, 8_belakangdiscstereo as belakangdiscstereo_5050kombinasi, 8_belakangdisckembang as belakangdisckembang_5050kombinasi, 8_belakangdisckempis as belakangdisckempis_5050kombinasi, 8_belakangdisckembar as belakangdisckembar_5050kombinasi, 
			8_tengahdiscmono as tengahdiscmono_5050kombinasi, 8_tengahdiscstereo as tengahdiscstereo_5050kombinasi, 8_tengahdisckembang as tengahdisckembang_5050kombinasi, 8_tengahdisckempis as tengahdisckempis_5050kombinasi, 8_tengahdisckembar as tengahdisckembar_5050kombinasi, 
			8_depandiscmono as depandiscmono_5050kombinasi, 8_depandiscstereo as depandiscstereo_5050kombinasi, 8_depandisckembang as depandisckembang_5050kombinasi, 8_depandisckempis as depandisckempis_5050kombinasi, 8_depandisckembar as depandisckembar_5050kombinasi, 
			8_limitbuang as limitglobal_5050kombinasi, 8_limittotal as limittotal_5050kombinasi 
			FROM ` + config.DB_tbl_mst_pasaran_togel + ` 
			WHERE idpasarantogel = ? 
		`
	var (
		minbet_5050kombinasi_db, maxbet_5050kombinasi_db                                                                                                                                       float32
		belakangkeimono_5050kombinasi_db, belakangkeistereo_5050kombinasi_db, belakangkeikembang_5050kombinasi_db, belakangkeikempis_5050kombinasi_db, belakangkeikembar_5050kombinasi_db      float32
		tengahkeimono_5050kombinasi_db, tengahkeistereo_5050kombinasi_db, tengahkeikembang_5050kombinasi_db, tengahkeikempis_5050kombinasi_db, tengahkeikembar_5050kombinasi_db                float32
		depankeimono_5050kombinasi_db, depankeistereo_5050kombinasi_db, depankeikembang_5050kombinasi_db, depankeikempis_5050kombinasi_db, depankeikembar_5050kombinasi_db                     float32
		belakangdiscmono_5050kombinasi_db, belakangdiscstereo_5050kombinasi_db, belakangdisckembang_5050kombinasi_db, belakangdisckempis_5050kombinasi_db, belakangdisckembar_5050kombinasi_db float32
		tengahdiscmono_5050kombinasi_db, tengahdiscstereo_5050kombinasi_db, tengahdisckembang_5050kombinasi_db, tengahdisckempis_5050kombinasi_db, tengahdisckembar_5050kombinasi_db           float32
		depandiscmono_5050kombinasi_db, depandiscstereo_5050kombinasi_db, depandisckembang_5050kombinasi_db, depandisckempis_5050kombinasi_db, depandisckembar_5050kombinasi_db                float32
		limitglobal_5050kombinasi_db, limittotal_5050kombinasi_db                                                                                                                              float32
	)
	rows_select := con.QueryRowContext(ctx, sql_pasaran, pasarancode)
	switch err := rows_select.Scan(
		&minbet_5050kombinasi_db, &maxbet_5050kombinasi_db,
		&belakangkeimono_5050kombinasi_db, &belakangkeistereo_5050kombinasi_db, &belakangkeikembang_5050kombinasi_db, &belakangkeikempis_5050kombinasi_db, &belakangkeikembar_5050kombinasi_db,
		&tengahkeimono_5050kombinasi_db, &tengahkeistereo_5050kombinasi_db, &tengahkeikembang_5050kombinasi_db, &tengahkeikempis_5050kombinasi_db, &tengahkeikembar_5050kombinasi_db,
		&depankeimono_5050kombinasi_db, &depankeistereo_5050kombinasi_db, &depankeikembang_5050kombinasi_db, &depankeikempis_5050kombinasi_db, &depankeikembar_5050kombinasi_db,
		&belakangdiscmono_5050kombinasi_db, &belakangdiscstereo_5050kombinasi_db, &belakangdisckembang_5050kombinasi_db, &belakangdisckempis_5050kombinasi_db, &belakangdisckembar_5050kombinasi_db,
		&tengahdiscmono_5050kombinasi_db, &tengahdiscstereo_5050kombinasi_db, &tengahdisckembang_5050kombinasi_db, &tengahdisckempis_5050kombinasi_db, &tengahdisckembar_5050kombinasi_db,
		&depandiscmono_5050kombinasi_db, &depandiscstereo_5050kombinasi_db, &depandisckembang_5050kombinasi_db, &depandisckempis_5050kombinasi_db, &depandisckembar_5050kombinasi_db,
		&limitglobal_5050kombinasi_db, &limittotal_5050kombinasi_db); err {
	case sql.ErrNoRows:
		flag = false
	case nil:
		flag = true

	default:
		flag = false
		helpers.ErrorCheck(err)
	}
	if flag {
		sql_update := `
			UPDATE   
			` + config.DB_tbl_mst_company_game_pasaran + ` 
			SET 8_minbet=? , 8_maxbet=?, 
			8_belakangkeimono=?, 8_belakangkeistereo=?, 8_belakangkeikembang=?, 8_belakangkeikempis=?, 8_belakangkeikembar=?, 
			8_tengahkeimono=?, 8_tengahkeistereo=?, 8_tengahkeikembang=?, 8_tengahkeikempis=?, 8_tengahkeikembar=?, 
			8_depankeimono=?, 8_depankeistereo=?, 8_depankeikembang=?, 8_depankeikempis=?, 8_depankeikembar=?, 
			8_belakangdiscmono=?, 8_belakangdiscstereo=?, 8_belakangdisckembang=?, 8_belakangdisckempis=?, 8_belakangdisckembar=?, 
			8_tengahdiscmono=?, 8_tengahdiscstereo=?, 8_tengahdisckembang=?, 8_tengahdisckempis=?, 8_tengahdisckembar=?, 
			8_depandiscmono=?, 8_depandiscstereo=?, 8_depandisckembang=?, 8_depandisckempis=?, 8_depandisckembar=?, 
			8_limitbuang=?, 8_limittotal=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`
		log.Println(sql_update)
		rows_update, err_update := con.PrepareContext(ctx, sql_update)

		helpers.ErrorCheck(err_update)

		rec_comp, err_comp := rows_update.ExecContext(ctx,
			minbet_5050kombinasi_db, maxbet_5050kombinasi_db,
			belakangkeimono_5050kombinasi_db, belakangkeistereo_5050kombinasi_db, belakangkeikembang_5050kombinasi_db, belakangkeikempis_5050kombinasi_db, belakangkeikembar_5050kombinasi_db,
			tengahkeimono_5050kombinasi_db, tengahkeistereo_5050kombinasi_db, tengahkeikembang_5050kombinasi_db, tengahkeikempis_5050kombinasi_db, tengahkeikembar_5050kombinasi_db,
			depankeimono_5050kombinasi_db, depankeistereo_5050kombinasi_db, depankeikembang_5050kombinasi_db, depankeikempis_5050kombinasi_db, depankeikembar_5050kombinasi_db,
			belakangdiscmono_5050kombinasi_db, belakangdiscstereo_5050kombinasi_db, belakangdisckembang_5050kombinasi_db, belakangdisckempis_5050kombinasi_db, belakangdisckembar_5050kombinasi_db,
			tengahdiscmono_5050kombinasi_db, tengahdiscstereo_5050kombinasi_db, tengahdisckembang_5050kombinasi_db, tengahdisckempis_5050kombinasi_db, tengahdisckembar_5050kombinasi_db,
			depandiscmono_5050kombinasi_db, depandiscstereo_5050kombinasi_db, depandisckembang_5050kombinasi_db, depandisckempis_5050kombinasi_db, depandisckembar_5050kombinasi_db,
			limitglobal_5050kombinasi_db, limittotal_5050kombinasi_db, idcomppasaran, company)
		helpers.ErrorCheck(err_comp)
		update_comp, err_comp := rec_comp.RowsAffected()
		helpers.ErrorCheck(err_comp)
		defer rows_update.Close()
		log.Println(update_comp)
		if update_comp > 0 {
			flag = true
			msg = "Success"
			log.Printf("Update %s Success : %d\n", config.DB_tbl_mst_company_game_pasaran, idcomppasaran)
		} else {
			flag = false
			msg = "Failed"
			log.Printf("Update %s Failed \n", config.DB_tbl_mst_company_game_pasaran)
		}
	}

	if flag {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	}

	return res, nil
}
func Fetch_companyPasaranmacau(master, company, pasarancode string, idcomppasaran int) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	msg := "Failed"
	flag := false

	sql_pasaran := `SELECT 
			9_minbet as minbet_kombinasi, 9_maxbet as maxbet_kombinasi, 9_win as win_kombinasi, 9_discount as disc_kombinasi, 
			9_limitbuang as limitglobal_kombinasi, 9_limittotal as limittotal_kombinasi 
			FROM ` + config.DB_tbl_mst_pasaran_togel + ` 
			WHERE idpasarantogel = ? 
		`
	var (
		minbet_kombinasi_db, maxbet_kombinasi_db, win_kombinasi_db, disc_kombinasi_db, limitglobal_kombinasi_db, limittotal_kombinasi_db float32
	)
	rows_select := con.QueryRowContext(ctx, sql_pasaran, pasarancode)
	switch err := rows_select.Scan(
		&minbet_kombinasi_db, &maxbet_kombinasi_db, &win_kombinasi_db, &disc_kombinasi_db, &limitglobal_kombinasi_db, &limittotal_kombinasi_db); err {
	case sql.ErrNoRows:
		flag = false
	case nil:
		flag = true

	default:
		flag = false
		helpers.ErrorCheck(err)
	}
	if flag {
		sql_update := `
			UPDATE   
			` + config.DB_tbl_mst_company_game_pasaran + ` 
			SET 9_minbet=? , 9_maxbet=?, 9_win=?, 9_discount=?, 
			9_limitbuang=?, 9_limittotal=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`
		log.Println(sql_update)
		rows_update, err_update := con.PrepareContext(ctx, sql_update)

		helpers.ErrorCheck(err_update)

		rec_comp, err_comp := rows_update.ExecContext(ctx,
			minbet_kombinasi_db, maxbet_kombinasi_db, win_kombinasi_db, disc_kombinasi_db, limitglobal_kombinasi_db, limittotal_kombinasi_db, idcomppasaran, company)
		helpers.ErrorCheck(err_comp)
		update_comp, err_comp := rec_comp.RowsAffected()
		helpers.ErrorCheck(err_comp)
		defer rows_update.Close()
		log.Println(update_comp)
		if update_comp > 0 {
			flag = true
			msg = "Success"
			log.Printf("Update %s Success : %d\n", config.DB_tbl_mst_company_game_pasaran, idcomppasaran)
		} else {
			flag = false
			msg = "Failed"
			log.Printf("Update %s Failed \n", config.DB_tbl_mst_company_game_pasaran)
		}
	}

	if flag {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	}

	return res, nil
}
func Fetch_companyPasarandasar(master, company, pasarancode string, idcomppasaran int) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	msg := "Failed"
	flag := false

	sql_pasaran := `SELECT 
			10_minbet as minbet_dasar, 10_maxbet as maxbet_dasar, 
			10_keibesar as keibesar_dasar, 10_keikecil as keikecil_dasar, 10_keigenap as keigenap_dasar, 10_keiganjil as keiganjil_dasar, 
			10_discbesar as discbesar_dasar, 10_disckecil as disckecil_dasar, 10_discigenap as discgenap_dasar, 10_discganjil as discganjil_dasar, 
			10_limitbuang as limitglobal_dasar, 10_limittotal as limittotal_dasar 
			FROM ` + config.DB_tbl_mst_pasaran_togel + ` 
			WHERE idpasarantogel = ? 
		`
	var (
		minbet_dasar_db, maxbet_dasar_db, keibesar_dasar_db, keikecil_dasar_db, keigenap_dasar_db, keiganjil_dasar_db, discbesar_dasar_db, disckecil_dasar_db, discgenap_dasar_db, discganjil_dasar_db, limitglobal_dasar_db, limittotal_dasar_db float32
	)
	rows_select := con.QueryRowContext(ctx, sql_pasaran, pasarancode)
	switch err := rows_select.Scan(
		&minbet_dasar_db, &maxbet_dasar_db, &keibesar_dasar_db, &keikecil_dasar_db, &keigenap_dasar_db, &keiganjil_dasar_db, &discbesar_dasar_db, &disckecil_dasar_db, &discgenap_dasar_db, &discganjil_dasar_db, &limitglobal_dasar_db, &limittotal_dasar_db); err {
	case sql.ErrNoRows:
		flag = false
	case nil:
		flag = true

	default:
		flag = false
		helpers.ErrorCheck(err)
	}
	if flag {
		sql_update := `
			UPDATE   
			` + config.DB_tbl_mst_company_game_pasaran + ` 
			SET 10_minbet=? , 10_maxbet=?, 
			10_keibesar=?, 10_keikecil=?, 10_keigenap=?, 10_keiganjil=?, 
			10_discbesar=?, 10_disckecil=?, 10_discigenap=?, 10_discganjil=?, 
			10_limitbuang=?, 10_limittotal=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`
		log.Println(sql_update)
		rows_update, err_update := con.PrepareContext(ctx, sql_update)

		helpers.ErrorCheck(err_update)

		rec_comp, err_comp := rows_update.ExecContext(ctx,
			minbet_dasar_db, maxbet_dasar_db, keibesar_dasar_db, keikecil_dasar_db, keigenap_dasar_db, keiganjil_dasar_db, discbesar_dasar_db, disckecil_dasar_db, discgenap_dasar_db, discganjil_dasar_db, limitglobal_dasar_db, limittotal_dasar_db, idcomppasaran, company)
		helpers.ErrorCheck(err_comp)
		update_comp, err_comp := rec_comp.RowsAffected()
		helpers.ErrorCheck(err_comp)
		defer rows_update.Close()
		log.Println(update_comp)
		if update_comp > 0 {
			flag = true
			msg = "Success"
			log.Printf("Update %s Success : %d\n", config.DB_tbl_mst_company_game_pasaran, idcomppasaran)
		} else {
			flag = false
			msg = "Failed"
			log.Printf("Update %s Failed \n", config.DB_tbl_mst_company_game_pasaran)
		}
	}

	if flag {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	}

	return res, nil
}
func Fetch_companyPasaranshio(master, company, pasarancode string, idcomppasaran int) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	msg := "Failed"
	flag := false

	sql_pasaran := `SELECT 
			11_minbet as minbet_shio, 11_maxbet as maxbet_shio, 11_win as win_shio, 11_disc as disc_shio, 11_limitbuang as limitglobal_shio, 11_limittotal as limittotal_shio, 
			11_shiotahunini as shioyear_shio 
			FROM ` + config.DB_tbl_mst_pasaran_togel + ` 
			WHERE idpasarantogel = ? 
		`
	var (
		minbet_shio_db, maxbet_shio_db, win_shio_db, disc_shio_db, limitglobal_shio_db, limittotal_shio_db float32
		shioyear_shio_db                                                                                   string
	)
	rows_select := con.QueryRowContext(ctx, sql_pasaran, pasarancode)
	switch err := rows_select.Scan(
		&minbet_shio_db, &maxbet_shio_db, &win_shio_db, &disc_shio_db, &limitglobal_shio_db, &limittotal_shio_db, &shioyear_shio_db); err {
	case sql.ErrNoRows:
		flag = false
	case nil:
		flag = true

	default:
		flag = false
		helpers.ErrorCheck(err)
	}
	if flag {
		sql_update := `
			UPDATE   
			` + config.DB_tbl_mst_company_game_pasaran + ` 
			SET 11_shiotahunini=? , 11_minbet=?, 11_maxbet=?, 
			11_win=?, 11_disc=?, 
			11_limitbuang=?, 11_limittotal=?  
			WHERE idcomppasaran=? AND idcompany=? 
		`
		log.Println(sql_update)
		rows_update, err_update := con.PrepareContext(ctx, sql_update)

		helpers.ErrorCheck(err_update)

		rec_comp, err_comp := rows_update.ExecContext(ctx,
			shioyear_shio_db, minbet_shio_db, maxbet_shio_db, win_shio_db, disc_shio_db, limitglobal_shio_db, limittotal_shio_db, idcomppasaran, company)
		helpers.ErrorCheck(err_comp)
		update_comp, err_comp := rec_comp.RowsAffected()
		helpers.ErrorCheck(err_comp)
		defer rows_update.Close()
		log.Println(update_comp)
		if update_comp > 0 {
			flag = true
			msg = "Success"
			log.Printf("Update %s Success : %d\n", config.DB_tbl_mst_company_game_pasaran, idcomppasaran)
		} else {
			flag = false
			msg = "Failed"
			log.Printf("Update %s Failed \n", config.DB_tbl_mst_company_game_pasaran)
		}
	}

	if flag {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = msg
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	}

	return res, nil
}
func Save_companyUpdatePasaran(
	master, company,
	pasarandiundi, pasaranurl, pasaranjamtutup, pasaranjamjadwal, pasaranjamopen, statuspasaranactive string,
	idcomppasaran int) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()
	msg := "Failed"
	flag := false

	sql_update := `
			UPDATE   
			` + config.DB_tbl_mst_company_game_pasaran + ` 
			SET pasarandiundi=? , pasaranurl=?, jamtutup=?, jamjadwal=?, jamopen=?, statuspasaranactive=?, 
			updatecomppas=?, updatedatecompas=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`

	rows_update, err_update := con.PrepareContext(ctx, sql_update)
	helpers.ErrorCheck(err_update)
	rec_comp, err_comp := rows_update.ExecContext(ctx,
		pasarandiundi,
		pasaranurl, pasaranjamtutup, pasaranjamjadwal, pasaranjamopen, statuspasaranactive,
		master, tglnow.Format("YYYY-MM-DD HH:mm:ss"),
		idcomppasaran, company)
	helpers.ErrorCheck(err_comp)
	update_comp, err_comp := rec_comp.RowsAffected()
	helpers.ErrorCheck(err_comp)
	defer rows_update.Close()
	log.Println(update_comp)
	if update_comp > 0 {
		flag = true
		msg = "Success"
		log.Printf("Update %s Success : %d\n", config.DB_tbl_mst_company_game_pasaran, idcomppasaran)
	} else {
		flag = false
		msg = "Failed"
		log.Printf("Update %s Failed \n", config.DB_tbl_mst_company_game_pasaran)
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
func Save_companyUpdatePasaranLine(
	master, company string, idcomppasaran int,
	limitline_4d, limitline_3d, limitline_2d, limitline_2dd, limitline_2dt, bbfs int) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()
	msg := "Failed"
	flag := false

	sql_update := `
			UPDATE   
			` + config.DB_tbl_mst_company_game_pasaran + ` 
			SET limitline_4d=? , limitline_3d=?, limitline_2d=?, limitline_2dd=?, limitline_2dt=?, bbfs=?, 
			updatecomppas=?, updatedatecompas=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`

	rows_update, err_update := con.PrepareContext(ctx, sql_update)
	helpers.ErrorCheck(err_update)
	rec_comp, err_comp := rows_update.ExecContext(ctx,
		limitline_4d,
		limitline_3d, limitline_2d, limitline_2dd, limitline_2dt, bbfs,
		master, tglnow.Format("YYYY-MM-DD HH:mm:ss"),
		idcomppasaran, company)
	helpers.ErrorCheck(err_comp)
	update_comp, err_comp := rec_comp.RowsAffected()
	helpers.ErrorCheck(err_comp)
	defer rows_update.Close()
	log.Println(update_comp)
	if update_comp > 0 {
		flag = true
		msg = "Success"
		log.Printf("Update %s Success : %d\n", config.DB_tbl_mst_company_game_pasaran, idcomppasaran)
	} else {
		flag = false
		msg = "Failed"
		log.Printf("Update %s Failed \n", config.DB_tbl_mst_company_game_pasaran)
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
func Save_companyUpdatePasaran432(
	master, company string, idcomppasaran int,
	minbet, maxbet4d, maxbet3d, maxbet2d, maxbet2dd, maxbet2dt int,
	win4d, win3d, win2d, win2dd, win2dt int,
	disc4d, disc3d, disc2d, disc2dd, disc2dt float32,
	limitglobal4d, limitglobal3d, limitglobal2d, limitglobal2dd, limitglobal2dt int,
	limittotal4d, limittotal3d, limittotal2d, limittotal2dd, limittotal2dt int) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()
	msg := "Failed"
	flag := false

	sql_update := `
			UPDATE   
			` + config.DB_tbl_mst_company_game_pasaran + ` 
			SET 1_minbet=? , 1_maxbet4d=?, 1_maxbet3d=?, 1_maxbet2d=?, 1_maxbet2dd=?, 1_maxbet2dt=?, 
			1_win4d=?, 1_win3d=?, 1_win2d=?, 1_win2dd=?, 1_win2dt=?, 
			1_disc4d=?, 1_disc3d=?, 1_disc2d=?, 1_disc2dd=?, 1_disc2dt=?, 
			1_limitbuang4d=?, 1_limitbuang3d=?, 1_limitbuang2d=?, 1_limitbuang2dd=?, 1_limitbuang2dt=?,  
			1_limittotal4d=?, 1_limittotal3d=?, 1_limittotal2d=?, 1_limittotal2dd=?, 1_limittotal2dt=?,  
			updatecomppas=?, updatedatecompas=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`

	rows_update, err_update := con.PrepareContext(ctx, sql_update)
	helpers.ErrorCheck(err_update)
	rec_comp, err_comp := rows_update.ExecContext(ctx,
		minbet, maxbet4d, maxbet3d, maxbet2d, maxbet2dd, maxbet2dt,
		win4d, win3d, win2d, win2dd, win2dt,
		disc4d, disc3d, disc2d, disc2dd, disc2dt,
		limitglobal4d, limitglobal3d, limitglobal2d, limitglobal2dd, limitglobal2dt,
		limittotal4d, limittotal3d, limittotal2d, limittotal2dd, limittotal2dt,
		master, tglnow.Format("YYYY-MM-DD HH:mm:ss"),
		idcomppasaran, company)
	helpers.ErrorCheck(err_comp)
	update_comp, err_comp := rec_comp.RowsAffected()
	helpers.ErrorCheck(err_comp)
	defer rows_update.Close()
	log.Println(update_comp)
	if update_comp > 0 {
		flag = true
		msg = "Success"
		log.Printf("Update %s Success : %d\n", config.DB_tbl_mst_company_game_pasaran, idcomppasaran)
	} else {
		flag = false
		msg = "Failed"
		log.Printf("Update %s Failed \n", config.DB_tbl_mst_company_game_pasaran)
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
func Save_companyUpdatePasarancolokbebas(
	master, company string, idcomppasaran int,
	minbet, maxbet int,
	win, disc float32,
	limitglobal, limittotal int) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()
	msg := "Failed"
	flag := false

	sql_update := `
			UPDATE   
			` + config.DB_tbl_mst_company_game_pasaran + ` 
			SET 2_minbet=? , 2_maxbet=?, 2_win=?, 2_disc=?, 
			2_limitbuang=?, 2_limitotal=?,   
			updatecomppas=?, updatedatecompas=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`

	rows_update, err_update := con.PrepareContext(ctx, sql_update)
	helpers.ErrorCheck(err_update)
	rec_comp, err_comp := rows_update.ExecContext(ctx,
		minbet, maxbet, win, disc, limitglobal, limittotal,
		master, tglnow.Format("YYYY-MM-DD HH:mm:ss"),
		idcomppasaran, company)
	helpers.ErrorCheck(err_comp)
	update_comp, err_comp := rec_comp.RowsAffected()
	helpers.ErrorCheck(err_comp)
	defer rows_update.Close()
	log.Println(update_comp)
	if update_comp > 0 {
		flag = true
		msg = "Success"
		log.Printf("Update %s Success : %d\n", config.DB_tbl_mst_company_game_pasaran, idcomppasaran)
	} else {
		flag = false
		msg = "Failed"
		log.Printf("Update %s Failed \n", config.DB_tbl_mst_company_game_pasaran)
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
func Save_companyUpdatePasarancolokmacau(
	master, company string, idcomppasaran int,
	minbet, maxbet int,
	win2, win3, win4, disc float32,
	limitglobal, limittotal int) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()
	msg := "Failed"
	flag := false

	sql_update := `
			UPDATE   
			` + config.DB_tbl_mst_company_game_pasaran + ` 
			SET 3_minbet=? , 3_maxbet=?, 3_win2digit=?, 3_win3digit=?, 3_win4digit=?, 
			3_disc=?, 3_limitbuang=?, 3_limittotal=?,   
			updatecomppas=?, updatedatecompas=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`

	rows_update, err_update := con.PrepareContext(ctx, sql_update)
	helpers.ErrorCheck(err_update)
	rec_comp, err_comp := rows_update.ExecContext(ctx,
		minbet, maxbet,
		win2, win3, win4,
		disc, limitglobal, limittotal,
		master, tglnow.Format("YYYY-MM-DD HH:mm:ss"),
		idcomppasaran, company)
	helpers.ErrorCheck(err_comp)
	update_comp, err_comp := rec_comp.RowsAffected()
	helpers.ErrorCheck(err_comp)
	defer rows_update.Close()
	log.Println(update_comp)
	if update_comp > 0 {
		flag = true
		msg = "Success"
		log.Printf("Update %s Success : %d\n", config.DB_tbl_mst_company_game_pasaran, idcomppasaran)
	} else {
		flag = false
		msg = "Failed"
		log.Printf("Update %s Failed \n", config.DB_tbl_mst_company_game_pasaran)
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
func Save_companyUpdatePasarancoloknaga(
	master, company string, idcomppasaran int,
	minbet, maxbet int,
	win3, win4, disc float32,
	limitglobal, limittotal int) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()
	msg := "Failed"
	flag := false

	sql_update := `
			UPDATE   
			` + config.DB_tbl_mst_company_game_pasaran + ` 
			SET 4_minbet=? , 4_maxbet=?, 4_win3digit=?, 4_win4digit=?,  
			4_disc=?, 4_limitbuang=?, 4_limittotal=?,   
			updatecomppas=?, updatedatecompas=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`

	rows_update, err_update := con.PrepareContext(ctx, sql_update)
	helpers.ErrorCheck(err_update)
	rec_comp, err_comp := rows_update.ExecContext(ctx,
		minbet, maxbet,
		win3, win4,
		disc, limitglobal, limittotal,
		master, tglnow.Format("YYYY-MM-DD HH:mm:ss"),
		idcomppasaran, company)
	helpers.ErrorCheck(err_comp)
	update_comp, err_comp := rec_comp.RowsAffected()
	helpers.ErrorCheck(err_comp)
	defer rows_update.Close()
	log.Println(update_comp)
	if update_comp > 0 {
		flag = true
		msg = "Success"
		log.Printf("Update %s Success : %d\n", config.DB_tbl_mst_company_game_pasaran, idcomppasaran)
	} else {
		flag = false
		msg = "Failed"
		log.Printf("Update %s Failed \n", config.DB_tbl_mst_company_game_pasaran)
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
func Save_companyUpdatePasarancolokjitu(
	master, company string, idcomppasaran int,
	minbet, maxbet int,
	winas, winkop, winkepala, winekor, disc float32,
	limitglobal, limittotal int) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()
	msg := "Failed"
	flag := false

	sql_update := `
			UPDATE   
			` + config.DB_tbl_mst_company_game_pasaran + ` 
			SET 5_minbet=? , 5_maxbet=?, 
			5_winas=?, 5_winkop=?, 5_winkepala=?, 5_winekor=?, 
			5_desic=?, 5_limitbuang=?, 5_limitotal=?,   
			updatecomppas=?, updatedatecompas=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`

	rows_update, err_update := con.PrepareContext(ctx, sql_update)
	helpers.ErrorCheck(err_update)
	rec_comp, err_comp := rows_update.ExecContext(ctx,
		minbet, maxbet,
		winas, winkop, winkepala, winekor,
		disc, limitglobal, limittotal,
		master, tglnow.Format("YYYY-MM-DD HH:mm:ss"),
		idcomppasaran, company)
	helpers.ErrorCheck(err_comp)
	update_comp, err_comp := rec_comp.RowsAffected()
	helpers.ErrorCheck(err_comp)
	defer rows_update.Close()
	log.Println(update_comp)
	if update_comp > 0 {
		flag = true
		msg = "Success"
		log.Printf("Update %s Success : %d\n", config.DB_tbl_mst_company_game_pasaran, idcomppasaran)
	} else {
		flag = false
		msg = "Failed"
		log.Printf("Update %s Failed \n", config.DB_tbl_mst_company_game_pasaran)
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
func Save_companyUpdatePasaran5050umum(
	master, company string, idcomppasaran int,
	minbet, maxbet int,
	keibesar, keikecil, keigenap, keiganjil, keitengah, keitepi float64,
	discbesar, disckecil, discgenap, discganjil, disctengah, disctepi float64,
	limitglobal, limittotal int) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()
	msg := "Failed"
	flag := false

	sql_update := `
			UPDATE   
			` + config.DB_tbl_mst_company_game_pasaran + ` 
			SET 6_minbet=? , 6_maxbet=?, 
			6_keibesar=?, 6_keikecil=?, 6_keigenap=?, 6_keiganjil=?, 6_keitengah=?, 6_keitepi=?, 
			6_discbesar=?, 6_disckecil=?, 6_discgenap=?, 6_discganjil=?, 6_disctengah=?, 6_disctepi=?,  
			6_limitbuang=?, 6_limittotal=?,    
			updatecomppas=?, updatedatecompas=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`

	rows_update, err_update := con.PrepareContext(ctx, sql_update)
	helpers.ErrorCheck(err_update)
	rec_comp, err_comp := rows_update.ExecContext(ctx,
		minbet, maxbet,
		helpers.ToFixed(keibesar, 3), helpers.ToFixed(keikecil, 3),
		helpers.ToFixed(keigenap, 3), helpers.ToFixed(keiganjil, 3),
		helpers.ToFixed(keitengah, 3), helpers.ToFixed(keitepi, 3),
		helpers.ToFixed(discbesar, 3), helpers.ToFixed(disckecil, 3),
		helpers.ToFixed(discgenap, 3), helpers.ToFixed(discganjil, 3),
		helpers.ToFixed(disctengah, 3), helpers.ToFixed(disctepi, 3),
		limitglobal, limittotal,
		master, tglnow.Format("YYYY-MM-DD HH:mm:ss"),
		idcomppasaran, company)
	helpers.ErrorCheck(err_comp)
	update_comp, err_comp := rec_comp.RowsAffected()
	helpers.ErrorCheck(err_comp)
	defer rows_update.Close()
	log.Println(update_comp)
	if update_comp > 0 {
		flag = true
		msg = "Success"
		log.Printf("Update %s Success : %d\n", config.DB_tbl_mst_company_game_pasaran, idcomppasaran)
	} else {
		flag = false
		msg = "Failed"
		log.Printf("Update %s Failed \n", config.DB_tbl_mst_company_game_pasaran)
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
func Save_companyUpdatePasaran5050special(
	master, company string, idcomppasaran int,
	minbet, maxbet int,
	keiasganjil, keiasgenap, keiasbesar, keiaskecil float64,
	keikopganjil, keikopgenap, keikopbesar, keikopkecil float64,
	keikepalaganjil, keikepalagenap, keikepalabesar, keikepalakecil float64,
	keiekorganjil, keiekorgenap, keiekorbesar, keiekorkecil float64,
	discasganjil, discasgenap, discasbesar, discaskecil float64,
	disckopganjil, disckopgenap, disckopbesar, disckopkecil float64,
	disckepalaganjil, disckepalagenap, disckepalabesar, disckepalakecil float64,
	discekorganjil, discekorgenap, discekorbesar, discekorkecil float64,
	limitglobal, limittotal int) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()
	msg := "Failed"
	flag := false

	sql_update := `
			UPDATE   
			` + config.DB_tbl_mst_company_game_pasaran + ` 
			SET 7_minbet=? , 7_maxbet=?, 
			7_keiasganjil=?, 7_keiasgenap=?, 7_keiasbesar=?, 7_keiaskecil=?, 
			7_keikopganjil=?, 7_keikopgenap=?, 7_keikopbesar=?, 7_keikopkecil=?, 
			7_keikepalaganjil=?, 7_keikepalagenap=?, 7_keikepalabesar=?, 7_keikepalakecil=?,  
			7_keiekorganjil=?, 7_keiekorgenap=?, 7_keiekorbesar=?, 7_keiekorkecil=?, 
			7_discasganjil=?, 7_discasgenap=?, 7_discasbesar=?, 7_discaskecil=?, 
			7_disckopganjil=?, 7_disckopgenap=?, 7_disckopbesar=?, 7_disckopkecil=?, 
			7_disckepalaganjil=?, 7_disckepalagenap=?, 7_disckepalabesar=?, 7_disckepalakecil=?, 
			7_discekorganjil=?, 7_discekorgenap=?, 7_discekorbesar=?, 7_discekorkecil=?, 
			7_limitbuang=?, 7_limittotal=?,   
			updatecomppas=?, updatedatecompas=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`

	rows_update, err_update := con.PrepareContext(ctx, sql_update)
	helpers.ErrorCheck(err_update)
	rec_comp, err_comp := rows_update.ExecContext(ctx,
		minbet, maxbet,
		helpers.ToFixed(keiasganjil, 3), helpers.ToFixed(keiasgenap, 3), helpers.ToFixed(keiasbesar, 3), helpers.ToFixed(keiaskecil, 3),
		helpers.ToFixed(keikopganjil, 3), helpers.ToFixed(keikopgenap, 3), helpers.ToFixed(keikopbesar, 3), helpers.ToFixed(keikopkecil, 3),
		helpers.ToFixed(keikepalaganjil, 3), helpers.ToFixed(keikepalagenap, 3), helpers.ToFixed(keikepalabesar, 3), helpers.ToFixed(keikepalakecil, 3),
		helpers.ToFixed(keiekorganjil, 3), helpers.ToFixed(keiekorgenap, 3), helpers.ToFixed(keiekorbesar, 3), helpers.ToFixed(keiekorkecil, 3),
		helpers.ToFixed(discasganjil, 3), helpers.ToFixed(discasgenap, 3), helpers.ToFixed(discasbesar, 3), helpers.ToFixed(discaskecil, 3),
		helpers.ToFixed(disckopganjil, 3), helpers.ToFixed(disckopgenap, 3), helpers.ToFixed(disckopbesar, 3), helpers.ToFixed(disckopkecil, 3),
		helpers.ToFixed(disckepalaganjil, 3), helpers.ToFixed(disckepalagenap, 3), helpers.ToFixed(disckepalabesar, 3), helpers.ToFixed(disckepalakecil, 3),
		helpers.ToFixed(discekorganjil, 3), helpers.ToFixed(discekorgenap, 3), helpers.ToFixed(discekorbesar, 3), helpers.ToFixed(discekorkecil, 3),
		limitglobal, limittotal,
		master, tglnow.Format("YYYY-MM-DD HH:mm:ss"),
		idcomppasaran, company)
	helpers.ErrorCheck(err_comp)
	update_comp, err_comp := rec_comp.RowsAffected()
	helpers.ErrorCheck(err_comp)
	defer rows_update.Close()
	log.Println(update_comp)
	if update_comp > 0 {
		flag = true
		msg = "Success"
		log.Printf("Update %s Success : %d\n", config.DB_tbl_mst_company_game_pasaran, idcomppasaran)
	} else {
		flag = false
		msg = "Failed"
		log.Printf("Update %s Failed \n", config.DB_tbl_mst_company_game_pasaran)
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
func Save_companyUpdatePasaran5050kombinasi(
	master, company string, idcomppasaran int,
	minbet, maxbet int,
	belakangkeimono, belakangkeistereo, belakangkeikembang, belakangkeikempis, belakangkeikembar float64,
	tengahkeimono, tengahkeistereo, tengahkeikembang, tengahkeikempis, tengahkeikembar float64,
	depankeimono, depankeistereo, depankeikembang, depankeikempis, depankeikembar float64,
	belakangdiscmono, belakangdiscstereo, belakangdisckembang, belakangdisckempis, belakangdisckembar float64,
	tengahdiscmono, tengahdiscstereo, tengahdisckembang, tengahdisckempis, tengahdisckembar float64,
	depandiscmono, depandiscstereo, depandisckembang, depandisckempis, depandisckembar float64,
	limitglobal, limittotal int) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()
	msg := "Failed"
	flag := false

	sql_update := `
			UPDATE   
			` + config.DB_tbl_mst_company_game_pasaran + ` 
			SET 8_minbet=? , 8_maxbet=?, 
			8_belakangkeimono=?, 8_belakangkeistereo=?, 8_belakangkeikembang=?, 8_belakangkeikempis=?, 8_belakangkeikembar=?, 
			8_tengahkeimono=?, 8_tengahkeistereo=?, 8_tengahkeikembang=?, 8_tengahkeikempis=?, 8_tengahkeikembar=?, 
			8_depankeimono=?, 8_depankeistereo=?, 8_depankeikembang=?, 8_depankeikempis=?, 8_depankeikembar=?, 
			8_belakangdiscmono=?, 8_belakangdiscstereo=?, 8_belakangdisckembang=?, 8_belakangdisckempis=?, 8_belakangdisckembar=?, 
			8_tengahdiscmono=?, 8_tengahdiscstereo=?, 8_tengahdisckembang=?, 8_tengahdisckempis=?, 8_tengahdisckembar=?, 
			8_depandiscmono=?, 8_depandiscstereo=?, 8_depandisckembang=?, 8_depandisckempis=?, 8_depandisckembar=?, 
			8_limitbuang=?, 8_limittotal=?, 
			updatecomppas=?, updatedatecompas=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`

	rows_update, err_update := con.PrepareContext(ctx, sql_update)
	helpers.ErrorCheck(err_update)
	rec_comp, err_comp := rows_update.ExecContext(ctx,
		minbet, maxbet,
		helpers.ToFixed(belakangkeimono, 3), helpers.ToFixed(belakangkeistereo, 3), helpers.ToFixed(belakangkeikembang, 3), helpers.ToFixed(belakangkeikempis, 3), helpers.ToFixed(belakangkeikembar, 3),
		helpers.ToFixed(tengahkeimono, 3), helpers.ToFixed(tengahkeistereo, 3), helpers.ToFixed(tengahkeikembang, 3), helpers.ToFixed(tengahkeikempis, 3), helpers.ToFixed(tengahkeikembar, 3),
		helpers.ToFixed(depankeimono, 3), helpers.ToFixed(depankeistereo, 3), helpers.ToFixed(depankeikembang, 3), helpers.ToFixed(depankeikempis, 3), helpers.ToFixed(depankeikembar, 3),
		helpers.ToFixed(belakangdiscmono, 3), helpers.ToFixed(belakangdiscstereo, 3), helpers.ToFixed(belakangdisckembang, 3), helpers.ToFixed(belakangdisckempis, 3), helpers.ToFixed(belakangdisckembar, 3),
		helpers.ToFixed(tengahdiscmono, 3), helpers.ToFixed(tengahdiscstereo, 3), helpers.ToFixed(tengahdisckembang, 3), helpers.ToFixed(tengahdisckempis, 3), helpers.ToFixed(tengahdisckembar, 3),
		helpers.ToFixed(depandiscmono, 3), helpers.ToFixed(depandiscstereo, 3), helpers.ToFixed(depandisckembang, 3), helpers.ToFixed(depandisckempis, 3), helpers.ToFixed(depandisckembar, 3),
		limitglobal, limittotal,
		master, tglnow.Format("YYYY-MM-DD HH:mm:ss"),
		idcomppasaran, company)
	helpers.ErrorCheck(err_comp)
	update_comp, err_comp := rec_comp.RowsAffected()
	helpers.ErrorCheck(err_comp)
	defer rows_update.Close()
	log.Println(update_comp)
	if update_comp > 0 {
		flag = true
		msg = "Success"
		log.Printf("Update %s Success : %d\n", config.DB_tbl_mst_company_game_pasaran, idcomppasaran)
	} else {
		flag = false
		msg = "Failed"
		log.Printf("Update %s Failed \n", config.DB_tbl_mst_company_game_pasaran)
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
func Save_companyUpdatePasarankombinasi(
	master, company string, idcomppasaran int,
	minbet, maxbet int,
	win, disc float32,
	limitglobal, limittotal int) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()
	msg := "Failed"
	flag := false

	sql_update := `
			UPDATE   
			` + config.DB_tbl_mst_company_game_pasaran + ` 
			SET 9_minbet=? , 9_maxbet=?, 9_win=?, 9_discount=?, 
			9_limitbuang=?, 9_limittotal=?, 
			updatecomppas=?, updatedatecompas=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`

	rows_update, err_update := con.PrepareContext(ctx, sql_update)
	helpers.ErrorCheck(err_update)
	rec_comp, err_comp := rows_update.ExecContext(ctx,
		minbet,
		maxbet,
		win,
		disc,
		limitglobal,
		limittotal,
		master, tglnow.Format("YYYY-MM-DD HH:mm:ss"),
		idcomppasaran, company)
	helpers.ErrorCheck(err_comp)
	update_comp, err_comp := rec_comp.RowsAffected()
	helpers.ErrorCheck(err_comp)
	defer rows_update.Close()
	log.Println(update_comp)
	if update_comp > 0 {
		flag = true
		msg = "Success"
		log.Printf("Update %s Success : %d\n", config.DB_tbl_mst_company_game_pasaran, idcomppasaran)
	} else {
		flag = false
		msg = "Failed"
		log.Printf("Update %s Failed \n", config.DB_tbl_mst_company_game_pasaran)
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
func Save_companyUpdatePasarandasar(
	master, company string, idcomppasaran int,
	minbet, maxbet int,
	keibesar, keikecil, keigenap, keiganjil float32,
	discbesar, disckecil, discigenap, discganjil float32,
	limitglobal, limittotal int) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()
	msg := "Failed"
	flag := false

	sql_update := `
			UPDATE   
			` + config.DB_tbl_mst_company_game_pasaran + ` 
			SET 10_minbet=? , 10_maxbet=?, 
			10_keibesar=?, 10_keikecil=?, 10_keigenap=?, 10_keiganjil=?, 
			10_discbesar=?, 10_disckecil=?, 10_discigenap=?, 10_discganjil=?, 
			10_limitbuang=?, 10_limittotal=?, 
			updatecomppas=?, updatedatecompas=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`

	rows_update, err_update := con.PrepareContext(ctx, sql_update)
	helpers.ErrorCheck(err_update)
	rec_comp, err_comp := rows_update.ExecContext(ctx,
		minbet,
		maxbet,
		keibesar, keikecil, keigenap, keiganjil,
		discbesar, disckecil, discigenap, discganjil,
		limitglobal,
		limittotal,
		master, tglnow.Format("YYYY-MM-DD HH:mm:ss"),
		idcomppasaran, company)
	helpers.ErrorCheck(err_comp)
	update_comp, err_comp := rec_comp.RowsAffected()
	helpers.ErrorCheck(err_comp)
	defer rows_update.Close()
	log.Println(update_comp)
	if update_comp > 0 {
		flag = true
		msg = "Success"
		log.Printf("Update %s Success : %d\n", config.DB_tbl_mst_company_game_pasaran, idcomppasaran)
	} else {
		flag = false
		msg = "Failed"
		log.Printf("Update %s Failed \n", config.DB_tbl_mst_company_game_pasaran)
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
func Save_companyUpdatePasaranshio(
	master, company string, idcomppasaran int,
	shiotahunini string,
	minbet, maxbet int,
	win, disc float32,
	limitglobal, limittotal int) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()
	msg := "Failed"
	flag := false

	sql_update := `
			UPDATE   
			` + config.DB_tbl_mst_company_game_pasaran + ` 
			SET 11_shiotahunini=? , 11_minbet=?, 11_maxbet=?, 
			11_win=?, 11_disc=?, 
			11_limitbuang=?, 11_limittotal=?, 
			updatecomppas=?, updatedatecompas=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`

	rows_update, err_update := con.PrepareContext(ctx, sql_update)
	helpers.ErrorCheck(err_update)
	rec_comp, err_comp := rows_update.ExecContext(ctx,
		shiotahunini,
		minbet,
		maxbet,
		win,
		disc,
		limitglobal,
		limittotal,
		master, tglnow.Format("YYYY-MM-DD HH:mm:ss"),
		idcomppasaran, company)
	helpers.ErrorCheck(err_comp)
	update_comp, err_comp := rec_comp.RowsAffected()
	helpers.ErrorCheck(err_comp)
	defer rows_update.Close()
	log.Println(update_comp)
	if update_comp > 0 {
		flag = true
		msg = "Success"
		log.Printf("Update %s Success : %d\n", config.DB_tbl_mst_company_game_pasaran, idcomppasaran)
	} else {
		flag = false
		msg = "Failed"
		log.Printf("Update %s Failed \n", config.DB_tbl_mst_company_game_pasaran)
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
func Save_companyInsertPasaranharionline(master, company string, idcomppasaran int, hari string) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()
	msg := "Failed"
	flag := false

	flag = Get_OnlinePasaran(company, idcomppasaran, hari, "hari")
	if !flag {
		year, _ := strconv.Atoi(tglnow.Format("YYYY"))
		field_col := config.DB_tbl_mst_company_game_pasaran_offline + "_" + tglnow.Format("YYYY")
		idrecord_counter := year + Get_counter(field_col)
		sql_insert := `
			INSERT INTO  
			` + config.DB_tbl_mst_company_game_pasaran_offline + ` (
				idcomppasaranoff, idcomppasaran, idcompany, haripasaran, createcomppasaranoff, createdatecomppasaranoff 
			)VALUES( 
				?,?,?,?,?,?
			) 
		`

		rows_insert, err_insert := con.PrepareContext(ctx, sql_insert)
		helpers.ErrorCheck(err_insert)
		rec_comp, err_comp := rows_insert.ExecContext(ctx,
			idrecord_counter,
			idcomppasaran,
			company,
			hari,
			master, tglnow.Format("YYYY-MM-DD HH:mm:ss"))
		helpers.ErrorCheck(err_comp)
		insert_comp, err_comp := rec_comp.RowsAffected()
		helpers.ErrorCheck(err_comp)
		defer rows_insert.Close()
		if insert_comp > 0 {
			flag = true
			msg = "Success"
			log.Printf("Update %s Success : %d\n", config.DB_tbl_mst_company_game_pasaran_offline, idcomppasaran)
		} else {
			flag = false
			msg = "Failed"
			log.Printf("Update %s Failed \n", config.DB_tbl_mst_company_game_pasaran_offline)
		}
	} else {
		msg = "Duplicate Entry"
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
func Delete_companyPasaranharionline(master, company string, idcomppasaran, idcomppasaranoff int) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	msg := "Failed"
	flag := false
	sql_delete := `
			DELETE FROM  
			` + config.DB_tbl_mst_company_game_pasaran_offline + ` 
			WHERE idcomppasaranoff=? AND idcomppasaran=? AND idcompany=? 
		`

	rows_delete, err_delete := con.PrepareContext(ctx, sql_delete)
	helpers.ErrorCheck(err_delete)
	rec_comp, err_comp := rows_delete.ExecContext(ctx, idcomppasaranoff, idcomppasaran, company)
	helpers.ErrorCheck(err_comp)
	delete_comp, err_comp := rec_comp.RowsAffected()
	helpers.ErrorCheck(err_comp)
	defer rows_delete.Close()
	if delete_comp > 0 {
		flag = true
		msg = "Success"
		log.Printf("Delete %s Success : %d\n", config.DB_tbl_mst_company_game_pasaran_offline, idcomppasaran)
	} else {
		flag = false
		msg = "Failed"
		log.Printf("Delete %s Failed \n", config.DB_tbl_mst_company_game_pasaran_offline)
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
func _winlose(company, start, end string, idcomppasaran int) int {
	con := db.CreateCon()
	ctx := context.Background()
	var winlose float64 = 0
	tbl_trx_keluarantogel, _, _, _ := Get_mappingdatabase(company)
	if idcomppasaran > 0 {
		sql_keluaran := `SELECT
			COALESCE(SUM(total_outstanding-total_cancel-winlose),0 )  as winlose
			FROM ` + tbl_trx_keluarantogel + `  
			WHERE idcompany = ? 
			AND idcomppasaran = ? 
			AND datekeluaran >= ? 
			AND datekeluaran <= ? 
			AND keluarantogel != ''  
		`
		row := con.QueryRowContext(ctx, sql_keluaran, company, idcomppasaran, start, end)
		switch e := row.Scan(&winlose); e {
		case sql.ErrNoRows:

		case nil:

		default:
			panic(e)
		}
		return int(winlose)
	} else {
		sql_keluaran := `SELECT
			COALESCE(SUM(total_outstanding-total_cancel-winlose),0 )  as winlose
			FROM ` + tbl_trx_keluarantogel + `  
			WHERE idcompany = ? 
			AND datekeluaran >= ? 
			AND datekeluaran <= ? 
			AND keluarantogel != ''  
		`
		row := con.QueryRowContext(ctx, sql_keluaran, company, start, end)
		switch e := row.Scan(&winlose); e {
		case sql.ErrNoRows:

		case nil:

		default:
			panic(e)
		}
		return int(winlose)
	}
}
func _winlosetemp(company, start, end string, invoice int) int {
	con := db.CreateCon()
	ctx := context.Background()
	var winlose float64 = 0
	tbl_trx_keluarantogel, _, tbl_trx_keluaranmember, _ := Get_mappingdatabase(company)
	if invoice > 0 {
		sql_keluaran := `SELECT
			COALESCE(SUM(totalbayar-totalcancel-totalwin),0 )  as winlose
			FROM ` + tbl_trx_keluaranmember + ` 
			WHERE idtrxkeluaran = ? 
			AND idcompany = ?  
		`
		row := con.QueryRowContext(ctx, sql_keluaran, invoice, company)
		switch e := row.Scan(&winlose); e {
		case sql.ErrNoRows:

		case nil:

		default:
			panic(e)
		}
		return int(winlose)
	} else {
		sql_keluaran := `SELECT
			COALESCE(SUM(A.totalbayar-A.totalcancel-A.totalwin),0 )  as winlose
			FROM ` + tbl_trx_keluaranmember + ` as A  
			JOIN ` + tbl_trx_keluarantogel + ` as B ON B.idtrxkeluaran = A.idtrxkeluaran  
			WHERE B.idcompany = ? 
			AND B.datekeluaran >= ? 
			AND B.datekeluaran <= ? 
			AND B.keluarantogel != ''  
		`
		row := con.QueryRowContext(ctx, sql_keluaran, company, start, end)
		switch e := row.Scan(&winlose); e {
		case sql.ErrNoRows:

		case nil:

		default:
			panic(e)
		}
		return int(winlose)
	}
}
func _rumuswinhasil(permainan string, bayar int, bet int, win float32) int {
	winhasil := 0
	if permainan == "50_50_UMUM" || permainan == "50_50_SPECIAL" ||
		permainan == "50_50_KOMBINASI" || permainan == "DASAR" || permainan == "COLOK_BEBAS" ||
		permainan == "COLOK_MACAU" || permainan == "COLOK_NAGA" || permainan == "COLOK_JITU" {

		winhasil = bayar + int(float32(bet)*win)
	} else {
		winhasil = int(float32(bet) * win)
	}
	return winhasil
}
