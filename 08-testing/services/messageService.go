package services

type MessageService interface {
	Send(msg string) bool
}
