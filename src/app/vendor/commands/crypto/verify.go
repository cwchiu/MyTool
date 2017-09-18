package crypto

import (
	"bufio"
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

func SetupVerifyCommand(rootCmd *cobra.Command) {
	var rsaPubKey string
	var signature string
	var method string

	cmd := &cobra.Command{
		Use:   "verify",
		Short: "產生訊息簽名",
		Run: func(cmd *cobra.Command, args []string) {
			bs, err := ioutil.ReadFile(rsaPubKey)
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

			signature_data, err := ioutil.ReadFile(signature)
			if err != nil {
				panic(err)
			}

			r := bufio.NewReader(os.Stdin)
			cbs, err := ioutil.ReadAll(r)
			if err != nil {
				panic(err)
			}
			switch method {
			case "sha224":
				hashed := sha256.Sum224(cbs)
				err = rsa.VerifyPKCS1v15(rsaPub, crypto.SHA224, hashed[:], signature_data)
			case "sha256":
				hashed := sha256.Sum256(cbs)
				err = rsa.VerifyPKCS1v15(rsaPub, crypto.SHA256, hashed[:], signature_data)
			case "sha384":
				hashed := sha512.Sum384(cbs)
				err = rsa.VerifyPKCS1v15(rsaPub, crypto.SHA384, hashed[:], signature_data)
			case "sha512":
				hashed := sha512.Sum512(cbs)
				err = rsa.VerifyPKCS1v15(rsaPub, crypto.SHA512, hashed[:], signature_data)
			default:
				hashed := sha1.Sum(cbs)
				err = rsa.VerifyPKCS1v15(rsaPub, crypto.SHA1, hashed[:], signature_data)
			}

			if err != nil {
				panic(err)
			}
			fmt.Println("verify OK")

		},
	}
	cmd.Flags().StringVarP(&rsaPubKey, "key", "k", "", "RSA PKCS1 Public Key")
	cmd.Flags().StringVarP(&signature, "signature", "s", "", "signature hex 字串")
	cmd.Flags().StringVarP(&method, "method", "m", "sha256", "簽名演算法, sha1, sha224, sha256, sha384, sha512")
	rootCmd.AddCommand(cmd)
}
