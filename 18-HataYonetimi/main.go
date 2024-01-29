package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {

	// Hata Yönetimi

	file, err := os.Open("dosyam.txt")
	if err != nil {
		// log.Fatal(err) programı sonlandırır.
		fmt.Println(err.Error())
	}
	defer file.Close()

	myError := errors.New("Bu bir hata mesajıdır.")
	log.Println(myError)

	_, err = os.Open("abc.txt")
	checkError(err)

}

func checkError(err error) {
	if err != nil {
		//fmt.Println("ERROR : ", err.Error())
		//os.Exit(1)
		log.Fatal("ERROR : ", err)
	}
}
