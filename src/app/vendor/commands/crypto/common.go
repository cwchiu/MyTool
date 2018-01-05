package crypto


import (
    // "fmt"
    "crypto/aes"
    "crypto/cipher"
    "io"
    "io/ioutil"
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha512"
    "crypto/x509"
    "encoding/pem"
)


func aesDecrypt(fn string, keystring string)  {
	// Byte array of the string
	// ciphertext := []byte(cipherstring)
    ciphertext, err := ioutil.ReadFile(fn)
    if err != nil {
        panic(err)
    }
    
	// Key
	key := []byte(keystring)

	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Before even testing the decryption,
	// if the text is too small, then it is incorrect
	if len(ciphertext) < aes.BlockSize {
		panic("Text is too short")
	}

	// Get the 16 byte IV
	iv := ciphertext[:aes.BlockSize]

	// Remove the IV from the ciphertext
	ciphertext = ciphertext[aes.BlockSize:]

	// Return a decrypted stream
	stream := cipher.NewCFBDecrypter(block, iv)

	// Decrypt bytes from ciphertext
	stream.XORKeyStream(ciphertext, ciphertext)

    err = ioutil.WriteFile("dec-" + fn, ciphertext, 0666)
    if err != nil {
        panic(err)
    }
}

func aesEncrypt(fn, keystring string)  {
	// Byte array of the string
	// plaintext := []byte(plainstring)
    plainstring, err := ioutil.ReadFile(fn)
    if err != nil {
        panic(err)
    }
	// Key
	key := []byte(keystring)

	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Empty array of 16 + plaintext length
	// Include the IV at the beginning
	ciphertext := make([]byte, aes.BlockSize+len(plainstring))

	// Slice of first 16 bytes
	iv := ciphertext[:aes.BlockSize]

	// Write 16 rand bytes to fill iv
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	// Return an encrypted stream
	stream := cipher.NewCFBEncrypter(block, iv)

	// Encrypt bytes from plaintext to ciphertext
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plainstring)
    err = ioutil.WriteFile("enc-" + fn, ciphertext, 0666)
    if err != nil {
        panic(err)
    }

}

func rsaKeyEncrypt(fn, public_key string) {

    bs, err := ioutil.ReadFile(public_key)
    if err != nil {
        panic(err)
    }

    block, _ := pem.Decode(bs)
    if block == nil {
        panic("rsa public_key error")
    }
    pub, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        panic(err)
    }
    rsaPub, ok := pub.(*rsa.PublicKey)
    if !ok {
        panic("Value returned from ParsePKIXPublicKey was not an RSA public key")
    }
    
    msg, err := ioutil.ReadFile(fn)
    if err != nil {
        panic(err)
    }
    
    encryptOAEP, err := rsa.EncryptOAEP(sha512.New(), rand.Reader, rsaPub, msg, nil)
    if err != nil {
        panic(err)
    }
    
    err = ioutil.WriteFile("enc-" + fn, encryptOAEP, 0666)
    if err != nil {
        panic(err)
    }
}

func rsaKeyDecrypt(fn, private_key string) {
    privateKeyData, err := ioutil.ReadFile(private_key)
    if err != nil {
        panic(err)
    }

    // 解析出私鑰
    priBlock, _ := pem.Decode([]byte(privateKeyData))
    if priBlock == nil {
        panic("rsa private_key error")
    }
    priKey, err := x509.ParsePKCS1PrivateKey(priBlock.Bytes)
    if err != nil {
        panic(err)
    }
    msg, err := ioutil.ReadFile(fn)
    if err != nil {
        panic(err)
    }
    
    // 解密RSA-OAEP方式加密後的內容
    decryptOAEP, err := rsa.DecryptOAEP(sha512.New(), rand.Reader, priKey, msg, nil)
    if err != nil {
        panic(err)
    }
    
    err = ioutil.WriteFile("dec-" + fn, decryptOAEP, 0666)
    if err != nil {
        panic(err)
    }
}