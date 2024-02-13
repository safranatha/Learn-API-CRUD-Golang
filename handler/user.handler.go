package handler

import (
	"go-fiber/database"
	"go-fiber/model/entity"
	"go-fiber/model/request"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserHandler(c *fiber.Ctx) error {
	var users []entity.User
	result := database.DB.Debug().Find(&users)
	if result.Error != nil {
		log.Println("error")
	}
	return c.JSON(users)
}

var validate = validator.New()

func UserCreate(c *fiber.Ctx) error {

	// create request
	user := new(request.UserCreateReq)
	err := c.BodyParser(user)
	if err != nil {
		return err
	}

	// validate
	err_val := validate.Struct(user)
	if err_val != nil {
		return err_val
	}

	// define apa yang bisa diisi (modelnya)
	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Phone:   user.Phone,
		Address: user.Address,
	}

	// error handling
	errCreate := database.DB.Create(&newUser).Error
	if errCreate != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": errCreate})
	}

	// jika sukses maka akan seperti ini
	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})

}

func GetDataById(c *fiber.Ctx) error {
	id := c.Params("id")

	user := entity.User{}

	err := database.DB.First(&user, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Product found", "data": user})

}

func UpdateData(c *fiber.Ctx) error {
	// create request
	userReq := new(request.UserUpdateReq)
	errReq := c.BodyParser(userReq)
	if errReq != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Bad Request"})
	}

	// get Data byID
	id := c.Params("id")

	user := entity.User{}

	err := database.DB.First(&user, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
	}

	// update
	// update only if fields are present in the request
	if userReq.Name != "" {
		user.Name = userReq.Name
	}

	if userReq.Address != "" {
		user.Address = userReq.Address
	}
	if userReq.Phone != "" {
		user.Phone = userReq.Phone
	}

	err_update := database.DB.Save(&user).Error
	if err_update != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update user"})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Update berhasil", "data": user})

}

func UpdateEmailData(c *fiber.Ctx) error {
	// create request
	userReq := new(request.UserUpdateEmailReq)
	errReq := c.BodyParser(userReq)
	if errReq != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Bad Request"})
	}

	// get Data byID
	id := c.Params("id")

	user := entity.User{}

	err := database.DB.First(&user, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
	}

	// update
	// update only if fields are present in the request
	if userReq.Email != "" {
		user.Email = userReq.Email
	}

	err_update := database.DB.Save(&user).Error
	if err_update != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update user"})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Update berhasil", "data": user})
}

func UserDelete(c *fiber.Ctx) error {
	// get Data byID
	id := c.Params("id")

	user := entity.User{}

	err := database.DB.First(&user, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
	}

	delete := database.DB.Delete(&user).Error
	if delete != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete user"})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Delete user success"})
}
