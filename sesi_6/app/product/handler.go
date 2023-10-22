package product

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return Handler{
		service: service,
	}
}

func (h Handler) CreateProduct(c *fiber.Ctx) error {
	var req = CreateProductRequest{}

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		})
	}

	model := Product{
		Name:     req.Name,
		Category: req.Category,
		Price:    req.Price,
		Stock:    req.Stock,
	}

	err = h.service.CreateProduct(c.UserContext(), model)
	if err != nil {
		var payload fiber.Map
		httpCode := 400

		switch err {
		case ErrEmptyName, ErrEmptyCategory, ErrEmptyPrice, ErrEmptyStock:
			payload = fiber.Map{
				"success": false,
				"message": "ERR BAD REQUEST",
				"error":   err.Error(),
			}
			httpCode = http.StatusBadRequest
		default:
			payload = fiber.Map{
				"success": false,
				"message": "ERR INTERNAL",
				"error":   err.Error(),
			}
			httpCode = http.StatusInternalServerError
		}
		return c.Status(httpCode).JSON(payload)
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "CREATE SUCCESS",
	})

}

func (h Handler) GetProducts(c *fiber.Ctx) error {

	products, err := h.service.GetProducts(c.UserContext())
	if err != nil {
		var payload fiber.Map
		httpCode := 400

		payload = fiber.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		}
		httpCode = http.StatusBadRequest

		return c.Status(httpCode).JSON(payload)
	}
	response := FormatProducts(products)
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "GET DATA SUCCESS",
		"data":    response,
	})

}

func (h Handler) GetProduct(c *fiber.Ctx) error {

	idStr := c.Params("id")
	id, _ := strconv.Atoi(idStr)

	product, err := h.service.GetProduct(c.UserContext(), id)
	if err != nil {
		var payload fiber.Map
		httpCode := 400

		payload = fiber.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		}
		httpCode = http.StatusBadRequest

		return c.Status(httpCode).JSON(payload)
	}
	response := FormatProduct(product)
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "GET DATA SUCCESS",
		"data":    response,
	})

}

func (h Handler) UpdateProduct(c *fiber.Ctx) error {
	var req = CreateProductRequest{}

	idStr := c.Params("id")
	reqId, _ := strconv.Atoi(idStr)

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		})
	}

	model := Product{
		Id:       reqId,
		Name:     req.Name,
		Category: req.Category,
		Price:    req.Price,
		Stock:    req.Stock,
	}

	err = h.service.UpdateProduct(c.UserContext(), model)
	if err != nil {
		var payload fiber.Map
		httpCode := 400

		switch err {
		case ErrEmptyName, ErrEmptyCategory, ErrEmptyPrice, ErrEmptyStock:
			payload = fiber.Map{
				"success": false,
				"message": "ERR BAD REQUEST",
				"error":   err.Error(),
			}
			httpCode = http.StatusBadRequest
		default:
			payload = fiber.Map{
				"success": false,
				"message": "ERR INTERNAL",
				"error":   err.Error(),
			}
			httpCode = http.StatusInternalServerError
		}
		return c.Status(httpCode).JSON(payload)
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "UPDATE SUCCESS",
	})

}

func (h Handler) DeleteProduct(c *fiber.Ctx) error {

	idStr := c.Params("id")
	id, _ := strconv.Atoi(idStr)

	err := h.service.DeleteProduct(c.UserContext(), id)
	if err != nil {
		var payload fiber.Map
		httpCode := 400

		payload = fiber.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		}
		httpCode = http.StatusBadRequest

		return c.Status(httpCode).JSON(payload)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "DELETE PRODUCT SUCCESS",
	})

}
