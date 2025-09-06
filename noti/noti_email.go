package noti

import "fmt"

type NotificadorEmail struct{}

func (n NotificadorEmail) Enviar(mensaje, destinatario string) bool {
	fmt.Printf("📧 Email a %s: %s\n", destinatario, mensaje)
	return true
}
