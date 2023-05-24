package main

import (
	"github.com/rkaneko95/golang-study-web-auth/introduction"
	"log"
	"net/http"
)

func main() {
	testHash()

	/*setHandleFunc()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic(err)
	}*/
}

func setHandleFunc() {
	http.HandleFunc("/encode", introduction.Foo)
	http.HandleFunc("/decode", introduction.Bar)
}

func testHash() {
	pass := "123456789"

	hashedPass, err := hashPassword(pass)
	if err != nil {
		panic(err)
	}

	err = comparePassword(pass, hashedPass)
	if err != nil {
		log.Fatalln("Not logged in")
	}

	log.Println("Logged in!")
}
