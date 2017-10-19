package web

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
	"time"
)

func debug(data []byte, err error) {
	if err == nil {
		fmt.Printf("%s\n\n", data)
	}
}

func printDownloadPercent(done chan int64, path string, total int64) {

	var stop bool = false

	for {
		select {
		case <-done:
			stop = true
		default:

			file, err := os.Open(path)
			if err != nil {
				fmt.Println(err)
			}

			fi, err := file.Stat()
			if err != nil {
				fmt.Println(err)
			}

			size := fi.Size()

			if size == 0 {
				size = 1
			}
			fmt.Printf("%d/%d(", size, total)
			if total <= 0 {
				fmt.Printf("??")
			} else {
				var percent float64 = float64(size) / float64(total) * 100

				fmt.Printf("%.0f", percent)
			}
			fmt.Println("%)")
		}

		if stop {
			break
		}

		time.Sleep(time.Second)
	}
}

func downloadFile(url string, dest string, verbose bool) {

	start := time.Now()
	stdout := false
	out, err := os.Create(dest)

	if err != nil {
		stdout = true
	} else {
		fmt.Printf("Downloading file %s from %s\n", dest, url)

		defer out.Close()
	}

	done := make(chan int64)

	if stdout == false {
		headResp, err := http.Head(url)

		if err != nil {
			panic(err)
		}

		defer headResp.Body.Close()

		size, err := strconv.Atoi(headResp.Header.Get("Content-Length"))

		if err != nil {
			size = -1
		}

		go printDownloadPercent(done, dest, int64(size))
	}

	req, err := http.NewRequest("GET", url, nil)
	if verbose {
		debug(httputil.DumpRequestOut(req, true))
	}

	if err != nil {
		panic(err)
	}

	resp, err := (&http.Client{}).Do(req)

	if err != nil {
		panic(err)
	}

	if verbose {
		debug(httputil.DumpResponse(resp, true))
	}

	defer resp.Body.Close()

	if stdout == false {
		n, err := io.Copy(out, resp.Body)

		if err != nil {
			panic(err)
		}

		done <- n

		elapsed := time.Since(start)
		fmt.Printf("Download completed in %s", elapsed)
	} else {
		io.Copy(os.Stdout, resp.Body)
	}
}

func SetupDownloadCommand(rootCmd *cobra.Command) {
	var file string
	var verbose bool
	cmd := &cobra.Command{
		Use: "http-get <url>",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <url>")
			}
			downloadFile(args[0], file, verbose)
		},
	}
	cmd.Flags().StringVarP(&file, "file", "f", "", "儲存檔案")
	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "除錯")
	rootCmd.AddCommand(cmd)

}
