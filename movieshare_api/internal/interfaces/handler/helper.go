package handler

func convertStringToStringPointer(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}
