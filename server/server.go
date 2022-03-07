package server

import (
	"strconv"

	"github.com/api-paradise/store"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const ApiVersions = 8

type ApiParadiseServer interface {
	Start()
}

func GetServerInstance(port int) ApiParadiseServer {

	s := &serverImpl{
		echoServer: echo.New(),
		port:       port,
		store:      store.GetInstance(),
	}

	s.registerEndpoints()

	return s
}

func (s *serverImpl) Start() {
	port := strconv.Itoa(s.port)
	s.echoServer.Logger.Info("Starting server on port " + port)
	s.echoServer.Logger.Fatal(s.echoServer.Start(":" + port))
}

func (s *serverImpl) registerEndpoints() {
	s.apiGroup = s.echoServer.Group("api")
	s.echoServer.Use(middleware.Static("api"))
	s.registerSwagger()
	s.registerV1()
	s.registerV2()
	s.registerV6()
}

func (s *serverImpl) registerSwagger() {

	for i := 1; i <= ApiVersions; i++ {
		url := echoSwagger.URL("openapi-v" + strconv.Itoa(i) + ".yaml")
		s.apiGroup.GET("/v"+strconv.Itoa(i)+"/swagger/*", echoSwagger.EchoWrapHandler(url))
	}

}
