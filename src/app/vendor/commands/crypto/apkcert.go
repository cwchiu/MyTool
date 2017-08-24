package crypto

import (
	"archive/zip"
	"crypto/sha1"
	"encoding/base64"
    "fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

func Sha1File(f string) error {
	sha1h := sha1.New()
	file, err := os.Open(f)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err = io.Copy(sha1h, file); err != nil {
		return err
	}
	log.Println("sha1:", base64.StdEncoding.EncodeToString(sha1h.Sum(nil)))
	return nil
}

func GenerateApkSignFile(apk_file, dic string) (err error) {
	var reader *zip.ReadCloser
	if reader, err = zip.OpenReader(apk_file); err != nil {
		return
	}
	defer reader.Close()
	if err = os.MkdirAll(dic, 0755); err != nil {
		return
	}
	manifest_mf := path.Join(dic, "MANIFEST.MF")
	cert_sf := path.Join(dic, "CERT.SF")

	mm_header := "Manifest-Version: 1.0\nBuilt-By: Generated-by-ADT\nCreated-By: Android Gradle 2.2.2\n\n"

	mm_sha1 := sha1.New()

	var mm_file *os.File
	var cf_file *os.File

	if mm_file, err = os.Create(manifest_mf); err != nil {
		return
	}
	if cf_file, err = os.Create(cert_sf); err != nil {
		return
	}
	defer cf_file.Close()
	mm_file.Write([]byte(mm_header))
	mm_sha1.Write([]byte(mm_header))

	for _, file := range reader.File {
		if strings.HasPrefix(file.Name, "META-INF") {
			continue
		}
		if strings.HasSuffix(file.Name, "/") {
			continue
		}
		file_name := "Name: " + file.Name
		if len(file_name) <= 70 {
			mm_file.Write([]byte(file_name + "\n"))
			mm_sha1.Write([]byte(file_name + "\n"))

		} else {
			mm_file.Write([]byte(file_name[0:70] + "\n"))
			mm_sha1.Write([]byte(file_name[0:70] + "\n"))

			mm_file.Write([]byte(" " + file_name[70:] + "\n"))
			mm_sha1.Write([]byte(" " + file_name[70:] + "\n"))

		}

		if rc, e := file.Open(); e != nil {
			err = e
			return
		} else {
			sha1h := sha1.New()
			if _, err = io.Copy(sha1h, rc); err != nil {
				return
			}
			sha1_data := base64.StdEncoding.EncodeToString(sha1h.Sum(nil))
			mm_file.Write([]byte("SHA1-Digest: " + sha1_data + "\n\n"))
			mm_sha1.Write([]byte("SHA1-Digest: " + sha1_data + "\n\n"))

			rc.Close()
		}
	}
	mm_file.Close()
	mm_file, err = os.Open(manifest_mf)
	mm_file.Seek(int64(len(mm_header)), os.SEEK_SET)
	cf_header := "Signature-Version: 1.0\nX-Android-APK-Signed: 2\nSHA1-Digest-Manifest: " + base64.StdEncoding.EncodeToString(mm_sha1.Sum(nil)) + "\nCreated-By: 1.0 (Android)\n\n"
	cf_file.Write([]byte(cf_header))
	io.Copy(cf_file, mm_file)
	mm_file.Close()
    
    fmt.Println("Export:")
    fmt.Println(manifest_mf)
    fmt.Println(cert_sf)
	return
}

func SetupApkCertCommand(rootCmd *cobra.Command) {
	var output_dir string
	cmd := &cobra.Command{
		Use:   "gen-apk-cert <apk>",
		Short: "取得apk的證書",
		Run: func(cmd *cobra.Command, args []string) {
            if len(args) != 1 {
                panic("required <apk>")
            }
            apk_fn := args[0]
			err := GenerateApkSignFile(apk_fn, output_dir)
			if err != nil {
				panic(err)
			}
		},
	}
	cmd.Flags().StringVarP(&output_dir, "output-dir", "o", ".", "輸出路徑")

	rootCmd.AddCommand(cmd)
}
