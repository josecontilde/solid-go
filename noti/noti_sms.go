package noti

import "fmt"

type NotificadorSMS struct{}

func (n NotificadorSMS) Enviar(mensaje, destinatario string) bool {
	fmt.Printf("📱 SMS a %s: %s\n", destinatario, mensaje)
	return true
}
