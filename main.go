package main

import (
	"github.com/rkaneko95/golang-study-web-auth/introduction"
	"log"
	"net/http"
)

/*
● Hashing
	○ MD5 - don’t use
	○ SHA
	○ Bcrypt
	○ Scrypt
● Signing
	○ Symmetric Key
		■ HMAC
		■ same key to sign (encrypt) / verify (decrypt)
	○ Asymmetric Key
		■ RSA
		■ ECDSA - better than RSA; faster; smaller keys
		■ private key to sign (encrypt) / public key to verify (decrypt)
	○ JWT
● Encryption
	○ Symmetric key
		■ AES
	○ Asymmetric Key
		■ RSA
*/

func main() {
	setHandleFunc()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic(err)
	}
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
