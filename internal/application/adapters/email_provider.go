package adapters

type EmailProviderInterface interface {
	SendEmail(to string, subject string, body string) error
}
