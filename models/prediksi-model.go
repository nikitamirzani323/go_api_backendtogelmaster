package models

import (
	"context"
	"database/sql"
	"fmt"
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

func Fetch_listpasaranwajib() (helpers.Response, error) {
	var obj entities.Model_pasaranlist
	var arraobj []entities.Model_pasaranlist
	var res helpers.Response
	msg := "Data Not Found"
	render_page := time.Now()
	con := db.CreateCon()
	ctx := context.Background()

	sql_periode := `SELECT 
			idpasarantogel, nmpasarantogel   
			FROM ` + config.DB_tbl_mst_pasaran_togel + ` 
			WHERE tipepasaran = 'WAJIB'
		`
	row, err := con.QueryContext(ctx, sql_periode)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idpasarantogel_db, nmpasarantogel_db string
		)

		err = row.Scan(&idpasarantogel_db, &nmpasarantogel_db)

		obj.Pasaranlist_idpasarantogel = idpasarantogel_db
		obj.Pasaranlist_nmpasarantogel = nmpasarantogel_db
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
func Fetch_prediksipasaranwajib(idpasarantogel, nomorprediksi string) (helpers.Response, error) {
	var obj entities.Model_prediksiwajib
	var arraobj []entities.Model_prediksiwajib
	var res helpers.Response
	msg := "Data Not Found"
	render_page := time.Now()
	con := db.CreateCon()
	ctx := context.Background()
	flag := false

	sql_select := `SELECT 
				A.idcomppasaran, A.idcompany, B.nmcompany    
				FROM ` + config.DB_tbl_mst_company_game_pasaran + ` as A 
				JOIN ` + config.DB_tbl_mst_company + ` as B ON B.idcompany = A.idcompany  
				WHERE A.idpasarantogel = ?
				AND B.statuscompany = 'ACTIVE' 
			`
	row, err := con.QueryContext(ctx, sql_select, idpasarantogel)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idcomppasaran_db           int
			idcompany_db, nmcompany_db string
		)

		err = row.Scan(&idcomppasaran_db, &idcompany_db, &nmcompany_db)

		result, totalbet, subtotal, subtotalwin := _listprediksi(idcompany_db, nomorprediksi, idcomppasaran_db)
		invoice, dateinvoice, periodeinvoice := _listprediksiparent(idcompany_db, idcomppasaran_db)

		obj.Prediksi_idcompany = idcompany_db
		obj.Prediksi_nmcompany = nmcompany_db
		obj.Prediksi_invoice = invoice
		obj.Prediksi_invoicedate = dateinvoice
		obj.Prediksi_invoiceperiode = periodeinvoice
		obj.Prediksi_result = result
		obj.Prediksi_totalbet = totalbet
		obj.Prediksi_subtotal = subtotal
		obj.Prediksi_subtotalwin = subtotalwin
		arraobj = append(arraobj, obj)
		msg = "Success"
		flag = true
	}
	defer row.Close()

	if flag {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = arraobj
		res.Time = time.Since(render_page).String()
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = msg
		res.Record = arraobj
		res.Time = time.Since(render_page).String()
	}

	return res, nil
}
func _listprediksiparent(company string, idcomppasaran int) (string, string, string) {
	ctx := context.Background()
	con := db.CreateCon()
	idtrxkeluaran := ""
	datekeluaran := ""
	periode := ""
	tbl_trx_keluarantogel, _, _, _ := Get_mappingdatabase(company)

	sql_parent := `SELECT
					idtrxkeluaran , datekeluaran, idcomppasaran , keluaranperiode 
					FROM ` + tbl_trx_keluarantogel + ` 
					WHERE idcompany = ? 
					AND idcomppasaran = ?  
					AND keluarantogel = '' 
					ORDER BY idtrxkeluaran DESC   
				`
	row, err := con.QueryContext(ctx, sql_parent, company, idcomppasaran)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idcomppasaran_db                                      int
			idtrxkeluaran_db, datekeluaran_db, keluaranperiode_db string
		)

		err = row.Scan(&idtrxkeluaran_db, &datekeluaran_db, &idcomppasaran_db, &keluaranperiode_db)
		helpers.ErrorCheck(err)
		idtrxkeluaran = idtrxkeluaran_db
		datekeluaran = datekeluaran_db
		idpasarantogel, _ := _pasaran_id(idcomppasaran_db, company, "idpasarantogel")
		periode = idpasarantogel + "-" + keluaranperiode_db

	}
	return idtrxkeluaran, datekeluaran, periode
}
func _listprediksi(company, nomorkeluaran string, idcomppasaran int) (interface{}, int, int, int) {
	var obj entities.Model_listPrediksi
	var arraobj []entities.Model_listPrediksi
	ctx := context.Background()
	con := db.CreateCon()
	totalbet := 0
	subtotal := 0
	subtotalwin := 0
	tbl_trx_keluarantogel, tbl_trx_keluarantogel_detail, _, _ := Get_mappingdatabase(company)

	sql_listprediksi := `SELECT
					A.datetimedetail , A.username , A.typegame, A.nomortogel, 
					A.bet, A.diskon, A.kei, A.win  
					FROM ` + tbl_trx_keluarantogel_detail + ` as A  
					JOIN ` + tbl_trx_keluarantogel + ` as B ON B.idtrxkeluaran = A.idtrxkeluaran  
					WHERE B.idcompany = ? 
					AND B.idcomppasaran = ?  
					AND B.keluarantogel = '' 
					ORDER BY B.idtrxkeluaran DESC   
				`
	row, err := con.QueryContext(ctx, sql_listprediksi, company, idcomppasaran)
	helpers.ErrorCheck(err)

	for row.Next() {
		var (
			bet_db                                                     int
			diskon_db, kei_db, win_db                                  float32
			datetimedetail_db, username_db, typegame_db, nomortogel_db string
		)

		err = row.Scan(
			&datetimedetail_db, &username_db, &typegame_db, &nomortogel_db,
			&bet_db, &diskon_db, &kei_db, &win_db)

		helpers.ErrorCheck(err)
		diskonpercen := diskon_db * 100
		diskonbet := int(float32(bet_db) * diskon_db)
		keipercen := kei_db * 100
		keibet := int(float32(bet_db) * kei_db)

		bayar := bet_db - int(float32(bet_db)*diskon_db) - int(float32(bet_db)*kei_db)
		subtotal = subtotal + bayar
		statuskeluarandetail, winrumus := _rumusTogel(nomorkeluaran, typegame_db, nomortogel_db, company, "N", idcomppasaran, 0)
		var winfixed float32 = 0
		winhasil := 0
		if typegame_db == "COLOK_BEBAS" || typegame_db == "COLOK_MACAU" || typegame_db == "COLOK_NAGA" {
			winhasil = _rumuswinhasil(typegame_db, bayar, bet_db, winrumus)
			winfixed = winrumus
		} else {
			winhasil = _rumuswinhasil(typegame_db, bayar, bet_db, win_db)
			winfixed = win_db
		}
		status_css := ""

		switch statuskeluarandetail {
		case "WINNER":
			totalbet = totalbet + 1
			subtotalwin = subtotalwin + winhasil
			status_css = "background:#8BC34A;color:black;font-weight:bold;"
		case "LOSE":
			status_css = "background:#E91E63;font-size:12px;font-weight:bold;color:white;"
		}
		if statuskeluarandetail == "WINNER" {
			obj.Prediksi_tanggal = datetimedetail_db
			obj.Prediksi_username = username_db
			obj.Prediksi_permainan = typegame_db
			obj.Prediksi_nomor = nomortogel_db
			obj.Prediksi_bet = bet_db
			obj.Prediksi_diskon = diskonbet
			obj.Prediksi_diskonpercen = diskonpercen
			obj.Prediksi_kei = keibet
			obj.Prediksi_keipercen = keipercen
			obj.Prediksi_bayar = bayar
			obj.Prediksi_win = winfixed
			obj.Prediksi_totalwin = winhasil
			obj.Prediksi_status = statuskeluarandetail
			obj.Prediksi_statuscss = status_css
			arraobj = append(arraobj, obj)
		}
	}
	defer row.Close()

	return arraobj, totalbet, subtotalwin, subtotal - subtotalwin
}
func _rumusTogel(angka, tipe, nomorkeluaran, company, simpandb string, idcomppasaran, idtrxkeluarandetail int) (string, float32) {
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	var result string = "LOSE"
	var win float32 = 0

	_, _, tbl_trx_keluarantogel_detail, _ := Get_mappingdatabase(company)

	temp := angka
	temp4d := string([]byte(temp)[0]) + string([]byte(temp)[1]) + string([]byte(temp)[2]) + string([]byte(temp)[3])
	temp3d := string([]byte(temp)[1]) + string([]byte(temp)[2]) + string([]byte(temp)[3])
	temp2d := string([]byte(temp)[2]) + string([]byte(temp)[3])
	temp2dd := string([]byte(temp)[0]) + string([]byte(temp)[1])
	temp2dt := string([]byte(temp)[1]) + string([]byte(temp)[2])

	switch tipe {
	case "4D":
		if temp4d == nomorkeluaran {
			result = "WINNER"
		}
	case "3D":
		if temp3d == nomorkeluaran {
			result = "WINNER"
		}
	case "2D":
		if temp2d == nomorkeluaran {
			result = "WINNER"
		}
	case "2DD":
		if temp2dd == nomorkeluaran {
			result = "WINNER"
		}
	case "2DT":
		if temp2dt == nomorkeluaran {
			result = "WINNER"
		}
	case "COLOK_BEBAS":
		flag := false
		count := 0
		for i := 0; i < len(temp); i++ {
			if string([]byte(temp)[i]) == nomorkeluaran {
				flag = true
				count = count + 1
			}
		}
		if flag {
			_, win_db := _pasaran_id(idcomppasaran, company, "2_win")
			if count == 1 {
				win = win_db
			}
			if count == 2 {
				win = win_db * 2
			}
			if count == 3 {
				win = win_db * 3
			}
			if count == 4 {
				win = win_db * 3
			}
			fmt.Println(win)

			if simpandb == "Y" {
				//UPDATE WIN DETAIL BET
				stmt, e := con.PrepareContext(ctx, `
					UPDATE 
					`+tbl_trx_keluarantogel_detail+`     
					SET win=?, 
					updatekeluarandetail=?, updatedatekeluarandetail=? 
					WHERE idtrxkeluarandetail=? 
				`)
				helpers.ErrorCheck(e)
				rec, e := stmt.ExecContext(ctx,
					win,
					"SYSTEM",
					tglnow.Format("YYYY-MM-DD HH:mm:ss"),
					idtrxkeluarandetail)
				helpers.ErrorCheck(e)

				a, e := rec.RowsAffected()
				helpers.ErrorCheck(e)
				fmt.Printf("The last id: %d\n", a)
				defer stmt.Close()
			}
			result = "WINNER"
		}
	case "COLOK_MACAU":
		flag_1 := false
		flag_2 := false
		count_1 := 0
		count_2 := 0
		totalcount := 0
		var win float32 = 0
		for i := 0; i < len(temp); i++ {
			if string([]byte(temp)[i]) == string([]byte(nomorkeluaran)[0]) {
				flag_1 = true
				count_1 = count_1 + 1
			}
			if string([]byte(temp)[i]) == string([]byte(nomorkeluaran)[1]) {
				flag_2 = true
				count_2 = count_2 + 1
			}
		}
		if flag_1 && flag_2 {
			totalcount = count_1 + count_2
			if totalcount == 2 {
				_, win = _pasaran_id(idcomppasaran, company, "3_win2digit")
			}
			if totalcount == 3 {
				_, win = _pasaran_id(idcomppasaran, company, "3_win3digit")
			}
			if totalcount == 4 {
				_, win = _pasaran_id(idcomppasaran, company, "3_win4digit")
			}
			if simpandb == "Y" {
				//UPDATE WIN DETAIL BET
				stmt, e := con.PrepareContext(ctx, `
					UPDATE 
					`+tbl_trx_keluarantogel_detail+`     
					SET win=?, 
					updatekeluarandetail=?, updatedatekeluarandetail=? 
					WHERE idtrxkeluarandetail=? 
				`)
				helpers.ErrorCheck(e)
				rec, e := stmt.ExecContext(ctx,
					win,
					"SYSTEM",
					tglnow.Format("YYYY-MM-DD HH:mm:ss"),
					idtrxkeluarandetail)
				helpers.ErrorCheck(e)

				a, e := rec.RowsAffected()
				helpers.ErrorCheck(e)
				fmt.Printf("The last id: %d\n", a)
				defer stmt.Close()
				fmt.Println(win)
			}
			result = "WINNER"
		}
	case "COLOK_NAGA":
		flag_1 := false
		flag_2 := false
		flag_3 := false
		count_1 := 0
		count_2 := 0
		count_3 := 0
		totalcount := 0
		var win float32 = 0
		for i := 0; i < len(temp); i++ {
			if string([]byte(temp)[i]) == string([]byte(nomorkeluaran)[0]) {
				flag_1 = true
				count_1 = count_1 + 1
			}
			if string([]byte(temp)[i]) == string([]byte(nomorkeluaran)[1]) {
				flag_2 = true
				count_2 = count_2 + 1
			}
			if string([]byte(temp)[i]) == string([]byte(nomorkeluaran)[2]) {
				flag_3 = true
				count_3 = count_3 + 1
			}
		}
		if flag_1 && flag_2 {
			if flag_3 {
				totalcount = count_1 + count_2 + count_3

				if totalcount == 3 {
					_, win = _pasaran_id(idcomppasaran, company, "4_win3digit")
				}
				if totalcount == 4 {
					_, win = _pasaran_id(idcomppasaran, company, "4_win4digit")
				}
				if simpandb == "Y" {
					//UPDATE WIN DETAIL BET
					stmt, e := con.PrepareContext(ctx, `
						UPDATE 
						`+tbl_trx_keluarantogel_detail+`     
						SET win=?,  
						updatekeluarandetail=?, updatedatekeluarandetail=? 
						WHERE idtrxkeluarandetail=? 
					`)
					helpers.ErrorCheck(e)
					rec, e := stmt.ExecContext(ctx,
						win,
						"SYSTEM",
						tglnow.Format("YYYY-MM-DD HH:mm:ss"),
						idtrxkeluarandetail)
					helpers.ErrorCheck(e)

					a, e := rec.RowsAffected()
					helpers.ErrorCheck(e)
					fmt.Printf("The last id: %d\n", a)
					defer stmt.Close()
					fmt.Println(win)
					fmt.Println(win)
				}
				result = "WINNER"
			}
		}
	case "COLOK_JITU":
		flag := false
		as := string([]byte(temp)[0]) + "_AS"
		kop := string([]byte(temp)[1]) + "_KOP"
		kepala := string([]byte(temp)[2]) + "_KEPALA"
		ekor := string([]byte(temp)[3]) + "_KEKOR"

		if as == nomorkeluaran {
			flag = true
		}
		if kop == nomorkeluaran {
			flag = true
		}
		if kepala == nomorkeluaran {
			flag = true
		}
		if ekor == nomorkeluaran {
			flag = true
		}
		if flag {
			result = "WINNER"
		}
	case "50_50_UMUM":
		flag := false
		data := []string{}
		kepala := string([]byte(temp)[2])
		ekor := string([]byte(temp)[3])
		kepala_2, _ := strconv.Atoi(kepala)
		ekor_2, _ := strconv.Atoi(ekor)
		dasar := kepala_2 + ekor_2
		//BESARKECIL
		if kepala_2 <= 4 {
			data = append(data, "KECIL")
		} else {
			data = append(data, "BESAR")
		}
		//GENAPGANJIL
		if ekor_2%2 == 0 {
			data = append(data, "GENAP")
		} else {
			data = append(data, "GANJIL")
		}
		//TEPITENGAH
		if dasar >= 0 && dasar <= 24 {
			data = append(data, "TEPI")
		}
		if dasar >= 25 && dasar <= 74 {
			data = append(data, "TENGAH")
		}
		if dasar >= 75 && dasar <= 99 {
			data = append(data, "TEPI")
		}
		for i := 0; i < len(data); i++ {
			if data[i] == nomorkeluaran {
				flag = true
			}
		}
		if flag {
			result = "WINNER"
		}
		fmt.Println(data)
	case "50_50_SPECIAL":
		flag := false
		as := string([]byte(temp)[0])
		kop := string([]byte(temp)[1])
		kepala := string([]byte(temp)[2])
		ekor := string([]byte(temp)[3])

		as_2, _ := strconv.Atoi(as)
		kop_2, _ := strconv.Atoi(kop)
		kepala_2, _ := strconv.Atoi(kepala)
		ekor_2, _ := strconv.Atoi(ekor)
		//AS - BESARKECIL == GENAPGANJIL
		if as_2 <= 4 {
			if nomorkeluaran == "AS_KECIL" {
				flag = true
			}
		} else {
			if nomorkeluaran == "AS_BESAR" {
				flag = true
			}
		}
		if as_2%2 == 0 {
			if nomorkeluaran == "AS_GENAP" {
				flag = true
			}
		} else {
			if nomorkeluaran == "AS_GANJIL" {
				flag = true
			}
		}

		//KOP - BESARKECIL == GENAPGANJIL
		if kop_2 <= 4 {
			if nomorkeluaran == "KOP_KECIL" {
				flag = true
			}
		} else {
			if nomorkeluaran == "KOP_BESAR" {
				flag = true
			}
		}
		if kop_2%2 == 0 {
			if nomorkeluaran == "KOP_GENAP" {
				flag = true
			}
		} else {
			if nomorkeluaran == "KOP_GANJIL" {
				flag = true
			}
		}

		//KEPALA - BESARKECIL == GENAPGANJIL
		if kepala_2 <= 4 {
			if nomorkeluaran == "KEPALA_KECIL" {
				flag = true
			}
		} else {
			if nomorkeluaran == "KEPALA_BESAR" {
				flag = true
			}
		}
		if kepala_2%2 == 0 {
			if nomorkeluaran == "KEPALA_GENAP" {
				flag = true
			}
		} else {
			if nomorkeluaran == "KEPALA_GANJIL" {
				flag = true
			}
		}

		//EKOR - BESARKECIL == GENAPGANJIL
		if ekor_2 <= 4 {
			if nomorkeluaran == "EKOR_KECIL" {
				flag = true
			}
		} else {
			if nomorkeluaran == "EKOR_BESAR" {
				flag = true
			}
		}
		if ekor_2%2 == 0 {
			if nomorkeluaran == "EKOR_GENAP" {
				flag = true
			}
		} else {
			if nomorkeluaran == "EKOR_GANJIL" {
				flag = true
			}
		}

		if flag {
			result = "WINNER"
		}
	case "50_50_KOMBINASI":
		flag := false
		data_1 := ""
		data_2 := ""
		data_3 := ""
		data_4 := ""
		depan := ""
		tengah := ""
		belakang := ""
		depan_1 := ""
		tengah_1 := ""
		belakang_1 := ""
		as := string([]byte(temp)[0])
		kop := string([]byte(temp)[1])
		kepala := string([]byte(temp)[2])
		ekor := string([]byte(temp)[3])

		as_2, _ := strconv.Atoi(as)
		kop_2, _ := strconv.Atoi(kop)
		kepala_2, _ := strconv.Atoi(kepala)
		ekor_2, _ := strconv.Atoi(ekor)

		if as_2%2 == 0 {
			data_1 = "GENAP"
		} else {
			data_1 = "GANJIL"
		}
		if kop_2%2 == 0 {
			data_2 = "GENAP"
		} else {
			data_2 = "GANJIL"
		}
		if kepala_2%2 == 0 {
			data_3 = "GENAP"
		} else {
			data_3 = "GANJIL"
		}
		if ekor_2%2 == 0 {
			data_4 = "GENAP"
		} else {
			data_4 = "GANJIL"
		}
		depan = data_1 + "-" + data_2
		tengah = data_2 + "-" + data_3
		belakang = data_3 + "-" + data_4

		if depan == "GENAP-GANJIL" || depan == "GANJIL-GENAP" {
			depan = "DEPAN_STEREO"
		} else {
			depan = "DEPAN_MONO"
		}
		if tengah == "GENAP-GANJIL" || tengah == "GANJIL-GENAP" {
			tengah = "TENGAH_STEREO"
		} else {
			tengah = "TENGAH_MONO"
		}
		if belakang == "GENAP-GANJIL" || belakang == "GANJIL-GENAP" {
			belakang = "BELAKANG_STEREO"
		} else {
			belakang = "BELAKANG_MONO"
		}
		if as_2 < kop_2 {
			depan_1 = "DEPAN_KEMBANG"
		}
		if as_2 > kop_2 {
			depan_1 = "DEPAN_KEMPIS"
		}
		if as_2 == kop_2 {
			depan_1 = "DEPAN_KEMBAR"
		}
		if kop_2 < kepala_2 {
			tengah_1 = "TENGAH_KEMBANG"
		}
		if kop_2 > kepala_2 {
			tengah_1 = "TENGAH_KEMPIS"
		}
		if kop_2 == kepala_2 {
			tengah_1 = "TENGAH_KEMBAR"
		}
		if kepala_2 < ekor_2 {
			belakang_1 = "BELAKANG_KEMBANG"
		}
		if kepala_2 > ekor_2 {
			belakang_1 = "BELAKANG_KEMPIS"
		}
		if kepala_2 == ekor_2 {
			belakang_1 = "BELAKANG_KEMBAR"
		}

		if depan == nomorkeluaran {
			flag = true
		}
		if tengah == nomorkeluaran {
			flag = true
		}
		if belakang == nomorkeluaran {
			flag = true
		}
		if depan_1 == nomorkeluaran {
			flag = true
		}
		if tengah_1 == nomorkeluaran {
			flag = true
		}
		if belakang_1 == nomorkeluaran {
			flag = true
		}

		if flag {
			result = "WINNER"
		}
	case "MACAU_KOMBINASI":
		flag := false
		data_1 := ""
		data_2 := ""
		data_3 := ""
		data_4 := ""
		depan := ""
		tengah := ""
		belakang := ""

		as := string([]byte(temp)[0])
		kop := string([]byte(temp)[1])
		kepala := string([]byte(temp)[2])
		ekor := string([]byte(temp)[3])

		as_2, _ := strconv.Atoi(as)
		kop_2, _ := strconv.Atoi(kop)
		kepala_2, _ := strconv.Atoi(kepala)
		ekor_2, _ := strconv.Atoi(ekor)

		if as_2 <= 4 {
			data_1 = "KECIL"
		} else {
			data_1 = "BESAR"
		}
		if kop_2%2 == 0 {
			data_2 = "GENAP"
		} else {
			data_2 = "GANJIL"
		}
		if kepala_2 <= 4 {
			data_3 = "KECIL"
		} else {
			data_3 = "BESAR"
		}
		if ekor_2%2 == 0 {
			data_4 = "GENAP"
		} else {
			data_4 = "GANJIL"
		}

		depan = "DEPAN_" + data_1 + "_" + data_2
		tengah = "TENGAH_" + data_2 + "_" + data_3
		belakang = "BELAKANG_" + data_3 + "_" + data_4

		if depan == nomorkeluaran {
			flag = true
		}
		if tengah == nomorkeluaran {
			flag = true
		}
		if belakang == nomorkeluaran {
			flag = true
		}

		if flag {
			result = "WINNER"
		}
	case "DASAR":
		flag := false
		data_1 := ""
		data_2 := ""

		kepala := string([]byte(temp)[2])
		ekor := string([]byte(temp)[3])

		kepala_2, _ := strconv.Atoi(kepala)
		ekor_2, _ := strconv.Atoi(ekor)

		dasar := kepala_2 + ekor_2

		if dasar > 9 {
			temp2 := strconv.Itoa(dasar) //int to string
			temp21 := string([]byte(temp2)[0])
			temp22 := string([]byte(temp2)[1])

			temp21_2, _ := strconv.Atoi(temp21)
			temp22_2, _ := strconv.Atoi(temp22)
			dasar = temp21_2 + temp22_2
		}
		if dasar <= 4 {
			data_1 = "KECIL"
		} else {
			data_1 = "BESAR"
		}
		if dasar%2 == 0 {
			data_2 = "GENAP"
		} else {
			data_2 = "GANJIL"
		}

		if data_1 == nomorkeluaran {
			flag = true
		}
		if data_2 == nomorkeluaran {
			flag = true
		}

		if flag {
			result = "WINNER"
		}
	case "SHIO":
		flag := false
		kepala := string([]byte(temp)[2])
		ekor := string([]byte(temp)[3])
		data := _tableshio(kepala + ekor)

		if data == nomorkeluaran {
			flag = true
		}

		if flag {
			result = "WINNER"
		}
	}
	return result, win
}
func _tableshio(shiodata string) string {
	log.Printf("Data shio : %s", shiodata)

	tglnow, _ := goment.New()
	yearnow := tglnow.Format("YYYY")
	log.Println(yearnow)
	result := ""
	switch yearnow {
	case "2022":
		harimau := []string{"01", "13", "25", "37", "49", "61", "73", "85", "97"}
		kerbau := []string{"02", "14", "26", "38", "50", "62", "74", "86", "98"}
		tikus := []string{"03", "15", "27", "39", "51", "63", "75", "87", "99"}
		babi := []string{"04", "16", "28", "40", "52", "64", "76", "88", "00"}
		anjing := []string{"05", "17", "29", "41", "53", "65", "77", "89", ""}
		ayam := []string{"06", "18", "30", "42", "54", "66", "78", "90", ""}
		monyet := []string{"07", "19", "31", "43", "55", "67", "79", "91", ""}
		kambing := []string{"08", "20", "32", "44", "56", "68", "80", "92", ""}
		kuda := []string{"09", "21", "33", "45", "57", "69", "81", "93", ""}
		ular := []string{"10", "22", "34", "46", "58", "70", "82", "94", ""}
		naga := []string{"11", "23", "35", "47", "59", "71", "83", "95", ""}
		kelinci := []string{"12", "24", "36", "48", "60", "72", "84", "96", ""}
		for i := 0; i < len(babi); i++ {
			if shiodata == babi[i] {
				result = "BABI"
			}
		}
		for i := 0; i < len(ular); i++ {
			if shiodata == ular[i] {
				result = "ULAR"
			}
		}
		for i := 0; i < len(anjing); i++ {
			if shiodata == anjing[i] {
				result = "ANJING"
			}
		}
		for i := 0; i < len(ayam); i++ {
			if shiodata == ayam[i] {
				result = "AYAM"
			}
		}
		for i := 0; i < len(monyet); i++ {
			if shiodata == monyet[i] {
				result = "MONYET"
			}
		}
		for i := 0; i < len(kambing); i++ {
			if shiodata == kambing[i] {
				result = "KAMBING"
			}
		}
		for i := 0; i < len(kuda); i++ {
			if shiodata == kuda[i] {
				result = "KUDA"
			}
		}
		for i := 0; i < len(naga); i++ {
			if shiodata == naga[i] {
				result = "NAGA"
			}
		}
		for i := 0; i < len(kelinci); i++ {
			if shiodata == kelinci[i] {
				result = "KELINCI"
			}
		}
		for i := 0; i < len(harimau); i++ {
			if shiodata == harimau[i] {
				result = "HARIMAU"
			}
		}
		for i := 0; i < len(kerbau); i++ {
			if shiodata == kerbau[i] {
				result = "KERBAU"
			}
		}
		for i := 0; i < len(tikus); i++ {
			if shiodata == tikus[i] {
				result = "TIKUS"
			}
		}
	case "2023":
		kelinci := []string{"01", "13", "25", "37", "49", "61", "73", "85", "97"}
		harimau := []string{"02", "14", "26", "38", "50", "62", "74", "86", "98"}
		kerbau := []string{"03", "15", "27", "39", "51", "63", "75", "87", "99"}
		tikus := []string{"04", "16", "28", "40", "52", "64", "76", "88", "00"}
		babi := []string{"05", "17", "29", "41", "53", "65", "77", "89", ""}
		anjing := []string{"06", "18", "30", "42", "54", "66", "78", "90", ""}
		ayam := []string{"07", "19", "31", "43", "55", "67", "79", "91", ""}
		monyet := []string{"08", "20", "32", "44", "56", "68", "80", "92", ""}
		kambing := []string{"09", "21", "33", "45", "57", "69", "81", "93", ""}
		kuda := []string{"10", "22", "34", "46", "58", "70", "82", "94", ""}
		ular := []string{"11", "23", "35", "47", "59", "71", "83", "95", ""}
		naga := []string{"12", "24", "36", "48", "60", "72", "84", "96", ""}
		for i := 0; i < len(babi); i++ {
			if shiodata == babi[i] {
				result = "BABI"
			}
		}
		for i := 0; i < len(ular); i++ {
			if shiodata == ular[i] {
				result = "ULAR"
			}
		}
		for i := 0; i < len(anjing); i++ {
			if shiodata == anjing[i] {
				result = "ANJING"
			}
		}
		for i := 0; i < len(ayam); i++ {
			if shiodata == ayam[i] {
				result = "AYAM"
			}
		}
		for i := 0; i < len(monyet); i++ {
			if shiodata == monyet[i] {
				result = "MONYET"
			}
		}
		for i := 0; i < len(kambing); i++ {
			if shiodata == kambing[i] {
				result = "KAMBING"
			}
		}
		for i := 0; i < len(kuda); i++ {
			if shiodata == kuda[i] {
				result = "KUDA"
			}
		}
		for i := 0; i < len(naga); i++ {
			if shiodata == naga[i] {
				result = "NAGA"
			}
		}
		for i := 0; i < len(kelinci); i++ {
			if shiodata == kelinci[i] {
				result = "KELINCI"
			}
		}
		for i := 0; i < len(harimau); i++ {
			if shiodata == harimau[i] {
				result = "HARIMAU"
			}
		}
		for i := 0; i < len(kerbau); i++ {
			if shiodata == kerbau[i] {
				result = "KERBAU"
			}
		}
		for i := 0; i < len(tikus); i++ {
			if shiodata == tikus[i] {
				result = "TIKUS"
			}
		}
	case "2024":
		naga := []string{"01", "13", "25", "37", "49", "61", "73", "85", "97"}
		kelinci := []string{"02", "14", "26", "38", "50", "62", "74", "86", "98"}
		harimau := []string{"03", "15", "27", "39", "51", "63", "75", "87", "99"}
		kerbau := []string{"04", "16", "28", "40", "52", "64", "76", "88", "00"}
		tikus := []string{"05", "17", "29", "41", "53", "65", "77", "89", ""}
		babi := []string{"06", "18", "30", "42", "54", "66", "78", "90", ""}
		anjing := []string{"07", "19", "31", "43", "55", "67", "79", "91", ""}
		ayam := []string{"08", "20", "32", "44", "56", "68", "80", "92", ""}
		monyet := []string{"09", "21", "33", "45", "57", "69", "81", "93", ""}
		kambing := []string{"10", "22", "34", "46", "58", "70", "82", "94", ""}
		kuda := []string{"11", "23", "35", "47", "59", "71", "83", "95", ""}
		ular := []string{"12", "24", "36", "48", "60", "72", "84", "96", ""}
		for i := 0; i < len(babi); i++ {
			if shiodata == babi[i] {
				result = "BABI"
			}
		}
		for i := 0; i < len(ular); i++ {
			if shiodata == ular[i] {
				result = "ULAR"
			}
		}
		for i := 0; i < len(anjing); i++ {
			if shiodata == anjing[i] {
				result = "ANJING"
			}
		}
		for i := 0; i < len(ayam); i++ {
			if shiodata == ayam[i] {
				result = "AYAM"
			}
		}
		for i := 0; i < len(monyet); i++ {
			if shiodata == monyet[i] {
				result = "MONYET"
			}
		}
		for i := 0; i < len(kambing); i++ {
			if shiodata == kambing[i] {
				result = "KAMBING"
			}
		}
		for i := 0; i < len(kuda); i++ {
			if shiodata == kuda[i] {
				result = "KUDA"
			}
		}
		for i := 0; i < len(naga); i++ {
			if shiodata == naga[i] {
				result = "NAGA"
			}
		}
		for i := 0; i < len(kelinci); i++ {
			if shiodata == kelinci[i] {
				result = "KELINCI"
			}
		}
		for i := 0; i < len(harimau); i++ {
			if shiodata == harimau[i] {
				result = "HARIMAU"
			}
		}
		for i := 0; i < len(kerbau); i++ {
			if shiodata == kerbau[i] {
				result = "KERBAU"
			}
		}
		for i := 0; i < len(tikus); i++ {
			if shiodata == tikus[i] {
				result = "TIKUS"
			}
		}
	}

	return result
}
func _pasaran_id(idcomppasaran int, company, tipecolumn string) (string, float32) {
	con := db.CreateCon()
	ctx := context.Background()
	var result string = ""
	var result_number float32 = 0
	sql_pasaran := `SELECT 
		idpasarantogel , 
		2_win as win_cbebas, 3_win2digit as win2_cmacau, 
		3_win3digit as win3_cmacau, 3_win4digit as win4_cmacau, 
		4_win3digit as win3_cnaga, 4_win4digit as win4_cnaga 
		FROM ` + config.DB_tbl_mst_company_game_pasaran + `  
		WHERE idcomppasaran  = ? 
		AND idcompany = ? 
	`
	var (
		idpasarantogel_db                                             string
		win_cbebas_db, win2_cmacau_db, win3_cmacau_db, win4_cmacau_db float32
		win3_cnaga_db, win4_cnaga_db                                  float32
	)
	rows := con.QueryRowContext(ctx, sql_pasaran, idcomppasaran, company)
	switch err := rows.Scan(
		&idpasarantogel_db,
		&win_cbebas_db, &win2_cmacau_db, &win3_cmacau_db, &win4_cmacau_db,
		&win3_cnaga_db, &win4_cnaga_db); err {
	case sql.ErrNoRows:
		result = ""
	case nil:
		switch tipecolumn {
		case "idpasarantogel":
			result = idpasarantogel_db
		case "2_win":
			result_number = win_cbebas_db
		case "3_win2digit":
			result_number = win2_cmacau_db
		case "3_win3digit":
			result_number = win3_cmacau_db
		case "3_win4digit":
			result_number = win4_cmacau_db
		case "4_win3digit":
			result_number = win3_cnaga_db
		}
	default:
		helpers.ErrorCheck(err)
	}
	return result, result_number
}
