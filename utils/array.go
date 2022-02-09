package utils

func RemoveDuplicateStrings(strSlice []string) (list []string) {
	allKeys := make(map[string]bool)

	for _, item := range strSlice {
		if _, ok := allKeys[item]; !ok {
			allKeys[item] = true
			list = append(list, item)
		}
	}

	return
}
