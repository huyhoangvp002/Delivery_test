package util

const (
	U = "user"
	A = "admin"
	C = "client"
	D = "driver"
)

func IsValidRole(role string) bool {
	switch role {
	case D, A, U, C:
		return true
	default:
		return false
	}
}
