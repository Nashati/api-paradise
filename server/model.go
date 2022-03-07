package server

import (
	"github.com/api-paradise/store"
	"github.com/labstack/echo/v4"
)

type serverImpl struct {
	echoServer *echo.Echo
	apiGroup   *echo.Group
	port       int
	store      store.Store
}
