package libs

import (
	"fmt"
	"github.com/cavaliercoder/grab"
	"time"
)

type Job struct {
	Url      string
	Filename string
	request  *grab.Request
}

type Downloader struct {
	links map[string]Job
}

func (d Downloader) Add(url, filename string) error {
	// client := grab.NewClient()
	req, err := grab.NewRequest(".", url)
	if err != nil {
		return err
	}
	req.Filename = filename
	req.NoResume = false

	d.links[url] = Job{
		Url:      url,
		Filename: filename,
		request:  req,
	}

	return nil
}

func (d Downloader) Start() {
	reqs := make([]*grab.Request, 0)
	for _, job := range d.links {
		reqs = append(reqs, job.request)
	}

	client := grab.NewClient()
	ch_resp := client.DoBatch(4, reqs...)

	for resp := range ch_resp {
		err := resp.Err()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("%v / %v bytes (%.2f%%)\n", resp.BytesComplete(), resp.Size, 100*resp.Progress())
	}

}

func DownloadOne(url, filename string) error {
	client := grab.NewClient()
	req, err := grab.NewRequest(".", url)
	if err != nil {
		return err
	}

	if filename != "" {
		req.Filename = filename
	}
	req.NoResume = false

	resp := client.Do(req)
    fmt.Println(`CanResume: `, resp.CanResume)
    fmt.Println(`DidResume: `, resp.DidResume)
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("%v / %v bytes (%.2f%%)\n", resp.BytesComplete(), resp.Size, 100*resp.Progress())
		case <-resp.Done:
			break Loop
		}
	}

	err = resp.Err()
	if err != nil {
		return err
	}
	fmt.Println(resp.Filename)

	return nil
}
