package communication

import (
	"fmt"
	"github.com/Ozenkol/rbk-go-final/internal/application/adapters"
)

type MessageProvider struct{}

func NewMessageProvider() adapters.MessageProviderInterface {
	return &MessageProvider{}
}

func (p *MessageProvider) SendMessage(to string, message string) error {
	fmt.Printf("Sending message to %s: %s\n", to, message)
	return nil
}

func (p *MessageProvider) SendWhatsApp(to string, message string) error {
	fmt.Printf("Sending WhatsApp to %s: %s\n", to, message)
	return nil
}

func (p *MessageProvider) SendSMS(to string, message string) error {
	fmt.Printf("Sending SMS to %s: %s\n", to, message)
	return nil
}
