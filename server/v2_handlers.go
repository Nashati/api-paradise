package server

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *serverImpl) registerV2() {

	s.apiGroup.GET("/jokes/getAll", s.handleJokesGetAllV2)
	s.apiGroup.GET("/jokes/get/:id", s.handleGetJokeByIdV2)
	s.apiGroup.POST("/jokes/createNew", s.okStatusHandlerV2)
	s.apiGroup.POST("/jokes/update/:id", s.okStatusHandlerV2)
	s.apiGroup.POST("/jokes/delete/:id", s.okStatusHandlerV2)

	s.apiGroup.GET("/categories/getAll", s.handleCategoriesGetAllV2)
	s.apiGroup.GET("/categories/get/:id", s.handleGetCategoryByIdV2)
	s.apiGroup.POST("/categories/createNew", s.okStatusHandlerV2)
	s.apiGroup.POST("/categories/update/:id", s.okStatusHandlerV2)
	s.apiGroup.POST("/categories/delete/:id", s.okStatusHandlerV2)
}

func (s *serverImpl) handleJokesGetAllV2(c echo.Context) error {
	print("Handling get all jokes - v2")
	return c.JSON(http.StatusOK, s.store.GetJokes())
}

func (s *serverImpl) handleGetJokeByIdV2(c echo.Context) error {
	print("Handling get joke by id - v2")
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

func (s *serverImpl) handleCategoriesGetAllV2(c echo.Context) error {
	print("Handling v1 get all categories - v2")
	return c.JSON(http.StatusOK, s.store.GetCategories())
}

func (s *serverImpl) handleGetCategoryByIdV2(c echo.Context) error {
	print("Handling get category by id - v2")
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

func (s *serverImpl) okStatusHandlerV2(c echo.Context) error {
	print("Handling create/update/delete - v2")
	return c.NoContent(http.StatusOK)
}
