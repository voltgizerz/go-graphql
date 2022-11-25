package utils

func GetSafeInt(num *int) int {
	if num != nil {
		return *num
	}
	return 0
}
