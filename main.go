package scrypt_playground

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"log"
)


// import "golang.org/x/crypto/scrypt"
func main() { // run the app

	// DO NOT use this salt value; generate your own random salt. 8 bytes is
	// a good length.

	salt := [32]byte{}
	_, err := rand.Read(salt[:])


	// salt := []byte{0xc8, 0x28, 0xf2, 0x58, 0xa7, 0x6a, 0xad, 0x7b}

	// N=512, r=1, p=1

	key := [32]byte{}
	_, err = rand.Read(key[:])

	dk, err := scrypt.Key(key[:], salt[:], 1<<9, 1, 1, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(dk))
}

