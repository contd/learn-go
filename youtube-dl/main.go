package main

import (
	"fmt"
	"github.com/cheggaaa/pb"
	"github.com/otium/ytdl"
	"io"
	"net/http"
	"os"
)

func main() {
	var out io.Writer
	var logOut io.Writer = os.Stdout

	info, err := ytdl.GetVideoInfo("https://www.youtube.com/watch?v=1rZ-JorHJEY")
	fileName := info.Title + ".mp4"

	formats := info.Formats
	format := formats[0]
	downloadURL, err := info.GetDownloadURL(format)

	flags := os.O_CREATE | os.O_WRONLY
	flags |= os.O_TRUNC

	var f *os.File
	// open as write only
	f, err = os.OpenFile(fileName, flags, 0666)
	if err != nil {
		err = fmt.Errorf("Unable to open output file: %s", err.Error())
		return
	}
	defer f.Close()
	out = f

	//log.Info("Downloading to ", out.(*os.File).Name())
	var req *http.Request
	req, err = http.NewRequest("GET", downloadURL.String(), nil)
	resp, err := http.DefaultClient.Do(req)

	if err != nil || resp.StatusCode < 200 || resp.StatusCode >= 300 {
		if err == nil {
			err = fmt.Errorf("Received status code %d from download url", resp.StatusCode)
		}
		err = fmt.Errorf("Unable to start download: %s", err.Error())
		return
	}

	defer resp.Body.Close()

	progressBar := pb.New64(resp.ContentLength)
	progressBar.SetUnits(pb.U_BYTES)
	progressBar.ShowTimeLeft = true
	progressBar.ShowSpeed = true
	progressBar.Output = logOut
	progressBar.Start()
	defer progressBar.Finish()
	out = io.MultiWriter(out, progressBar)

	_, err = io.Copy(out, resp.Body)

	//vid.Download(file)
}
