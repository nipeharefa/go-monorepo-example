package packages

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

type httpClient struct {
	c *http.Client
}

type HTTPClient interface {
	Get(u string) ([]byte, int, error)
	Post() ([]byte, int, error)
}

func NewHTTPClient() HTTPClient {
	h := httpClient{}
	return &h
}

//go:noinline
func (h *httpClient) Get(u string) ([]byte, int, error) {

	if h.c == nil {
		defaultTransport := &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
		}

		client := &http.Client{
			Transport: defaultTransport,
		}

		h.c = client
	}

	request, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	response, err := h.c.Do(request)
	if err != nil {
		return nil, http.StatusInternalServerError, nil
	}

	if response != nil {
		defer response.Body.Close()
	}

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, response.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, nil
	}

	return buf.Bytes(), http.StatusOK, nil
}

func (*httpClient) Post() ([]byte, int, error) {

	return nil, http.StatusInternalServerError, errors.New("error")
}
