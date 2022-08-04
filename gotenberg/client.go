package gotenberg

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

var (
	GOTENBERG_URL = "http://gotenberg:3000/forms/chromium/convert/url"
	BASE_URL      = os.Getenv("GOTENBERG_URL")
	PATH          = "/forms/chromium/convert/url"
)

type Client struct{}

func (c *Client) GetPdfStream(url string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	w, _ := writer.CreateFormField("url")
	io.Copy(w, strings.NewReader(url))
	writer.Close()

	pdfRequest, _ := http.NewRequest("POST", BASE_URL+PATH, bytes.NewReader(body.Bytes()))
	pdfRequest.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(pdfRequest)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	fileBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return fileBytes, nil
}
