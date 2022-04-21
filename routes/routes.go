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
	api := app.Group("/api/", middleware.JWTProtected())

	api.Post("home", middleware.JWTProtected(), controller.Home)
	api.Post("allinvoice", middleware.JWTProtected(), controller.InvoiceHome)
	api.Post("saveinvoice", middleware.JWTProtected(), controller.InvoiceSave)
	api.Post("saveinvoicewinlosestatus", middleware.JWTProtected(), controller.InvoiceSavewinlosestatus)

	api.Post("allcompany", middleware.JWTProtected(), controller.CompanyHome)
	api.Post("detailcompany", middleware.JWTProtected(), controller.CompanyDetail)
	api.Post("companylistadmin", middleware.JWTProtected(), controller.CompanyDetailListAdmin)
	api.Post("companylistpasaran", middleware.JWTProtected(), controller.CompanyDetailListPasaran)
	api.Post("companylistpasaranonline", middleware.JWTProtected(), controller.CompanyDetailListPasaranOnline)
	api.Post("companylistpasaranconf", middleware.JWTProtected(), controller.CompanyDetailListPasaranConf)
	api.Post("companylistkeluaran", middleware.JWTProtected(), controller.CompanyListKeluaran)
	api.Post("companyinvoicemember", middleware.JWTProtected(), controller.CompanyInvoiceMember)
	api.Post("companyinvoicemembertemp", middleware.JWTProtected(), controller.CompanyInvoiceMemberTemp)
	api.Post("companyinvoicemembersync", middleware.JWTProtected(), controller.CompanyInvoiceMemberSync)
	api.Post("companyinvoicegrouppermainan", middleware.JWTProtected(), controller.CompanyInvoiceGroupPermainan)
	api.Post("companyinvoicelistpermainan", middleware.JWTProtected(), controller.CompanyInvoicelistpermainan)
	api.Post("companyinvoicelistpermainanstatus", middleware.JWTProtected(), controller.CompanyInvoicelistpermainanbystatus)
	api.Post("companyinvoicelistpermainanmember", middleware.JWTProtected(), controller.CompanyInvoicelistpermainanbyusername)
	api.Post("savecompanypasaranonline", middleware.JWTProtected(), controller.CompanySaveNewPasaranHariOnline)
	api.Post("deletecompanypasaranonline", middleware.JWTProtected(), controller.CompanyDeletePasaranHariOnline)
	api.Post("savecompany", middleware.JWTProtected(), controller.CompanySave)
	api.Post("savecompanyadmin", middleware.JWTProtected(), controller.CompanySaveNewAdmin)
	api.Post("savecompanypasaran", middleware.JWTProtected(), controller.CompanySaveNewPasaran)
	api.Post("savecompanyfetchpasaranlimitline", middleware.JWTProtected(), controller.CompanyFetchPasaranlimitline)
	api.Post("savecompanyfetchpasaran432", middleware.JWTProtected(), controller.CompanyFetchPasaran432)
	api.Post("savecompanyfetchpasarancolokbebas", middleware.JWTProtected(), controller.CompanyFetchPasarancolokbebas)
	api.Post("savecompanyfetchpasarancolokmacau", middleware.JWTProtected(), controller.CompanyFetchPasarancolokmacau)
	api.Post("savecompanyfetchpasarancoloknaga", middleware.JWTProtected(), controller.CompanyFetchPasarancoloknaga)
	api.Post("savecompanyfetchpasarancolokjitu", middleware.JWTProtected(), controller.CompanyFetchPasarancolokjitu)
	api.Post("savecompanyfetchpasaran5050umum", middleware.JWTProtected(), controller.CompanyFetchPasaran5050umum)
	api.Post("savecompanyfetchpasaran5050special", middleware.JWTProtected(), controller.CompanyFetchPasaran5050special)
	api.Post("savecompanyfetchpasaran5050kombinasi", middleware.JWTProtected(), controller.CompanyFetchPasaran5050kombinasi)
	api.Post("savecompanyfetchpasaranmacaukombinasi", middleware.JWTProtected(), controller.CompanyFetchPasaranmacaukombinasi)
	api.Post("savecompanyfetchpasarandasar", middleware.JWTProtected(), controller.CompanyFetchPasarandasar)
	api.Post("savecompanyfetchpasaranshio", middleware.JWTProtected(), controller.CompanyFetchPasaranshio)
	api.Post("savecompanyupdatepasaran", middleware.JWTProtected(), controller.CompanyPasaranUpdate)
	api.Post("savecompanyupdatepasaranline", middleware.JWTProtected(), controller.CompanyPasaranUpdateLimitline)
	api.Post("savecompanyupdatepasaran432", middleware.JWTProtected(), controller.CompanyPasaranUpdate432)
	api.Post("savecompanyupdatepasarancolokbebas", middleware.JWTProtected(), controller.CompanyPasaranUpdatecolokbebas)
	api.Post("savecompanyupdatepasarancolokmacau", middleware.JWTProtected(), controller.CompanyPasaranUpdatecolokmacau)
	api.Post("savecompanyupdatepasarancoloknaga", middleware.JWTProtected(), controller.CompanyPasaranUpdatecoloknaga)
	api.Post("savecompanyupdatepasarancolokjitu", middleware.JWTProtected(), controller.CompanyPasaranUpdatecolokjitu)
	api.Post("savecompanyupdatepasaran5050umum", middleware.JWTProtected(), controller.CompanyPasaranUpdate5050umum)
	api.Post("savecompanyupdatepasaran5050special", middleware.JWTProtected(), controller.CompanyPasaranUpdate5050special)
	api.Post("savecompanyupdatepasaran5050kombinasi", middleware.JWTProtected(), controller.CompanyPasaranUpdate5050kombinasi)
	api.Post("savecompanyupdatepasarankombinasi", middleware.JWTProtected(), controller.CompanyPasaranUpdatekombinasi)
	api.Post("savecompanyupdatepasarandasar", middleware.JWTProtected(), controller.CompanyPasaranUpdatedasar)
	api.Post("savecompanyupdatepasaranshio", middleware.JWTProtected(), controller.CompanyPasaranUpdateshio)

	api.Post("allpasaran", middleware.JWTProtected(), controller.PasaranHome)
	api.Post("detailpasaran", middleware.JWTProtected(), controller.PasaranDetail)
	api.Post("detailconfpasaran", middleware.JWTProtected(), controller.PasaranDetailConf)
	api.Post("savepasaran", middleware.JWTProtected(), controller.PasaranSave)
	api.Post("savepasaranlimitline", middleware.JWTProtected(), controller.PasaranSaveLimitLine)
	api.Post("savepasaranconf432d", middleware.JWTProtected(), controller.PasaranSaveConf432d)
	api.Post("savepasaranconfcolokbebas", middleware.JWTProtected(), controller.PasaranSaveConfColokBebas)
	api.Post("savepasaranconfcolokmacau", middleware.JWTProtected(), controller.PasaranSaveConfColokMacau)
	api.Post("savepasaranconfcoloknaga", middleware.JWTProtected(), controller.PasaranSaveConfColokNaga)
	api.Post("savepasaranconfcolokjitu", middleware.JWTProtected(), controller.PasaranSaveConfColokJitu)
	api.Post("savepasaranconf5050umum", middleware.JWTProtected(), controller.PasaranSaveConf5050umum)
	api.Post("savepasaranconf5050special", middleware.JWTProtected(), controller.PasaranSaveConf5050special)
	api.Post("savepasaranconf5050kombinasi", middleware.JWTProtected(), controller.PasaranSaveConf5050kombinasi)
	api.Post("savepasaranconfmacaukombinasi", middleware.JWTProtected(), controller.PasaranSaveConfmacaukombinasi)
	api.Post("savepasaranconfdasar", middleware.JWTProtected(), controller.PasaranSaveConfdasar)
	api.Post("savepasaranconfshio", middleware.JWTProtected(), controller.PasaranSaveConfshio)

	api.Post("setting", middleware.JWTProtected(), controller.SettingHome)
	api.Post("savesetting", middleware.JWTProtected(), controller.SettingSave)

	api.Post("domain", middleware.JWTProtected(), controller.Domainhome)
	api.Post("savedomain", middleware.JWTProtected(), controller.DomainSave)

	api.Post("listpasaranwajib", middleware.JWTProtected(), controller.Listpasaranwajib)
	api.Post("prediksiwajib", middleware.JWTProtected(), controller.Prediksiwajib)
	return app
}
