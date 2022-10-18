package handler

import (
	"log"
	"msa/database"
	"msa/model/entity"
	"msa/model/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func TransactionGetAll(c *fiber.Ctx)error{
	var transaction []entity.Transaction
	result := database.DB.Debug().Find(&transaction)
	if result.Error != nil{
		return fiber.NewError(fiber.StatusNotFound,"No Transaction Data")
	}
	return c.JSON(transaction)
}

func TransactionCreate(c *fiber.Ctx)error{
	transaction := new(request.TransactionCreateRequest)
	log.Print(transaction)
	errTransaction := c.BodyParser(transaction)
	
	if errTransaction != nil{
		return fiber.NewError(fiber.StatusNotFound,"TransactionData not Input")
	}

	validate := validator.New()
	errValidate := validate.Struct(transaction)
	if errValidate != nil{
		return fiber.NewError(fiber.StatusNotFound,"Data Transacion Not True")
	}

	newTransaction := entity.Transaction{
		Status: transaction.Status,
		Evalue: transaction.Evalue,
		Wvalue: transaction.Wvalue,
	}

	createTransaction := database.DB.Create(&newTransaction)
	if createTransaction.Error != nil{
		return fiber.NewError(fiber.StatusNotFound,"Not Create Transaction")
	}
	return c.JSON(transaction)
}

func TransactionUpdate(c *fiber.Ctx)error{
	return nil
}

func TransactionDelete(c *fiber.Ctx)error{
	return nil
}