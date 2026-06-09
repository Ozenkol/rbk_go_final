package adapters

type MessageProviderInterface interface {
	SendMessage(to string, message string) error
	SendWhatsApp(to string, message string) error
	SendSMS(to string, message string) error
}
