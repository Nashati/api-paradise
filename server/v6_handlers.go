package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/api-paradise/model"
	"github.com/labstack/echo/v4"
)

func (s *serverImpl) registerV6() {
	
	s.apiGroup.GET("/v6/jokes", s.handleJokesGetAllV6)
	s.apiGroup.GET("/v6/jokes/:id", s.handleJokeByIdV6)
	s.apiGroup.POST("/v6/jokes", s.createdStatusHandlerV6)
}

func (s *serverImpl) handleJokesGetAllV6(c echo.Context) error {
	print("Handling get all jokes - v6")
	return c.JSON(http.StatusOK, s.store.GetJokes())
}

func (s *serverImpl) handleJokeByIdV6(c echo.Context) error {
	print("Handling get joke by id - v6")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}
	c.Response().Header().Set("Link", "<http://localhost:3030/v6/jokes/"+strconv.Itoa(id)+"?historyNumber=2&historySize=100>; rel=\"next\", <http://localhost:3030/v6/jokes/"+strconv.Itoa(id)+"?historyNumber=3&historySize=100>; rel=\"last\"")
	j, err := s.store.GetJoke(int64(id))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, *j)
}

func (s *serverImpl) createdStatusHandlerV6(c echo.Context) error {
	print("Handling create/update/delete - v6")
	var joke model.Joke
	if c.Bind(&joke) != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	joke.Id = 123
	joke.LastUpdated = time.Now().Unix()
	c.Response().Header().Set("Location", "/api/v6/jokes/123")

	return c.JSON(http.StatusCreated, joke)
}
