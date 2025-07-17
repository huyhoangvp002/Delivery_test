package util

const (
	Cr = "created"
	P  = "picked"
	I  = "in_transit"
	De = "delivered"
	F  = "failed"
)

func IsValidStatus(status string) bool {
	switch status {
	case De, I, P, Cr, F:
		return true
	default:
		return false
	}
}
