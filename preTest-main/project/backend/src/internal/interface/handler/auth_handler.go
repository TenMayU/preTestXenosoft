package handler

import (
	"backendrest/src/internal/domain/user"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	handler user.AuthService
}

func NewHttpAuth(service user.AuthService) *AuthHandler {
	return &AuthHandler{handler: service}
}

func (s *AuthHandler) Login(c *fiber.Ctx) error {
	var bodyParse user.User
	if err := c.BodyParser(&bodyParse); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request body",
			"status":  false,
		})
	}

	user, ok, err := s.handler.Login(c.Context(), bodyParse.UserName, bodyParse.Password)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"status":  false,
		})
	}

	if ok == "" {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Loing has fail",
			"status":  false,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user login successfully",
		"user":    user,
		"status":  true,
		"token":   ok,
	})

}

func (s *AuthHandler) Register(c *fiber.Ctx) error {

	var bodyParse user.User

	if err := c.BodyParser(&bodyParse); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	ok, err := s.handler.Register(c.Context(), bodyParse.UserName, bodyParse.Password, bodyParse.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  err.Error(),
			"status": false,
		})
	}

	if !ok {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "register has fail",
			"status":  false,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user registered successfully",
		"status":  true,
	})

}

func (s *AuthHandler) TestAuth(c *fiber.Ctx) error {
	tests := []user.User{
		{UserName: "May", Password: "123"}, // success
		{UserName: "May", Password: "1"},   // wrong password
		{UserName: "M", Password: "123"},   // wrong username
		{UserName: "", Password: "123"},    // missing username
		{UserName: "May", Password: ""},    // missing password
	}

	var results []fiber.Map

	for _, test := range tests {
		u, token, err := s.handler.Login(c.Context(), test.UserName, test.Password)

		result := fiber.Map{
			"input": fiber.Map{
				"username": test.UserName,
				"password": test.Password,
			},
			"user":   u,
			"token":  token,
			"status": "success",
		}

		if err != nil {
			result["status"] = "error"
			result["error"] = err.Error()
		}

		results = append(results, result)
	}

	return c.JSON(results)
}

func (s *AuthHandler) TestRegisterCases(c *fiber.Ctx) error {
	testCases := []user.User{
		{UserName: "may", Password: "123", Name: "May"}, // ✅ success
		{UserName: "", Password: "123", Name: "May"},    // ❌ missing username
		{UserName: "may", Password: "", Name: "May"},    // ❌ missing password
		{UserName: "may", Password: "123", Name: ""},    // ❌ missing name
	}

	var results []fiber.Map

	for _, input := range testCases {
		u, err := s.handler.Register(c.Context(), input.UserName, input.Password, input.Name)

		result := fiber.Map{
			"input":  input,
			"user":   u,
			"status": "success",
		}
		if err != nil {
			result["status"] = "error"
			result["error"] = err.Error()
		}
		results = append(results, result)
	}

	return c.JSON(results)
}
