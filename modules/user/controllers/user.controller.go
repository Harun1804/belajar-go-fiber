package controllers

import (
	"belajar-go-fiber/modules/user/dtos"
	"belajar-go-fiber/modules/user/services"
	"belajar-go-fiber/modules/user/validators"
	"belajar-go-fiber/utils/responseformatter"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var userService = services.NewUserService()

func GetAllUsers(c *fiber.Ctx) error {
    page, _ := strconv.Atoi(c.Query("page", "1"))
    pageSize, _ := strconv.Atoi(c.Query("pageSize", "10"))
    sortBy := c.Query("sortBy", "id")
    sortOrder := c.Query("sortOrder", "asc")

    users, totalData, totalPage, err := userService.GetAllUsers(page, pageSize, sortBy, sortOrder)
    if err != nil {
        return responseformatter.SendError(c, fiber.StatusInternalServerError, "Failed to get users", err.Error())
    }
    return responseformatter.SendWithPaginationSuccess(c, "Success Get All Users", users, page, pageSize, totalData, totalPage)
}

func GetUserById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return responseformatter.SendError(c, fiber.StatusBadRequest, "Invalid user ID")
	}
	user, err := userService.GetUserById(id)
	if err != nil {
		return responseformatter.SendError(c, fiber.StatusBadRequest, "User not found")
	}
	return responseformatter.SendSuccess(c, "Success Get User By Id", user)
}

func CreateUsers(c *fiber.Ctx) error {
	user := new(dtos.UserCreateRequest)
	if err := c.BodyParser(user); err != nil {
		return responseformatter.SendError(c, fiber.StatusServiceUnavailable, "Failed to parse body", err.Error())
	}
	if messages, errValidate := validators.ValidateUserCreateRequest(user); errValidate != nil {
		return responseformatter.SendError(c, fiber.StatusUnprocessableEntity, "Failed to validate", messages)
	}
	if _, err := userService.CreateUser(user); err != nil {
		return responseformatter.SendError(c, fiber.StatusInternalServerError, "Failed to create user", err.Error())
	}
	return responseformatter.SendSuccess(c, "Success Create New User")
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return responseformatter.SendError(c, fiber.StatusBadRequest, "Invalid user ID")
	}
	user := new(dtos.UserUpdateRequest)
	user.ID = id
	if err := c.BodyParser(user); err != nil {
		return responseformatter.SendError(c, fiber.StatusBadRequest, "Failed to parse body", err.Error())
	}
	if messages, errValidate := validators.ValidateUserUpdateRequest(user); errValidate != nil {
		return responseformatter.SendError(c, fiber.StatusUnprocessableEntity, "Failed to validate", messages)
	}
	if _, err := userService.UpdateUser(user); err != nil {
		return responseformatter.SendError(c, fiber.StatusInternalServerError, "Failed to update user", err.Error())
	}
	return responseformatter.SendSuccess(c, "Update user successfully")
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return responseformatter.SendError(c, fiber.StatusBadRequest, "Invalid user ID")
	}
	user, err := userService.GetUserById(id)
	if err != nil {
		return responseformatter.SendError(c, fiber.StatusNotFound, "User not found")
	}
	if err := userService.DeleteUser(user.ID); err != nil {
		return responseformatter.SendError(c, fiber.StatusInternalServerError, "Failed to delete user", err.Error())
	}
	return responseformatter.SendSuccess(c, "Delete user successfully")
}
