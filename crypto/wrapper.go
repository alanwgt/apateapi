package crypto

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"os"

	"github.com/alanwgt/apateapi/cache"

	"github.com/alanwgt/apateapi/protos"

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

// GetServerB64PubK returns the public key encoded in base64
func GetServerB64PubK() string {
	return b64PubKey
}

// GetServerSecK returns a pointer to the raw secret key
func GetServerSecK() *[32]byte {
	return secretKey
}

// OpenUserBox opens a crypto box and returns the raw message
func OpenUserBox(dr *protos.DeviceRequest) ([]byte, *cache.UserCache, error) {
	// we need to decode all the base64 data first
	bPayload, err := base64.StdEncoding.DecodeString(dr.Paylod)

	if err != nil {
		return nil, nil, errors.New("Payload couldn't be decoded")
	}

	bNonce, err := base64.StdEncoding.DecodeString(dr.Nonce)

	if err != nil {
		return nil, nil, errors.New("Nonce couldn't be decoded")
	}

	// u, err := ca.GetUser(dr.Username)
	u, err := cache.GetUser(dr.Username)

	if err != nil {
		return nil, nil, err
	}

	var out []byte
	var nonce [24]byte

	copy(nonce[:], bNonce)
	tout, ok := box.Open(
		out,
		bPayload,
		&nonce,
		u.PubK,
		secretKey,
	)

	if !ok {
		return nil, nil, errors.New("Couldn't open the box")
	}

	return tout, u, nil
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
