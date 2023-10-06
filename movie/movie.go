package movie

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	agent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36"
)

type Movie struct {
	c *http.Client
}

func New() *Movie {
	return &Movie{
		c: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Movie) Get(url string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if _, ok := headers["User-Agent"]; !ok {
		headers["User-Agent"] = agent
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http code: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

func (c *Movie) LoadSubject(url string, headers map[string]string) (*Subject, error) {
	data, err := c.Get(url, headers)
	if err != nil {
		return nil, err
	}

	return &Subject{data: data}, nil
}

func (c *Movie) LoadCelebrity(url string, headers map[string]string) (*Celebrity, error) {
	data, err := c.Get(url, headers)
	if err != nil {
		return nil, err
	}

	return &Celebrity{data: data}, nil
}
