package noti

type Notificador interface {
	Enviar(mensaje, destinatario string) bool
}
