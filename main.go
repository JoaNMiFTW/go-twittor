package main

import (
	"log"

	"github.com/JoaNMiFTW/go-twittor/bd"
	"github.com/JoaNMiFTW/go-twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}

	handlers.Manejadores()

}
