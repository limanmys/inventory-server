package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"os"

	"github.com/limanmys/inventory-server/app/entities"
)

func EncryptProfile(profile entities.Profile) (entities.Profile, error) {
	// Set error variable for later use
	var err error
	// Encode username
	profile.Username, err = encrypt(profile.Username)
	if err != nil {
		return profile, nil
	}

	// Encode password
	profile.Password, err = encrypt(profile.Password)
	if err != nil {
		return profile, nil
	}

	return profile, nil
}

func encrypt(message string) (encoded string, err error) {
	// Set key
	key := []byte(os.Getenv("APP_KEY"))
	// Create byte array from the input string
	plainText := []byte(message)

	// Create a new AES cipher using the key
	block, err := aes.NewCipher(key)

	//IF NewCipher failed, exit:
	if err != nil {
		return
	}

	// Make the cipher text a byte array of size BlockSize + the length of the message
	cipherText := make([]byte, aes.BlockSize+len(plainText))

	// iv is the ciphertext up to the blocksize (16)
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	// Encrypt the data:
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	// Return string encoded in base64
	return base64.RawStdEncoding.EncodeToString(cipherText), err
}
