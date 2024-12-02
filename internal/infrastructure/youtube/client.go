package youtube

import (
	"io"
	"net/http"
)

type Client struct {
	*http.Client
}

func New(client *http.Client) *Client {
	return &Client{
		Client: client,
	}
}

func (c *Client) GetPreview(url string) ([]byte, error) {
	response, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	image, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return image, nil
}
