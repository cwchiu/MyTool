package fs

import (
	"archive/zip"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
)

func SetupUnzipCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "unzip [filename]",
		Short: "unzip file",
		Long:  `unzip file`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("need an filename")
			}
			zipReader, err := zip.OpenReader(args[0])
			if err != nil {
				panic(err)
			}
			defer zipReader.Close()

			for _, file := range zipReader.Reader.File {
				// Open the file inside the zip archive
				// like a normal file
				zippedFile, err := file.Open()
				if err != nil {
					panic(err)
				}
				defer zippedFile.Close()

				// Specify what the extracted file name should be.
				// You can specify a full path or a prefix
				// to move it to a different directory.
				// In this case, we will extract the file from
				// the zip to a file of the same name.
				targetDir := "./"
				extractedFilePath := filepath.Join(
					targetDir,
					file.Name,
				)

				// Extract the item (or create directory)
				if file.FileInfo().IsDir() {
					// Create directories to recreate directory
					// structure inside the zip archive. Also
					// preserves permissions
					// log.Println("Creating directory:", extractedFilePath)
					os.MkdirAll(extractedFilePath, file.Mode())
				} else {
					// Extract regular file since not a directory
					// log.Println("Extracting file:", file.Name)

					// Open an output file for writing
					outputFile, err := os.OpenFile(
						extractedFilePath,
						os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
						file.Mode(),
					)
					if err != nil {
						panic(err)
					}
					defer outputFile.Close()

					// "Extract" the file by copying zipped file
					// contents to the output file
					_, err = io.Copy(outputFile, zippedFile)
					if err != nil {
						panic(err)
					}
				}
			}

		},
	}
	rootCmd.AddCommand(cmd)
}
