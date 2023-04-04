package log

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Logger(c *fiber.Ctx) error {
	file, err := os.Create("logfile.log")
	if err != nil {
		log.Fatal("No se pudo crear el archivo de registro:", err)
	}
	defer file.Close()

	logger := log.New(file, "", 0)

	// Establece el formato de fecha y hora y el prefijo
	logger.SetFlags(log.Ldate | log.Ltime)
	logger.SetPrefix("INFO: ")

	realizarAccionImportante(logger)
	return (err)
}

func realizarAccionImportante(logger *log.Logger) {
	logger.Println("La acción importante ha comenzado.")

	// Realiza la acción importante aquí...

	logger.Println("La acción importante se ha completado exitosamente.")
}
