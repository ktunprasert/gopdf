package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ClientInterface interface {
	Create(path string, body, v interface{}) error
	Fetch(path string, v interface{}) error
	Delete(path string, v interface{}) error
}

type Config struct {
	host string
}

type Client struct {
	config     *Config
	httpClient *http.Client
}

func New(host string) ClientInterface {
	if host == "" {
		host = fmt.Sprintf("%s/%s/", os.Getenv("COUCHDB_URL"), os.Getenv("COUCHDB_TABLE"))
	}

	return &Client{
		config:     &Config{host},
		httpClient: &http.Client{},
	}
}

func (c *Client) Create(path string, body, v interface{}) error {
	buffer := new(bytes.Buffer)

	err := json.NewEncoder(buffer).Encode(body)
    fmt.Println(buffer)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPut, c.resolveUrl(path), buffer)
	if err != nil {
		return err
	}

	err = c.doRequest(request, v)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Fetch(path string, v interface{}) error {
	request, err := http.NewRequest(http.MethodGet, c.resolveUrl(path), nil)
	if err != nil {
		return err
	}

	err = c.doRequest(request, v)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Delete(path string, v interface{}) error {
	request, err := http.NewRequest(http.MethodDelete, c.resolveUrl(path), nil)
	if err != nil {
		return err
	}

	err = c.doRequest(request, v)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) doRequest(request *http.Request, v interface{}) error {
	request.Header.Set("Accept", "application/json")

	res, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		var restError *ErrHttp

		bytes, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		json.Unmarshal(bytes, &restError)
		restError.StatusCode = res.StatusCode

		return restError
	}

	if v != nil {
		err := json.NewDecoder(res.Body).Decode(&v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) resolveUrl(path string) string {
    fmt.Println(c.config.host + path)
	return c.config.host + path
}
