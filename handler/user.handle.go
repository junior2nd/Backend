package handler

import (
	"log"
	"msa/database"
	"msa/model/entity"
	"msa/model/request"
	"msa/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserGetAll(c *fiber.Ctx)error{
	var user []entity.User
	result := database.DB.Debug().Find(&user)
	if result.Error != nil {
		log.Print(result.Error)
	}
	return c.JSON(user)
}

func UserCreate(c *fiber.Ctx)error{
	user := new(request.UserCreateRequest)
	err := c.BodyParser(user)
	if err != nil{
		return err
	}
	// Validate Data
	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil{
		return fiber.NewError(fiber.StatusNotFound,"Not Data")
	}

	newUser := entity.User{
		Name: user.Name,
		Email: user.Email,
		Adress: user.Address,
		Phone: user.Phone,
		Image: user.Image,
	}

	hashPassword,err := utils.HashingPassword(user.Password)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusNotFound,"server error")
	}

	newUser.Password = hashPassword

	errCreate := database.DB.Create(&newUser).Error

	if errCreate != nil{
		return fiber.NewError(fiber.StatusNotFound,"Not CrateUser")
	}
	return c.Status(fiber.StatusCreated).JSON(newUser)
}

func UserUpdate(c *fiber.Ctx)error{
	return nil
}

func UserDelete(c *fiber.Ctx)error{
	return nil
}