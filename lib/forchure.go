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

func FetchAnimalTrivia(animal string) (io.Reader, error) {
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

func ParseAnimalTrivia(reader io.Reader) (string, error) {
	var fact Fact
	if err := json.NewDecoder(reader).Decode(&fact); err != nil {
		return "", err
	}
	return fact.Text, nil
}
