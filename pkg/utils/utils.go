package utils

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func ConvertGBKToUTF8(s string) (string, error) {
	reader := transform.NewReader(bytes.NewReader([]byte(s)), simplifiedchinese.GBK.NewDecoder())
	d, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func DownloadDemoHTML(url string) {
	r, _ := http.Get(url)
	defer r.Body.Close()
	all, _ := io.ReadAll(r.Body)
	s, _ := ConvertGBKToUTF8(string(all))

	f, _ := os.Create("./demo.html")
	f.Write([]byte(s))
}

func ConvertEncoding(originalData []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(originalData), simplifiedchinese.GBK.NewDecoder())
	return io.ReadAll(reader)
}
