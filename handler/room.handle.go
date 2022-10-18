package handler

import (
	"msa/database"
	"msa/model/entity"
	"msa/model/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func RoomGetAll(c *fiber.Ctx)error{
	var room []entity.Room
	result := database.DB.Debug().Find(&room)
	if result.Error != nil{
		return fiber.NewError(fiber.StatusNotFound,"Not RoomData")
	}
	return c.JSON(room)
}

func RoomCreate(c *fiber.Ctx)error{
	room := new(request.RoomCreateRequest)
	errRoom := c.BodyParser(room)
	if errRoom != nil{
		return fiber.NewError(fiber.StatusNotFound,"Not Input Data")
	}

	validate := validator.New()
	errValidate := validate.Struct(room)
	if errValidate != nil{
		return fiber.NewError(fiber.StatusNotFound,"Data Not True")
	}

	newRoom := entity.Room{
		Name: room.Name,
		Status: room.Status,
	}

	createRoom := database.DB.Create(&newRoom)
	if createRoom.Error != nil{
		return fiber.NewError(fiber.StatusNotFound,"Not Create Room")
	}
	return c.JSON(room)
}

func RoomUpdate(c *fiber.Ctx)error{
	return nil
}

func RoomDelete(c *fiber.Ctx)error{
	return nil
}