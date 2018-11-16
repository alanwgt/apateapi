package crypto

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"os"

	"golang.org/x/crypto/nacl/box"

	"github.com/alanwgt/apateapi/util"
)

var publicKey, secretKey *[32]byte
var b64PubKey string

func init() {
	_, err := os.Stat(util.Conf.Crypto.SecKeyFile)

	// this means that the file doesn't exist
	if os.IsNotExist(err) {
		log.Println("Crypto Box not found! Generating one...")
		pubKey, secKey, err := box.GenerateKey(rand.Reader)

		if err != nil {
			log.Println("Couldn't generate a keypair from box!")
			log.Panic(err)
		}

		publicKey = pubKey
		secretKey = secKey
		storeSecKey()
	} else {
		secretKey, publicKey = loadKeyPair()
	}

	b64PubKey = base64.StdEncoding.EncodeToString(publicKey[:])
}

// GetB64PubK returns the public key encoded in base64
func GetB64PubK() string {
	return b64PubKey
}

// Loads the .der secret key from a .der file
func loadKeyPair() (secKey, pubK *[32]byte) {
	f, err := os.Open(util.Conf.Crypto.SecKeyFile)
	sSecK, sPubK := make([]byte, 32), make([]byte, 32)

	if err != nil {
		log.Println("Couldn't open file:", util.Conf.Crypto.SecKeyFile)
		log.Fatal(err)
	}

	i, err := f.Read(sSecK)
	if err != nil || i != 32 {
		log.Println(sSecK)
		log.Fatal("Couldn't read the secret key from:", util.Conf.Crypto.SecKeyFile)
	}

	f, err = os.Open(util.Conf.Crypto.PubKeyFile)
	if err != nil {
		log.Println("Couldn't open file:", util.Conf.Crypto.SecKeyFile)
		log.Fatal(err)
	}

	i, err = f.Read(sPubK)
	if err != nil || i != 32 {
		log.Fatal("Couldn't read the public key from:", util.Conf.Crypto.PubKeyFile)
	}

	var bSecK, bPubK [32]byte
	copy(bSecK[:], sSecK)
	copy(bPubK[:], sPubK)

	log.Println("Keys successfully loaded from files.")

	return &bSecK, &bPubK
}

// Stores the secret key a .der file
// this means that the key is not encoded
func storeSecKey() {
	if publicKey == nil || secretKey == nil {
		log.Panic("The keys weren't generated! There is nothing to store.")
	}

	f, err := os.Create(util.Conf.Crypto.SecKeyFile)
	if err != nil {
		log.Println("Couldn't create a file in:", util.Conf.Crypto.SecKeyFile)
		log.Fatal(err)
	}

	f.Write(secretKey[:])
	f.Close()

	f, err = os.Create(util.Conf.Crypto.PubKeyFile)
	if err != nil {
		log.Println("Couldn't create a file in:", util.Conf.Crypto.SecKeyFile)
		log.Fatal(err)
	}

	f.Write(publicKey[:])
	f.Close()

	log.Println("Secret and public key successfully stored.")
}
