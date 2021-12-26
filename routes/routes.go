package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/controller"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/middleware"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())

	app.Post("/api/login", controller.CheckLogin)
	app.Post("/api/loginother", controller.CheckLoginOtherWebsite)

	app.Post("/api/home", middleware.JWTProtected(), controller.Home)
	app.Post("/api/allinvoice", middleware.JWTProtected(), controller.InvoiceHome)
	app.Post("/api/saveinvoice", middleware.JWTProtected(), controller.InvoiceSave)
	app.Post("/api/saveinvoicewinlosestatus", middleware.JWTProtected(), controller.InvoiceSavewinlosestatus)

	app.Post("/api/allcompany", middleware.JWTProtected(), controller.CompanyHome)
	app.Post("/api/detailcompany", middleware.JWTProtected(), controller.CompanyDetail)
	app.Post("/api/companylistadmin", middleware.JWTProtected(), controller.CompanyDetailListAdmin)
	app.Post("/api/companylistpasaran", middleware.JWTProtected(), controller.CompanyDetailListPasaran)
	app.Post("/api/companylistpasaranonline", middleware.JWTProtected(), controller.CompanyDetailListPasaranOnline)
	app.Post("/api/companylistpasaranconf", middleware.JWTProtected(), controller.CompanyDetailListPasaranConf)
	app.Post("/api/companylistkeluaran", middleware.JWTProtected(), controller.CompanyListKeluaran)
	app.Post("/api/companyinvoicemember", middleware.JWTProtected(), controller.CompanyInvoiceMember)
	app.Post("/api/companyinvoicemembertemp", middleware.JWTProtected(), controller.CompanyInvoiceMemberTemp)
	app.Post("/api/companyinvoicemembersync", middleware.JWTProtected(), controller.CompanyInvoiceMemberSync)
	app.Post("/api/companyinvoicegrouppermainan", middleware.JWTProtected(), controller.CompanyInvoiceGroupPermainan)
	app.Post("/api/companyinvoicelistpermainan", middleware.JWTProtected(), controller.CompanyInvoicelistpermainan)
	app.Post("/api/companyinvoicelistpermainanstatus", middleware.JWTProtected(), controller.CompanyInvoicelistpermainanbystatus)
	app.Post("/api/companyinvoicelistpermainanmember", middleware.JWTProtected(), controller.CompanyInvoicelistpermainanbyusername)
	app.Post("/api/savecompanypasaranonline", middleware.JWTProtected(), controller.CompanySaveNewPasaranHariOnline)
	app.Post("/api/deletecompanypasaranonline", middleware.JWTProtected(), controller.CompanyDeletePasaranHariOnline)
	app.Post("/api/savecompany", middleware.JWTProtected(), controller.CompanySave)
	app.Post("/api/savecompanyadmin", middleware.JWTProtected(), controller.CompanySaveNewAdmin)
	app.Post("/api/savecompanypasaran", middleware.JWTProtected(), controller.CompanySaveNewPasaran)
	app.Post("/api/savecompanyfetchpasaranlimitline", middleware.JWTProtected(), controller.CompanyFetchPasaranlimitline)
	app.Post("/api/savecompanyfetchpasaran432", middleware.JWTProtected(), controller.CompanyFetchPasaran432)
	app.Post("/api/savecompanyfetchpasarancolokbebas", middleware.JWTProtected(), controller.CompanyFetchPasarancolokbebas)
	app.Post("/api/savecompanyfetchpasarancolokmacau", middleware.JWTProtected(), controller.CompanyFetchPasarancolokmacau)
	app.Post("/api/savecompanyfetchpasarancoloknaga", middleware.JWTProtected(), controller.CompanyFetchPasarancoloknaga)
	app.Post("/api/savecompanyfetchpasarancolokjitu", middleware.JWTProtected(), controller.CompanyFetchPasarancolokjitu)
	app.Post("/api/savecompanyfetchpasaran5050umum", middleware.JWTProtected(), controller.CompanyFetchPasaran5050umum)
	app.Post("/api/savecompanyfetchpasaran5050special", middleware.JWTProtected(), controller.CompanyFetchPasaran5050special)
	app.Post("/api/savecompanyfetchpasaran5050kombinasi", middleware.JWTProtected(), controller.CompanyFetchPasaran5050kombinasi)
	app.Post("/api/savecompanyfetchpasaranmacaukombinasi", middleware.JWTProtected(), controller.CompanyFetchPasaranmacaukombinasi)
	app.Post("/api/savecompanyfetchpasarandasar", middleware.JWTProtected(), controller.CompanyFetchPasarandasar)
	app.Post("/api/savecompanyfetchpasaranshio", middleware.JWTProtected(), controller.CompanyFetchPasaranshio)
	app.Post("/api/savecompanyupdatepasaran", middleware.JWTProtected(), controller.CompanyPasaranUpdate)
	app.Post("/api/savecompanyupdatepasaranline", middleware.JWTProtected(), controller.CompanyPasaranUpdateLimitline)
	app.Post("/api/savecompanyupdatepasaran432", middleware.JWTProtected(), controller.CompanyPasaranUpdate432)
	app.Post("/api/savecompanyupdatepasarancolokbebas", middleware.JWTProtected(), controller.CompanyPasaranUpdatecolokbebas)
	app.Post("/api/savecompanyupdatepasarancolokmacau", middleware.JWTProtected(), controller.CompanyPasaranUpdatecolokmacau)
	app.Post("/api/savecompanyupdatepasarancoloknaga", middleware.JWTProtected(), controller.CompanyPasaranUpdatecoloknaga)
	app.Post("/api/savecompanyupdatepasarancolokjitu", middleware.JWTProtected(), controller.CompanyPasaranUpdatecolokjitu)
	app.Post("/api/savecompanyupdatepasaran5050umum", middleware.JWTProtected(), controller.CompanyPasaranUpdate5050umum)
	app.Post("/api/savecompanyupdatepasaran5050special", middleware.JWTProtected(), controller.CompanyPasaranUpdate5050special)
	app.Post("/api/savecompanyupdatepasaran5050kombinasi", middleware.JWTProtected(), controller.CompanyPasaranUpdate5050kombinasi)
	app.Post("/api/savecompanyupdatepasarankombinasi", middleware.JWTProtected(), controller.CompanyPasaranUpdatekombinasi)
	app.Post("/api/savecompanyupdatepasarandasar", middleware.JWTProtected(), controller.CompanyPasaranUpdatedasar)
	app.Post("/api/savecompanyupdatepasaranshio", middleware.JWTProtected(), controller.CompanyPasaranUpdateshio)

	app.Post("/api/allpasaran", middleware.JWTProtected(), controller.PasaranHome)
	app.Post("/api/detailpasaran", middleware.JWTProtected(), controller.PasaranDetail)
	app.Post("/api/detailconfpasaran", middleware.JWTProtected(), controller.PasaranDetailConf)
	app.Post("/api/savepasaran", middleware.JWTProtected(), controller.PasaranSave)
	app.Post("/api/savepasaranlimitline", middleware.JWTProtected(), controller.PasaranSaveLimitLine)
	app.Post("/api/savepasaranconf432d", middleware.JWTProtected(), controller.PasaranSaveConf432d)
	app.Post("/api/savepasaranconfcolokbebas", middleware.JWTProtected(), controller.PasaranSaveConfColokBebas)
	app.Post("/api/savepasaranconfcolokmacau", middleware.JWTProtected(), controller.PasaranSaveConfColokMacau)
	app.Post("/api/savepasaranconfcoloknaga", middleware.JWTProtected(), controller.PasaranSaveConfColokNaga)
	app.Post("/api/savepasaranconfcolokjitu", middleware.JWTProtected(), controller.PasaranSaveConfColokJitu)
	app.Post("/api/savepasaranconf5050umum", middleware.JWTProtected(), controller.PasaranSaveConf5050umum)
	app.Post("/api/savepasaranconf5050special", middleware.JWTProtected(), controller.PasaranSaveConf5050special)
	app.Post("/api/savepasaranconf5050kombinasi", middleware.JWTProtected(), controller.PasaranSaveConf5050kombinasi)
	app.Post("/api/savepasaranconfmacaukombinasi", middleware.JWTProtected(), controller.PasaranSaveConfmacaukombinasi)
	app.Post("/api/savepasaranconfdasar", middleware.JWTProtected(), controller.PasaranSaveConfdasar)
	app.Post("/api/savepasaranconfshio", middleware.JWTProtected(), controller.PasaranSaveConfshio)

	app.Post("/api/setting", middleware.JWTProtected(), controller.SettingHome)
	app.Post("/api/savesetting", middleware.JWTProtected(), controller.SettingSave)

	app.Post("/api/domain", middleware.JWTProtected(), controller.Domainhome)
	app.Post("/api/savedomain", middleware.JWTProtected(), controller.DomainSave)

	app.Post("/api/listpasaranwajib", middleware.JWTProtected(), controller.Listpasaranwajib)
	app.Post("/api/prediksiwajib", middleware.JWTProtected(), controller.Prediksiwajib)
	return app
}
