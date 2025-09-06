package validation

func ErrorMessages() map[string]string {
	return map[string]string{
		"required": "This field is required",
		"email":    "Invalid email format",
		"min":      "Should be more than the limit",
		"max":      "Should be less than the limit",
	}
}
