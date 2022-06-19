package helps

import "log"

func CheckFolder(path string) bool {
	log.Println("check")
	return true
}

func CreateFolder(path string) error {
	log.Println("create")

	return nil
}
