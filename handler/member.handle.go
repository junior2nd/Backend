package handler

import (
	"msa/database"
	"msa/model/entity"
	"msa/model/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func MemberGetAll(c *fiber.Ctx)error{
	var member []entity.Member
	result := database.DB.Debug().Find(&member)
	if result.Error != nil{
		return fiber.NewError(fiber.StatusNotFound,"Not Member Data")
	}
	return c.JSON(member)
}

func MemberCreate(c *fiber.Ctx)error{
	member := new(request.MemberCreateRequest)
	err := c.BodyParser(member)
	if err != nil{
		return fiber.NewError(fiber.StatusNotFound,"No, Input Data")
	}

	// Validate Data
	validate := validator.New()
	errValidate := validate.Struct(member)
	if errValidate != nil{
		return fiber.NewError(fiber.StatusNotFound,"Validate")
	}

	newMember := entity.Member{
		Name: member.Name,
		Adress: member.Address,
		Phone: member.Phone,
		Image: member.Image,
	}

	MemberCreate := database.DB.Create(&newMember)
	if MemberCreate.Error != nil {
		return fiber.NewError(fiber.StatusNotFound,"Not Create Member")
	}

	return c.JSON(member)
}

func MemberUpdate(c *fiber.Ctx)error{
	return nil
}

func MemberDelete(c *fiber.Ctx)error{
	return nil
}