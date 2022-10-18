package utils

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

const DefaultPathAssetImage = "./public/covers/"

func HadleSingleFile(c *fiber.Ctx) error {
	//Handle file
	file, errFile := c.FormFile("cover")
	if errFile != nil {
		log.Println("error file = ", errFile)
	}

	var filename *string
	if file != nil {
		errCheckComtentType := checkContentFile(file, "image/jpeg", "image/png", "image/gif")
		if errCheckComtentType != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"msg": errCheckComtentType.Error(),
			})
		}
		filename = &file.Filename
		// type file
		extenstionFile := filepath.Ext(*filename)
		// Rename File Upload
		newFilename := fmt.Sprintf("image-%s", extenstionFile)

		errSaveFile := c.SaveFile(file, fmt.Sprintf("./public/covers/%s", newFilename))
		if errSaveFile != nil {
			log.Println("Upload Fail")
		}
	} else {
		log.Println("Nothing file to upload")
	}

	if filename != nil {
		c.Locals("filename", *filename)
	} else {
		c.Locals("filename", nil)
	}

	return c.Next()
}

func HandleMultipleFile(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		log.Println("error read multipart from request,error = ", err)
	}

	files := form.File["photos"]

	var filenames []string

	for i, file := range files {
		var filename string
		if file != nil {
			extestionFile := filepath.Ext(file.Filename)
			filename = fmt.Sprintf("%d-%s%s", i, "image", extestionFile)

			errSaveFile := c.SaveFile(file, fmt.Sprintf("./public/covers/%s", filename))
			if errSaveFile != nil {
				log.Println("Upload Fail")
			}
		} else {
			log.Println("Nothing file to upload")
		}

		if filename != "" {
			filenames = append(filenames, filename)
		}
	}
	c.Locals("filenames", filenames)

	return c.Next()
}

func HandleRemoveFile(filename string, pathFIle ...string) error {
	if len(pathFIle) > 0 {
		err := os.Remove(pathFIle[0] + filename)
		if err != nil {
			log.Println("Failed to remove file")
			return err
		}
	} else {
		err := os.Remove(DefaultPathAssetImage + filename)
		if err != nil {
			log.Println("Failed to remove file")
			return err
		}
	}

	return nil
}

func checkContentFile(file *multipart.FileHeader, contentTypes ...string) error {
	if len(contentTypes) > 0 {
		for _, contentType := range contentTypes {
			contentTypeFile := file.Header.Get("Content-Type")
			if contentTypeFile == contentType {
				return nil
			}
		}
		return nil
	} else {
		return errors.New("not found content type")
	}
}
