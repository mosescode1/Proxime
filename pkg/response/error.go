package response

import "github.com/gofiber/fiber/v2"

type StatusCode struct {
	InvalidRequest      int
	InvalidCredentials  int
	Unauthorized        int
	Forbidden           int
	NotFound            int
	Conflict            int
	InternalError       int
	UnprocessableEntity int
}
type ErrorResponse struct {
	InvalidRequest      string
	InvalidToken        string
	InvalidCredentials  string
	Unauthorized        string
	Forbidden           string
	NotFound            string
	Conflict            string
	InternalError       string
	UnprocessableEntity string
}

var Errors = ErrorResponse{
	InvalidRequest:      "Invalid request",
	InvalidToken:        "Invalid token",
	InvalidCredentials:  "Invalid credentials",
	Unauthorized:        "Unauthorized",
	Forbidden:           "Forbidden",
	NotFound:            "Not found",
	Conflict:            "Conflict",
	InternalError:       "Internal error",
	UnprocessableEntity: "Unprocessable entity",
}

var StatusCodes = StatusCode{
	InvalidRequest:      400,
	InvalidCredentials:  401,
	Unauthorized:        401,
	Forbidden:           403,
	NotFound:            404,
	Conflict:            409,
	InternalError:       500,
	UnprocessableEntity: 422,
}

// Invalid Request Response: 400
func InvalidRequestResponse(c *fiber.Ctx, msg string, errors ...any) error {
	if msg == "" {
		msg = Errors.InvalidRequest
	}
	return c.Status(StatusCodes.InvalidRequest).JSON(fiber.Map{
		"message": msg,
		"errors":  errors,
	})
}

// Invalid Credentials Response: 401
func InvalidCredentials(c *fiber.Ctx, msg string, errors ...any) error {
	if msg == "" {
		msg = Errors.InvalidCredentials
	}
	return c.Status(StatusCodes.InvalidCredentials).JSON(fiber.Map{
		"message": msg,
		"errors":  errors,
	})
}

// Forbidden Response: 403
func ForbiddenResponse(c *fiber.Ctx, msg string, errors ...any) error {
	if msg == "" {
		msg = Errors.Forbidden
	}
	return c.Status(StatusCodes.Forbidden).JSON(fiber.Map{
		"message": msg,
		"errors":  errors,
	})
}

// Not Found Response: 404
func NotFoundResponse(c *fiber.Ctx, msg string, errors ...any) error {
	if msg == "" {
		msg = Errors.NotFound
	}
	return c.Status(StatusCodes.NotFound).JSON(fiber.Map{
		"message": msg,
		"errors":  errors,
	})
}

// Conflict Response: 409
func ConflictResponse(c *fiber.Ctx, msg string, errors ...any) error {
	if msg == "" {
		msg = Errors.Conflict
	}
	return c.Status(StatusCodes.Conflict).JSON(fiber.Map{
		"message": msg,
		"errors":  errors,
	})
}

// UnAuthorized Response: 401
func UnauthorizedResponse(c *fiber.Ctx, msg string, errors ...any) error {
	if msg == "" {
		msg = Errors.Unauthorized
	}
	return c.Status(StatusCodes.Unauthorized).JSON(fiber.Map{
		"message": msg,
		"errors":  errors,
	})
}
