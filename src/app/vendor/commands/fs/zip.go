package fs

import (
	"archive/zip"
	"github.com/spf13/cobra"
	"io"
	"os"
)

func SetupZipCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "zip [filename]",
		Short: "zip files",
		Long:  `zip files`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				panic("need an filename")
			}
			outFile, err := os.Create(args[0])
			if err != nil {
				panic(err)
			}
			defer outFile.Close()

			zipWriter := zip.NewWriter(outFile)
			defer zipWriter.Close()

			buf := make([]byte, 32*1024)
			for _, fn := range args[1:] {
				fileWriter, err := zipWriter.Create(fn)
				if err != nil {
					panic(err)
				}

				file, err := os.Open(fn)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				for {
					n, err := file.Read(buf)

					if n > 0 {
						_, err = fileWriter.Write(buf)
						if err != nil {
							panic(err)
						}
					}

					if err == io.EOF {
						break
					}
					if err != nil {
						panic(err)
					}
				}

			}

		},
	}
	rootCmd.AddCommand(cmd)
}
