package url

import (
	"github.com/gofiber/swagger"
	"github.com/indrariksa/ws-indra2024/controller"

	"github.com/gofiber/fiber/v2"
)

func Web(page *fiber.App) {
	// page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	// page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth

	page.Get("/", controller.Sink)
	page.Post("/", controller.Sink)
	page.Put("/", controller.Sink)
	page.Patch("/", controller.Sink)
	page.Delete("/", controller.Sink)
	page.Options("/", controller.Sink)

	page.Get("/checkip", controller.Homepage)
	page.Get("/presensi", controller.GetPresensi)
	page.Get("/presensi/:id", controller.GetPresensiID)
	page.Post("/insert", controller.InsertDataPresensi)
	page.Put("/update/:id", controller.UpdateData)
	page.Delete("/delete/:id", controller.DeletePresensiByID)

	//swagger
	page.Get("/docs/*", swagger.HandlerDefault)
}
