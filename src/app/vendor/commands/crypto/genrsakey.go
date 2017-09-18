package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func saveGobKey(fileName string, key interface{}) (err error){
    outFile, err := os.Create(fileName)
	if err != nil {
        return err
    }
	defer outFile.Close()

	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
    
    return err
}

func savePEMKey(fileName string, key *rsa.PrivateKey) (err error) {
    outFile, err := os.Create(fileName)
	if err != nil {
        return err
    }
	
	defer outFile.Close()

	err = pem.Encode(outFile, &pem.Block{
		Type:  "RSA PRIVATE KEY",
        Headers: nil,
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})
    
    return err
}

func savePublicPEMKey(fileName string, pubkey rsa.PublicKey) (err error) {
    pub_der, err := x509.MarshalPKIXPublicKey(&pubkey);

	pemfile, err := os.Create(fileName)
    if err != nil {
        panic(err)
    }
	defer pemfile.Close()

	err = pem.Encode(pemfile, &pem.Block{
		Type:  "PUBLIC KEY",
        Headers: nil,
		Bytes: pub_der,
	})
    return err
}

func SetupGenRsaKeyCommand(rootCmd *cobra.Command) {
	var rsaBits int
	cmd := &cobra.Command{
		Use:   "gen-rsa-key",
		Short: "產生RSA公私鑰",
		Run: func(cmd *cobra.Command, args []string) {
			reader := rand.Reader

			key, err := rsa.GenerateKey(reader, rsaBits)
			if err != nil {
                panic(err)
            }
            
            err = key.Validate();
            if err != nil {
                panic(err)
            }
            
			prefix := fmt.Sprintf("rsa_%d_", rsaBits)
			err = saveGobKey(prefix+"private.key", key)
            if err != nil {
                panic(err)
            }
            fmt.Println(prefix+"private.key")
            
			err = savePEMKey(prefix+"private.pem", key)
            if err != nil {
                panic(err)
            }
            fmt.Println(prefix+"private.pem")
            
			publicKey := key.PublicKey
			err = saveGobKey(prefix+"public.key", publicKey)
            if err != nil {
                panic(err)
            }
            fmt.Println(prefix+"public.key")
            
			err = savePublicPEMKey(prefix+"public.pem", publicKey)
            if err != nil {
                panic(err)
            }
            fmt.Println(prefix+"public.pem")
		},
	}
	cmd.Flags().IntVarP(&rsaBits, "rsa-bits", "r", 2048, "Size of RSA key to generate. Ignored if --ecdsa-curve is set")

	rootCmd.AddCommand(cmd)
}
