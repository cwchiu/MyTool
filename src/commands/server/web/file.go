package web

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func CreateUploadHandler(root string) http.HandlerFunc {
	upload_folder := filepath.Join(root, "upload")
	os.Mkdir(upload_folder, 0777)
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		//POST takes the uploaded file(s) and saves it to disk.
		case "POST":
			//parse the multipart form in the request
			err := r.ParseMultipartForm(100000)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			//get a ref to the parsed multipart form
			m := r.MultipartForm

			//get the *fileheaders
			files := m.File["uploadfile"]
			for i, _ := range files {
				//for each fileheader, get a handle to the actual file
				file, err := files[i].Open()
				defer file.Close()
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				//create destination file making sure the path is writeable.
				target_fn := filepath.Join(upload_folder, files[i].Filename)
				fmt.Println(filepath.Base(target_fn))
				dst, err := os.Create(target_fn)
				defer dst.Close()
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				//copy the uploaded file to the destination file
				if _, err := io.Copy(dst, file); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

			}

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func RegisterFile(root string) {
	http.HandleFunc("/file/upload", CreateUploadHandler(root))
	// http.HandleFunc("/file/download", createUploadHandler(root))

}
