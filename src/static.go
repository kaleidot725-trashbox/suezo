package main

func contains(str string, array []string) bool {
	for _, item := range array {
		if str == item {
			return true
		}
	}
	return false
}
