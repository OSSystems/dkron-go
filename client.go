package dkron

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
)

type Client struct {
	client *http.Client

	baseURL *url.URL

	Jobs JobsService
}

func NewClient(addr string) (*Client, error) {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}

	baseURL, err := url.Parse(fmt.Sprintf("http://%s:%s", host, port))
	if err != nil {
		return nil, err
	}

	c := &Client{client: http.DefaultClient, baseURL: baseURL}
	c.Jobs = &JobsServiceOp{client: c}

	return c, nil
}

func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	u, _ := url.Parse(path)
	uri := c.baseURL.ResolveReference(u)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, uri.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	return req, nil
}

func (c *Client) Do(req *http.Request, value interface{}) (*http.Response, error) {
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(res.Body).Decode(value)
	if err != nil {
		return nil, err
	}

	return res, nil
}
