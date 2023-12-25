package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"os"

	"github.com/limanmys/inventory-server/app/entities"
)

func DecryptProfile(profile entities.Profile) (entities.Profile, error) {
	// Set error variable for later use
	var err error

	// Encode username
	profile.Username, err = decrypt(profile.Username)
	if err != nil {
		return profile, nil
	}

	// Encode password
	profile.Password, err = decrypt(profile.Password)
	if err != nil {
		return profile, nil
	}

	return profile, nil
}

func decrypt(secure string) (decoded string, err error) {
	// Set key
	key := []byte(os.Getenv("APP_KEY"))
	//Remove base64 encoding:
	cipherText, err := base64.RawStdEncoding.DecodeString(secure)

	//IF DecodeString failed, exit:
	if err != nil {
		return
	}

	//Create a new AES cipher with the key and encrypted message
	block, err := aes.NewCipher(key)

	//IF NewCipher failed, exit:
	if err != nil {
		return
	}

	//IF the length of the cipherText is less than 16 Bytes:
	if len(cipherText) < aes.BlockSize {
		err = errors.New("ciphertext block size is too short")
		return
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	//Decrypt the message
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), err
}
