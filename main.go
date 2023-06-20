package main

import (
	"bytes"
	"github.com/rkaneko95/golang-study-web-auth/encryption"
	"github.com/rkaneko95/golang-study-web-auth/introduction"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"net/http"
)

func main() {
	/*testBasicAuth()
	testHash()
	testBase64()
	testAes()*/
	testSha()
}

func setHandleFunc() {
	http.HandleFunc("/encode", introduction.Foo)
	http.HandleFunc("/decode", introduction.Bar)
}

func testBasicAuth() {
	setHandleFunc()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic(err)
	}
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

func testBase64() {
	msg := "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
	encoded := encryption.Base64Encoded(msg)
	log.Printf(encoded)

	decoded, err := encryption.Base64Decoded(encoded)
	if err != nil {
		log.Fatalln("error decoded")
	}
	log.Printf(decoded)
}

func testAes() {
	msg := "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."

	pss := "ilovedogs"
	bs, err := bcrypt.GenerateFromPassword([]byte(pss), bcrypt.MinCost)
	if err != nil {
		log.Fatalln(err)
	}
	bs = bs[:16]

	w := &bytes.Buffer{}
	encWriter, err := encryption.EncryptWriter(w, bs)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = io.WriteString(encWriter, msg)
	if err != nil {
		log.Fatalln(err)
	}

	enc := w.String()
	log.Printf("Before: %s", enc)

	rslt, err := encryption.EnDecode(bs, []byte(enc))
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf(string(rslt))

	/*rslt, err = encryption.EnDecode(bs, rslt)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf(string(rslt))*/
}

func testSha() {
	b, err := encryption.Sha()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf(string(b))
}
