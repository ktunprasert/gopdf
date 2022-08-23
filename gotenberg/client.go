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

	urlField, _ := writer.CreateFormField("url")
	io.Copy(urlField, strings.NewReader(url))

	cssPageSizeField, _ := writer.CreateFormField("preferCssPageSize")
	io.Copy(cssPageSizeField, strings.NewReader("true"))


    margin := "0.3"
	marginTopField, _ := writer.CreateFormField("marginTop")
	io.Copy(marginTopField, strings.NewReader(margin))

	marginBottomField, _ := writer.CreateFormField("marginBottom")
	io.Copy(marginBottomField, strings.NewReader(margin))

	marginLeftField, _ := writer.CreateFormField("marginLeft")
	io.Copy(marginLeftField, strings.NewReader(margin))

	marginRightField, _ := writer.CreateFormField("marginRight")
	io.Copy(marginRightField, strings.NewReader(margin))

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
