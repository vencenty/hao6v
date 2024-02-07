package utils

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func ConvertGBKToUTF8(s string) (string, error) {
	reader := transform.NewReader(bytes.NewReader([]byte(s)), simplifiedchinese.GBK.NewDecoder())
	d, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func DownloadDemoHTML(url string, saveAsFileName string) {
	r, _ := http.Get(url)
	defer r.Body.Close()
	all, _ := io.ReadAll(r.Body)
	s, _ := ConvertGBKToUTF8(string(all))

	f, _ := os.Create("./demo/" + saveAsFileName)
	f.Write([]byte(s))
}

func ConvertEncoding(originalData []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(originalData), simplifiedchinese.GBK.NewDecoder())
	return io.ReadAll(reader)
}

func IdentifyLinkType(url string) string {
	if strings.HasPrefix(url, "ed2k") {
		return "ed2k"
	} else if strings.HasPrefix(url, "magnet") {
		return "magnet"
	} else {
		return "other"
	}
}
