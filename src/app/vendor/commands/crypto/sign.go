package crypto

import (
	"bufio"
	"crypto"
	"crypto/rand"
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

func SetupSignCommand(rootCmd *cobra.Command) {
	var rsaPrivKey string
	var outFile string
	var method string
	cmd := &cobra.Command{
		Use:   "sign",
		Short: "產生訊息簽名",
		Run: func(cmd *cobra.Command, args []string) {
			bs, err := ioutil.ReadFile(rsaPrivKey)
			if err != nil {
				panic(err)
			}

			block, _ := pem.Decode(bs)
			if block == nil {
				panic("rsa private_key error")
			}

			privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
			if err != nil {
				panic(err)
			}
			rng := rand.Reader

			r := bufio.NewReader(os.Stdin)
			cbs, err := ioutil.ReadAll(r)
			if err != nil {
				panic(err)
			}
			var signature []byte
			switch method {
			case "sha224":
				hashed := sha256.Sum224(cbs)
				signature, err = rsa.SignPKCS1v15(rng, privateKey, crypto.SHA224, hashed[:])
			case "sha256":
				hashed := sha256.Sum256(cbs)
				signature, err = rsa.SignPKCS1v15(rng, privateKey, crypto.SHA256, hashed[:])
			case "sha384":
				hashed := sha512.Sum384(cbs)
				signature, err = rsa.SignPKCS1v15(rng, privateKey, crypto.SHA384, hashed[:])
			case "sha512":
				hashed := sha512.Sum512(cbs)
				signature, err = rsa.SignPKCS1v15(rng, privateKey, crypto.SHA512, hashed[:])
			default:
				hashed := sha1.Sum(cbs)
				signature, err = rsa.SignPKCS1v15(rng, privateKey, crypto.SHA1, hashed[:])
			}

			if err != nil {
				panic(err)
			}

			err = ioutil.WriteFile(outFile, signature, 0666)
			if err != nil {
				panic(err)
			}
			fmt.Printf("output %s\n", outFile)
		},
	}
	cmd.Flags().StringVarP(&rsaPrivKey, "key", "k", "", "RSA PKCS1 Private Key")
	cmd.Flags().StringVarP(&outFile, "out", "o", "file.sig", "簽名檔輸出")
	cmd.Flags().StringVarP(&method, "method", "m", "sha256", "簽名演算法, sha1, sha224, sha256, sha384, sha512")

	rootCmd.AddCommand(cmd)
}
