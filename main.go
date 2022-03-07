package main

import (
	"github.com/api-paradise/server"
)

const PORT = 3030

func main() {
	s := server.GetServerInstance(PORT)
	s.Start()
}
