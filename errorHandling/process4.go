package errorhandling

import (
	"log"
	"os"
)

func Process4() {
	_, err := os.Open("abc.rar")
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal("ERROR : ", err)
	}
}
