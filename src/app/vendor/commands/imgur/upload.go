package imgur

import (
    "fmt"
	"github.com/spf13/cobra"
)

func SetupUploadCommand(rootCmd *cobra.Command) {
	var cid string
	cmd := &cobra.Command{
		Use:   "upload <filename>",
		Short: "上傳檔案到 Imgur",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("need an filename")
			}
            resp, err := UploadImgur(cid, args[0])
            if err != nil {
                panic(err)
            }
            
            if resp.Success != true {
                panic("upload fail")
            }
            
            fmt.Printf("Image ID: %s\n", resp.Data.HashImage)
            fmt.Printf("Delete Hash: %s\n", resp.Data.HashDelete)
            fmt.Printf("Image Original: %s\n", resp.GetImageOriginal())
            fmt.Printf("Image Large: %s\n", resp.GetImageLargeThumbnail())
            fmt.Printf("Image Small: %s\n", resp.GetImageSmallThumbnail())
            fmt.Printf("Web: %s\n", resp.GetImagePage())
            fmt.Printf("Delete: %s\n", resp.GetImageDeletePage())
		},
	}

	cmd.Flags().StringVarP(&cid, "cid", "c", "", "Client ID")
	rootCmd.AddCommand(cmd)
}
