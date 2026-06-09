package telephony

import (
	"fmt"
	"github.com/Ozenkol/rbk-go-final/internal/application/adapters"
)

type TelephonyProvider struct{}

func NewTelephonyProvider() adapters.TelephonyProviderInterface {
	return &TelephonyProvider{}
}

func (p *TelephonyProvider) LogCall(from string, to string, duration int, notes string) error {
	fmt.Printf("Logging call from %s to %s, duration: %d, notes: %s\n", from, to, duration, notes)
	return nil
}
