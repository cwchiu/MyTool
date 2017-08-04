package imgur

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

type ImgurData struct {
	Link        string `json:"link"`
	ContentType string `json:"type"`
	HashDelete  string `json:"deletehash"`
	HashImage   string `json:"id"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
}

type ImgurResponse struct {
	Status  int       `json:"status"`
	Success bool      `json:"success"`
	Data    ImgurData `json:"data"`
}

type ImgurDeleteResponse struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func (self ImgurResponse) GetImageOriginal() string {
	return self.Data.Link
}

func (self ImgurResponse) GetImageSmallThumbnail() string {
	return fmt.Sprintf(`https://i.imgur.com/%st.jpg`, self.Data.HashImage)
}

func (self ImgurResponse) GetImagePage() string {
	return fmt.Sprintf(`https://i.imgur.com/%s`, self.Data.HashImage)
}

func (self ImgurResponse) GetImageDeletePage() string {
	return fmt.Sprintf(`https://i.imgur.com/delete/%s.jpg`, self.Data.HashDelete)
}

func (self ImgurResponse) GetImageLargeThumbnail() string {
	return fmt.Sprintf(`https://i.imgur.com/%sl.jpg`, self.Data.HashImage)
}

func DeleteImgur(client_id, delete_hash string) (*ImgurDeleteResponse, error) {
	url := fmt.Sprintf(`https://api.imgur.com/3/image/%s.json`, delete_hash)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.78 Safari/537.36")
	req.Header.Add("Authorization", fmt.Sprintf(`Client-ID %s`, client_id))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// fmt.Println(string(body))
	var ret ImgurDeleteResponse
	err2 := json.Unmarshal([]byte(body), &ret)
	if err2 != nil {
		return nil, err2
	}
	return &ret, nil
}

func UploadImgur(client_id string, filename string) (*ImgurResponse, error) {
	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)
	_, err := body_writer.CreateFormFile("image", filename)
	if err != nil {
		return nil, err
	}
	fh, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	boundary := body_writer.Boundary()
	close_buf := bytes.NewBufferString(fmt.Sprintf("\r\n--%s--\r\n", boundary))
	request_reader := io.MultiReader(body_buf, fh, close_buf)
	fi, err := fh.Stat()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "https://api.imgur.com/3/image.json", request_reader)
	if err != nil {
		return nil, err
	}

	// Set headers for multipart, and Content Length
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.78 Safari/537.36")
	req.Header.Add("Content-Type", "multipart/form-data; boundary="+boundary)
	req.Header.Add("Authorization", fmt.Sprintf(`Client-ID %s`, client_id))
	req.ContentLength = fi.Size() + int64(body_buf.Len()) + int64(close_buf.Len())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ret ImgurResponse
	err2 := json.Unmarshal([]byte(body), &ret)
	if err2 != nil {
		return nil, err2
	}
	return &ret, nil
}
