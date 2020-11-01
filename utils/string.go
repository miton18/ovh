package utils

// NullableString return a nullable string
func NullableString(str string) *string {
	if str == "" {
		return nil
	}
	return &str
}

