package adapters

type MessageProviderInterface interface {
	SendMessage(to string, message string) error
}