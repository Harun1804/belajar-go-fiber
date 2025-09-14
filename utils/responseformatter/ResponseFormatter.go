package responseformatter

import "github.com/gofiber/fiber/v2"

type SuccessResponse struct {
	Status     bool        `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
}

type Pagination struct {
	Page      int `json:"page"`
	PageSize  int `json:"pageSize"`
	Total     int `json:"total"`
	TotalPage int `json:"totalPage"`
}

type ErrorResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
}

// Helper to send error response
func SendError(c *fiber.Ctx, status int, message string, errors ...interface{}) error {
	var errDetail interface{}
	if len(errors) > 0 {
		errDetail = errors[0]
	}
	return c.Status(status).JSON(ErrorResponse{
		Status:  false,
		Message: message,
		Errors:  errDetail,
	})
}

// Helper to send success response
func SendSuccess(c *fiber.Ctx, message string, data ...interface{}) error {
	var respData interface{}
	if len(data) > 0 {
		respData = data[0]
	}
	return c.Status(fiber.StatusOK).JSON(SuccessResponse{
		Status:     true,
		Message:    message,
		Data:       respData,
		Pagination: nil,
	})
}

func SendWithPaginationSuccess(c *fiber.Ctx, message string, data interface{}, page, pageSize, totalData, totalPage int) error {
	pagination := Pagination{
		Page:      page,
		PageSize:  pageSize,
		Total:     totalData,
		TotalPage: totalPage,
	}

	return c.Status(fiber.StatusOK).JSON(SuccessResponse{
		Status:     true,
		Message:    message,
		Data:       data,
		Pagination: pagination,
	})
}
