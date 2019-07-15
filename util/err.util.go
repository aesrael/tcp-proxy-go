package util

import "log"

//HandleErr //generic function to handle all errors
func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
