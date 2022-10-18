package handler

import (
	"msa/database"
	"msa/model/entity"
	"msa/model/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ApartmentGetAll(c *fiber.Ctx)error{
	var apartment []entity.Apartment
	result := database.DB.Debug().Find(&apartment)
	if result.Error != nil{
		return fiber.NewError(fiber.StatusNotFound,"Apartment Data NotFound")
	}
	return c.JSON(apartment)
}

func ApartmentCreate(c *fiber.Ctx)error{
	apartment := new(request.ApartmentCreateRequest)
	errApartment := c.BodyParser(apartment)
	if errApartment != nil{
		return fiber.NewError(fiber.StatusNotFound,"No input Apartment Data")
	}

	validate := validator.New()
	errValidate := validate.Struct(apartment)
	if errValidate != nil{
		return fiber.NewError(fiber.StatusNotFound,"Data Not True")
	}

	newApartment := entity.Apartment{
		Name: apartment.Name,
		Branch: apartment.Branch,
		Phone: apartment.Phone,
		Adress: apartment.Adress,
		Image: apartment.Image,
	}

	createApartment := database.DB.Create(&newApartment) 
	if createApartment.Error != nil{
		return fiber.NewError(fiber.StatusNotFound,"Not Create Apartment")
	}
	return c.JSON(apartment)
}

func ApartmentUpdate(c *fiber.Ctx)error{
	return nil
}

func ApartmentDelete(c *fiber.Ctx)error{
	return nil
}