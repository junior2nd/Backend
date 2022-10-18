package handler

import (
	"log"
	"msa/database"
	"msa/model/entity"
	"msa/model/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func RoomTypeGetAll(c *fiber.Ctx) error {
	var roomtype []entity.Roomtype
	result := database.DB.Debug().Find(&roomtype)
	if result.Error != nil{
		return fiber.NewError(fiber.StatusNotFound,"No Data Roomtype")
	}
	return c.JSON(roomtype)
}

func RoomTypeCreate(c *fiber.Ctx) error {
	roomtype := new(request.RoomtypeCreateRequest)
	log.Print(roomtype)
	err := c.BodyParser(roomtype)
	if err != nil{
		return fiber.NewError(fiber.StatusNotFound,"Not input Data (price :: int ??)")
	}

	// Valedate
	validate := validator.New()
	errVaridate := validate.Struct(roomtype)
	if errVaridate != nil{
		return fiber.NewError(fiber.StatusNotFound,"Error Validate")
	}

	newRoomtype := entity.Roomtype{
		Name: roomtype.Name,
		Detail: roomtype.Detail,
		Price: roomtype.Price,
	}

	createRoomtype := database.DB.Create(&newRoomtype)
	if createRoomtype.Error != nil{
		return fiber.NewError(fiber.StatusNotFound,"Not Create Roomtype")
	}

	return c.JSON(roomtype)
}

func RoomTypeUpdate(c *fiber.Ctx) error {
	return nil
}

func RoomTypeDelete(c *fiber.Ctx) error {
	return nil
}
