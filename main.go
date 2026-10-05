package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const ServerAddress = "https://openlibrary.org/search.json?author="

type OpenLibrarySearchResponse struct {
	NumFound         int    `json:"numFound"`
	Start            int    `json:"start"`
	NumFoundExact    bool   `json:"numFoundExact"`
	NumFoundSnake    int    `json:"num_found"`
	DocumentationURL string `json:"documentation_url"`
	Query            string `json:"q"`
	Offset           *int   `json:"offset"`
	Docs             []Book `json:"docs"`
}

type Book struct {
	AuthorKey        []string `json:"author_key,omitempty"`
	AuthorName       []string `json:"author_name,omitempty"`
	CoverEditionKey  string   `json:"cover_edition_key,omitempty"`
	CoverID          int      `json:"cover_i,omitempty"`
	EbookAccess      string   `json:"ebook_access,omitempty"`
	EditionCount     int      `json:"edition_count,omitempty"`
	FirstPublishYear int      `json:"first_publish_year,omitempty"`
	HasFulltext      bool     `json:"has_fulltext,omitempty"`
	IA               []string `json:"ia,omitempty"`
	IACollection     []string `json:"ia_collection,omitempty"`
	Key              string   `json:"key,omitempty"`
	Language         []string `json:"language,omitempty"`
	PublicScan       bool     `json:"public_scan_b,omitempty"`
	Title            string   `json:"title,omitempty"`
}

func getBookData(author string) ([]Book, error) {
	var result []Book
	apiClient := &http.Client{
		Timeout: 30 * time.Second,
	}
	fullAddress := fmt.Sprintf("%s%s", ServerAddress, author)
	request, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet, fullAddress, nil)
	if err != nil {
		return nil, err
	}
	response, err := apiClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}
	defer response.Body.Close()

	var jsonResponse = OpenLibrarySearchResponse{}
	err = json.NewDecoder(response.Body).Decode(&jsonResponse)
	if err != nil && err != io.EOF {
		print("Error the actual response body was: ", response.Body, err)
		return nil, err
	}

	result = jsonResponse.Docs
	return result, nil
}
