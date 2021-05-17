package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const API_ENDPOINT = "https://cat-fact.herokuapp.com/facts/random"

type Fact struct {
	Text string `json:"text"`
}

type Fetcher interface {
	requestAnimalTrivia(animal string) (io.Reader, error)
}

type Client struct {
	Fetcher Fetcher
}

func (c *Client) requestAnimalTrivia(animal string) (io.Reader, error) {
	resp, err := http.Get(fmt.Sprintf("%s?amount=1&animal_type=%s", API_ENDPOINT, animal))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var buffer bytes.Buffer
	if _, err := io.Copy(&buffer, resp.Body); err != nil {
		return nil, err
	}
	return &buffer, nil
}

func (c *Client) FetchAnimalTrivia(animal string) (string, error) {
	buffer, err := c.requestAnimalTrivia(animal)
	if err != nil {
		return "", err
	}
	var fact Fact
	if err := json.NewDecoder(buffer).Decode(&fact); err != nil {
		return "", err
	}
	return fact.Text, nil

}
