package web

import (
	"fmt"
	"github.com/spf13/cobra"
	api "github.com/cwchiu/MyTool/libs/api/music163"
)

func setupDownloadSongCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "download-song <song-id> <filename>",
		Short: "下載音樂",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				panic("required <song-id> <filename>")
			}
			err := api.DownloadSong(args[0], args[1])
			if err != nil {
				fmt.Println(err)
			}
		},
	}
	rootCmd.AddCommand(cmd)

}

func setupEncryptCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "encrypt <data>",
		Short: "加密資料",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <data>")
			}
			fmt.Println(args[0])

			enc, err := api.Music163Encrypt(args[0])
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(enc)
			}
		},
	}
	rootCmd.AddCommand(cmd)
}

func setupDownloadProgramCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "download-program <program-id>",
		Short: "下載電台節目",
		Long:  `http://music.163.com/#/program?id=793372667<program-id>`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <data>")
			}

			err := api.DownloadRadio(args[0])
			if err != nil {
				fmt.Println(err)
			}
		},
	}
	rootCmd.AddCommand(cmd)
}

func SetupMusic163Command(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "music163",
		Short: "網易雲音樂",
	}

	setupDownloadSongCommand(cmd)
	setupDownloadProgramCommand(cmd)
	setupEncryptCommand(cmd)

	rootCmd.AddCommand(cmd)
}
