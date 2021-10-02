package main

import (
	"log"

	"github.com/DiegoCV/twiterGo/bd"
	"github.com/DiegoCV/twiterGo/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexion")
		return
	}
	handlers.Manejadores()
}
