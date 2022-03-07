package server

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *serverImpl) registerV1() {

	s.apiGroup.GET("/joke/getAll", s.handleJokesGetAllV1)
	s.apiGroup.GET("/joke/get/:id", s.handleGetJokeByIdV1)
	s.apiGroup.POST("/joke/createNew", s.okStatusHandlerV1)
	s.apiGroup.POST("/joke/update/:id", s.okStatusHandlerV1)
	s.apiGroup.POST("/joke/delete/:id", s.okStatusHandlerV1)

	s.apiGroup.GET("/category/getAll", s.handleCategoriesGetAllV1)
	s.apiGroup.GET("/category/get/:id", s.handleGetCategoryByIdV1)
	s.apiGroup.POST("/category/createNew", s.okStatusHandlerV1)
	s.apiGroup.POST("/category/update/:id", s.okStatusHandlerV1)
	s.apiGroup.POST("/category/delete/:id", s.okStatusHandlerV1)
}

func (s *serverImpl) handleJokesGetAllV1(c echo.Context) error {
	print("Handling v1 get all jokes - v1")
	return c.JSON(http.StatusOK, s.store.GetJokes())
}

func (s *serverImpl) handleGetJokeByIdV1(c echo.Context) error {
	print("Handling get joke by id - v1")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	j, err := s.store.GetJoke(int64(id))
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, *j)
}

func (s *serverImpl) handleCategoriesGetAllV1(c echo.Context) error {
	print("Handling v1 get all categories - v1")
	return c.JSON(http.StatusOK, s.store.GetCategories())
}

func (s *serverImpl) handleGetCategoryByIdV1(c echo.Context) error {
	print("Handling get category by id - v1")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	cat, err := s.store.GetCategory(int64(id))
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, *cat)
}

func (s *serverImpl) okStatusHandlerV1(c echo.Context) error {
	print("Handling create/update/delete - v1")
	return c.NoContent(http.StatusOK)
}
