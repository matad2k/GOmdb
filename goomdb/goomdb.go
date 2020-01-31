package goomdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	baseUrl      = "http://www.omdbapi.com/"
	movieByTitle = 0
	movieByID    = 1
	serachMovie  = 2
)

type client struct {
	apikey string
}

// Function creating Client with provided api-key
func NewClient(api string) (*client, error) {
	if api == "" {
		return &client{}, errors.New("No Api key provided")
	}
	return &client{apikey: api}, nil
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

// Function for generating query string for api
func (c *client) generateQueryString(query string, mode uint) string {
	var u *url.URL
	u, _ = url.Parse(baseUrl)
	params := url.Values{}
	params.Add("apikey", c.apikey)
	switch mode {
	case movieByTitle:
		params.Add("t", query)
	case movieByID:
		params.Add("i", query)
	case serachMovie:
		params.Add("s", query)
	default:
		params.Add("t", query)

	}
	u.RawQuery = params.Encode()
	return u.String()
}

// Getting title for argument as client argument create with NewClient Function
func (c *client) GetDataByTitle(title string) *OmdbTitle {
	var movie OmdbTitle
	resp, err := http.Get(c.generateQueryString(title, movieByTitle))
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

func (c *client) GetDataById(id string) *OmdbTitle {
	var movie OmdbTitle
	resp, err := http.Get(c.generateQueryString(id, movieByID))
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

func (m *OmdbTitle) MovieInfo() string {
	return fmt.Sprintf("%s (%s) %s IMDB, %s RT \n ", m.Title, m.Year, m.Ratings[0].Value, m.Ratings[1].Value)
}
