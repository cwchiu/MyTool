package epub

import (
	"fmt"
	"github.com/kapmahc/epub"
	"github.com/spf13/cobra"
	"net/url"
)

func SetupInfoCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "info <filename>",
		Short: "epub 資訊",
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) != 1 {
				panic("required <filename>")
			}
			bk, err := epub.Open(args[0])
			if err != nil {
				panic(err)
			}
			defer bk.Close()
			fmt.Println("== Bookmark ==")
			for _, v := range bk.Ncx.Points {
				fmt.Printf("%v : %v\n", v.Content.Src, v.Text)
			}

			fmt.Println("== Manifest ==")
			for _, manifest := range bk.Opf.Manifest {
				// fmt.Printf("id: %v", manifest.ID)
				// fmt.Printf("href: %v", manifest.Href)
				// fmt.Printf("type: %v", manifest.MediaType)
				// fmt.Printf("fallback: %v", manifest.Fallback)
				// fmt.Printf("properties: %v", manifest.Properties)
				// fmt.Printf("overlay: %v", manifest.MediaOverlay)
				name, err := url.QueryUnescape(manifest.Href)
				if err != nil {
					panic(err)
				}
				fmt.Printf("%v : %v\n", name, manifest.MediaType)
			}
		},
	}
	rootCmd.AddCommand(cmd)
}
