package main

import "os"

func contains(str string, array []string) bool {
	for _, item := range array {
		if str == item {
			return true
		}
	}
	return false
}

func fcopydir(source string, destination string, mode os.FileMode) (err error) {
	// if !fexist(source) {
	// 	err = os.Mkdir(destination, mode)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	// return err
	return err
}
