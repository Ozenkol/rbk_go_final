package adapters

type TelephonyProviderInterface interface {
	LogCall(from string, to string, duration int, notes string) error
}
