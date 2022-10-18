package route

import (
	"msa/config"
	"msa/handler"
	// "msa/handler"
	// "msa/middleware"
	// "msa/utils"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {

	r.Static("/public", config.ProjectRootPath+"/public/asset")

	// UserHandle
	r.Get("/user",handler.UserGetAll)
	r.Post("/user",handler.UserCreate)

	// MemberHandle
	r.Get("/member",handler.MemberGetAll)
	r.Post("/member",handler.MemberCreate)

	// RoomTypeHandle
	r.Get("/roomtype",handler.RoomTypeGetAll)
	r.Post("/roomtype",handler.RoomTypeCreate)

	// RoomHandle
	r.Get("/room",handler.RoomGetAll)
	r.Post("/room",handler.RoomCreate)

	// TransactionHandle
	r.Get("/transaction",handler.TransactionGetAll)
	r.Post("/transaction",handler.TransactionCreate)

	// ApartmentHandle
	r.Get("/apartment",handler.ApartmentGetAll)
	r.Post("/apartment",handler.ApartmentCreate)


}
