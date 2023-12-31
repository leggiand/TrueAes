package TrueAes

import (
	"crypto/aes"
	"encoding/hex"
	"github.com/zenazn/pkcs7pad"
)

// Encrypt A simple function for encrypting a 16 bytes string
func Encrypt(key string, mess string) string {
	cypher, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	encrypted := make([]byte, len(mess))
	cypher.Encrypt(encrypted, []byte(mess))
	return hex.EncodeToString(encrypted)
}

// Decrypt A simple function for decrypting a 16 bytes string
func Decrypt(key *string, mess string) string {
	txt, _ := hex.DecodeString(mess)
	cypher, err := aes.NewCipher([]byte(*key))
	if err != nil {
		panic(err)
	}
	decryptedByte := make([]byte, len(txt))
	cypher.Decrypt(decryptedByte, txt)
	decrypted := string(decryptedByte)
	return decrypted
}
func TotalUnpad(topad *string) string {
	var result = ""
	if len(*topad) < 16 {
		tmp, _ := pkcs7pad.Unpad([]byte(*topad))
		result = string(tmp)
		return result
	} else if len(*topad) == 16 {

		tmp, _ := pkcs7pad.Unpad([]byte(*topad))
		result = string(tmp)
	} else {
		result = ""
		buffer := make([]byte, 0)
		for index, element := range []byte(*topad) {
			buffer = append(buffer, element)
			if (index+1)%16 == 0 {
				tmp, err := pkcs7pad.Unpad(buffer)
				if err != nil {
					result += string(buffer)
				}
				result += string(tmp)
				buffer = nil
			}
		}
		tmp, _ := pkcs7pad.Unpad(buffer)
		result += string(tmp)
	}
	return result
}

// A recursive function to apply the Pkcs7padding to a string not multiple of 16
func TotalPad(topad string) string {
	var result = ""
	if len(topad) < 16 {
		result = string(pkcs7pad.Pad([]byte(topad), 16))
		return result
	} else if len(topad) == 16 {

		return topad
	} else {
		result = ""
		buffer := make([]byte, 0)
		for index, element := range []byte(topad) {
			buffer = append(buffer, element)
			if (index+1)%16 == 0 {
				result += string(buffer)
				buffer = nil
			}
		}
		result += string(pkcs7pad.Pad(buffer, 16))
	}
	return result
}

// A Simple Loop to encrypt strings larger then 16bytes
func FullEncrypt(key string, todecr string) string {
	var result = ""
	buffer := make([]byte, 0)
	for index, element := range []byte(todecr) {
		buffer = append(buffer, element)
		if (index+1)%16 == 0 {
			result += Encrypt(key, string(buffer))
			buffer = nil
		}
	}
	return result
}

// A Simple Loop to decrypt strings larger then 16bytes
func FullDecrypt(key string, todecr string) string {
	var result = ""
	buffer := make([]byte, 0)
	for index, element := range []byte(todecr) {
		buffer = append(buffer, element)
		if (index+1)%32 == 0 {
			result += Decrypt(&key, string(buffer))
			buffer = nil
		}
	}
	return result
}
