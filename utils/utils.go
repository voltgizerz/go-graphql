package utils

// GetSafeInt - get pointer number safely prevent panic
func GetSafeInt(num *int) int {
	if num != nil {
		return *num
	}

	return 0
}
