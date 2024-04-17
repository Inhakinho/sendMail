package main

import (
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Carga las variables de entorno desde el archivo .env
	err := godotenv.Load() // Asume que el archivo .env está en el mismo directorio que tu código
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Configura los parámetros del servidor SMTP y las credenciales de autenticación
	from := os.Getenv("EMAIL")              //tu correo electronico
	password := os.Getenv("EMAIL_PASSWORD") // La contraseña de tu correo de Gmail
	to := []string{"inhaki.pi@gmail.com"}   // correo de destino para las pruebas

	// Define el servidor SMTP y el puerto
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Estructura el mensaje que deseas enviar
	message := []byte("To: inhaki.pi@gmail.com\r\n" + //remitente
		"Subject: Saludos desde Go!\r\n" + //asunto
		"\r\n" +
		"Hola Amigo, esto es un correo de prueba enviado desde una aplicación Go.\r\n") //mensaje

	// Configura la autenticación SMTP
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Intenta enviar el correo
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		panic(err)
	} else {
		println("Correo enviado correctamente!")
	}
}
