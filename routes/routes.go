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

	api.Post("home", controller.Home)
	api.Post("allinvoice", controller.InvoiceHome)
	api.Post("invoicedetail", controller.InvoiceDetail)
	api.Post("saveinvoice", controller.InvoiceSave)
	api.Post("saveinvoicewinlosestatus", controller.InvoiceSavewinlosestatus)
	api.Post("saveinvoicepasaran", controller.InvoiceSavePasaran)

	api.Post("allcompany", controller.CompanyHome)
	api.Post("detailcompany", controller.CompanyDetail)
	api.Post("companylistadmin", controller.CompanyDetailListAdmin)
	api.Post("companylistpasaran", controller.CompanyDetailListPasaran)
	api.Post("companylistpasaranonline", controller.CompanyDetailListPasaranOnline)
	api.Post("companylistpasaranconf", controller.CompanyDetailListPasaranConf)
	api.Post("companylistkeluaran", controller.CompanyListKeluaran)
	api.Post("companyinvoicemember", controller.CompanyInvoiceMember)
	api.Post("companyinvoicemembertemp", controller.CompanyInvoiceMemberTemp)
	api.Post("companyinvoicemembersync", controller.CompanyInvoiceMemberSync)
	api.Post("companyinvoicegrouppermainan", controller.CompanyInvoiceGroupPermainan)
	api.Post("companyinvoicelistpermainan", controller.CompanyInvoicelistpermainan)
	api.Post("companyinvoicelistpermainanstatus", controller.CompanyInvoicelistpermainanbystatus)
	api.Post("companyinvoicelistpermainanmember", controller.CompanyInvoicelistpermainanbyusername)
	api.Post("savecompanypasaranonline", controller.CompanySaveNewPasaranHariOnline)
	api.Post("deletecompanypasaranonline", controller.CompanyDeletePasaranHariOnline)
	api.Post("savecompany", controller.CompanySave)
	api.Post("savecompanyadmin", controller.CompanySaveNewAdmin)
	api.Post("savecompanypasaran", controller.CompanySaveNewPasaran)
	api.Post("savecompanyfetchpasaranlimitline", controller.CompanyFetchPasaranlimitline)
	api.Post("savecompanyfetchpasaran432", controller.CompanyFetchPasaran432)
	api.Post("savecompanyfetchpasarancolokbebas", controller.CompanyFetchPasarancolokbebas)
	api.Post("savecompanyfetchpasarancolokmacau", controller.CompanyFetchPasarancolokmacau)
	api.Post("savecompanyfetchpasarancoloknaga", controller.CompanyFetchPasarancoloknaga)
	api.Post("savecompanyfetchpasarancolokjitu", controller.CompanyFetchPasarancolokjitu)
	api.Post("savecompanyfetchpasaran5050umum", controller.CompanyFetchPasaran5050umum)
	api.Post("savecompanyfetchpasaran5050special", controller.CompanyFetchPasaran5050special)
	api.Post("savecompanyfetchpasaran5050kombinasi", controller.CompanyFetchPasaran5050kombinasi)
	api.Post("savecompanyfetchpasaranmacaukombinasi", controller.CompanyFetchPasaranmacaukombinasi)
	api.Post("savecompanyfetchpasarandasar", controller.CompanyFetchPasarandasar)
	api.Post("savecompanyfetchpasaranshio", controller.CompanyFetchPasaranshio)
	api.Post("savecompanyupdatepasaran", controller.CompanyPasaranUpdate)
	api.Post("savecompanyupdatepasaranline", controller.CompanyPasaranUpdateLimitline)
	api.Post("savecompanyupdatepasaran432", controller.CompanyPasaranUpdate432)
	api.Post("savecompanyupdatepasarancolokbebas", controller.CompanyPasaranUpdatecolokbebas)
	api.Post("savecompanyupdatepasarancolokmacau", controller.CompanyPasaranUpdatecolokmacau)
	api.Post("savecompanyupdatepasarancoloknaga", controller.CompanyPasaranUpdatecoloknaga)
	api.Post("savecompanyupdatepasarancolokjitu", controller.CompanyPasaranUpdatecolokjitu)
	api.Post("savecompanyupdatepasaran5050umum", controller.CompanyPasaranUpdate5050umum)
	api.Post("savecompanyupdatepasaran5050special", controller.CompanyPasaranUpdate5050special)
	api.Post("savecompanyupdatepasaran5050kombinasi", controller.CompanyPasaranUpdate5050kombinasi)
	api.Post("savecompanyupdatepasarankombinasi", controller.CompanyPasaranUpdatekombinasi)
	api.Post("savecompanyupdatepasarandasar", controller.CompanyPasaranUpdatedasar)
	api.Post("savecompanyupdatepasaranshio", controller.CompanyPasaranUpdateshio)

	api.Post("allpasaran", controller.PasaranHome)
	api.Post("detailpasaran", controller.PasaranDetail)
	api.Post("detailconfpasaran", controller.PasaranDetailConf)
	api.Post("savepasaran", controller.PasaranSave)
	api.Post("savepasaranlimitline", controller.PasaranSaveLimitLine)
	api.Post("savepasaranconf432d", controller.PasaranSaveConf432d)
	api.Post("savepasaranconfcolokbebas", controller.PasaranSaveConfColokBebas)
	api.Post("savepasaranconfcolokmacau", controller.PasaranSaveConfColokMacau)
	api.Post("savepasaranconfcoloknaga", controller.PasaranSaveConfColokNaga)
	api.Post("savepasaranconfcolokjitu", controller.PasaranSaveConfColokJitu)
	api.Post("savepasaranconf5050umum", controller.PasaranSaveConf5050umum)
	api.Post("savepasaranconf5050special", controller.PasaranSaveConf5050special)
	api.Post("savepasaranconf5050kombinasi", controller.PasaranSaveConf5050kombinasi)
	api.Post("savepasaranconfmacaukombinasi", controller.PasaranSaveConfmacaukombinasi)
	api.Post("savepasaranconfdasar", controller.PasaranSaveConfdasar)
	api.Post("savepasaranconfshio", controller.PasaranSaveConfshio)

	api.Post("setting", controller.SettingHome)
	api.Post("savesetting", controller.SettingSave)

	api.Post("domain", controller.Domainhome)
	api.Post("savedomain", controller.DomainSave)

	api.Post("listpasaranwajib", controller.Listpasaranwajib)
	api.Post("prediksiwajib", controller.Prediksiwajib)
	return app
}
