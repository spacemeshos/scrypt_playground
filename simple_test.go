package scrypt_playground

import (
	"crypto/rand"
	"golang.org/x/crypto/scrypt"
	"log"
	"testing"
)

func BenchmarkBasic(b *testing.B) {
	salt := [32]byte{}
	_, err := rand.Read(salt[:])
	if err != nil {
		log.Fatal(err)
	}

	key := [64]byte{}
	_, err = rand.Read(key[:])
	if err != nil {
		log.Fatal(err)
	}

	b.ResetTimer()

	// salt := []byte{0xc8, 0x28, 0xf2, 0x58, 0xa7, 0x6a, 0xad, 0x7b}
	// N=512, r=1, p=1
	for i := 0; i < b.N; i++ {
		_, err = scrypt.Key(key[:], salt[:], 1<<22, 1, 1, 1)
		if err != nil {
			log.Fatal(err)
		}
	}
}


func TestBasic(t *testing.T) {
	salt := [32]byte{}
	_, err := rand.Read(salt[:])
	if err != nil {
		log.Fatal(err)
	}

	key := [64]byte{}
	_, err = rand.Read(key[:])
	if err != nil {
		log.Fatal(err)
	}

	// salt := []byte{0xc8, 0x28, 0xf2, 0x58, 0xa7, 0x6a, 0xad, 0x7b}
	// N=512, r=1, p=1
	_, err = scrypt.Key(key[:], salt[:], 1<<22, 1, 1, 1)
	if err != nil {
		log.Fatal(err)
	}

}
