package goOmdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	baseUrl = "http://www.omdbapi.com/"
)

type client struct {
	apikey string
}

// Function creating Client with apikey
func NewClient(api string) *client {
	return &client{apikey: api}
}

type Rating struct {
	Source string
	Value  string
}

type OmdbTitle struct {
	Title    string
	Year     string
	Rated    string
	Released string
	Runtime  string
	Genre    string
	Director string
	Writer   string
	Actors   string
	Plot     string
	Country  string
	Awards   string
	Poster   string
	Ratings  []Rating
}

func (c *client) getByTitle(query string) string {
	return baseUrl +
		"?t=" + strings.Replace(query, " ", ".", -1) +
		"&apikey=" +
		c.apikey
}

func (c *client) getDataById(query string) string {
	return baseUrl + "?i=" + query + "&apikey=" + c.apikey
}

// Getting title for argument as client argument create with NewClient Function
func (c *client) GetDataByTitle(title string) *OmdbTitle {
	var movie OmdbTitle
	resp, err := http.Get(c.getByTitle(title))
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(data, &movie)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	}
	return &movie
}

func (m *OmdbTitle) ShowMovieInfo() string {
	return fmt.Sprintf("%s (%s): %s\n", m.Title, m.Year, m.Ratings[0].Value)
}
