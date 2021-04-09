package commons

func Contains(list []string, value string) bool {
	if len(list) == 0 {
		return false
	}

	for _, v := range list {
		return v == value
	}

	return false
}