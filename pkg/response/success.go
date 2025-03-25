package response

import "github.com/gofiber/fiber/v2"

type SuccessCode struct {
	Success int
	Created int
	Deleted int
	Updated int
}

type Success struct {
	Success string
	Created string
	Deleted string
	Updated string
}

var SuccessResp = Success{
	Success: "Success",
	Created: "Created",
	Deleted: "Deleted",
	Updated: "Updated",
}

var SuccessCodes = SuccessCode{
	Success: 200,
	Created: 201,
	Deleted: 204,
	Updated: 200,
}

// Success Response: 200
func SuccessResponse(c *fiber.Ctx, msg string, data ...any) error {
	if msg == "" {
		msg = SuccessResp.Success
	}
	return c.Status(SuccessCodes.Success).JSON(fiber.Map{
		"message": msg,
		"data":    data,
	})
}

// Created Response: 201
func CreatedResponse(c *fiber.Ctx, msg string, data ...any) error {
	if msg == "" {
		msg = SuccessResp.Created
	}
	return c.Status(SuccessCodes.Created).JSON(fiber.Map{
		"message": msg,
		"data":    data,
	})
}

// Updated Response: 200
func UpdatedResponse(c *fiber.Ctx, msg string, data ...any) error {
	if msg == "" {
		msg = SuccessResp.Updated
	}
	return c.Status(SuccessCodes.Updated).JSON(fiber.Map{
		"message": msg,
		"data":    data,
	})
}

// Deleted Response: 204
func DeletedResponse(c *fiber.Ctx, msg string) error {
	if msg == "" {
		msg = SuccessResp.Deleted
	}
	return c.Status(SuccessCodes.Deleted).JSON(fiber.Map{
		"message": msg,
	})
}
