package route

import (
	"go-fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func Route_init(app *fiber.App) {
	app.Get("/user", handler.UserHandler)
	app.Get("/user/:id", handler.GetDataById)
	app.Post("/user", handler.UserCreate)
	app.Put("/user/:id/UpdateDataAll", handler.UpdateData)
	app.Put("/user/:id/UpdateEmail", handler.UpdateEmailData)
	app.Delete("/user/:id/DeleteUser", handler.UserDelete)
}
