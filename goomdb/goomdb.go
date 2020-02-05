package goomdb

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

var (
	NoApiError = errors.New("NO API KEY PROVIDED")
)

// Function creating Client with provided api-key
func NewClient(api string) (*client, error) {

	// Checking if website is avalible if not returns error
	_, err := net.Dial("tcp", "www.omdbapi.com:80")
	if err != nil {
		return &client{}, err
	}

	// Checking on aplication level if apikey is empty if yes return NoApiError
	if api == "" {
		return &client{}, NoApiError
	}

	return &client{apikey: api}, nil
}

// Structure for OmdbTitle Rating table
type Rating struct {
	Source string
	Value  string
}

// Strcture for JSON object from OMDB API for GetDataByTitle and GetDataBy
type OmdbTitle struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Country    string
	Awards     string
	Poster     string
	Ratings    []Rating
	Metascore  string
	ImdbRating string
	ImdbVotes  string
	ImdbID     string
	Type       string
	BoxOffice  string
	Production string
	Website    string
	Response   string
	Error      string
}

// Function for generating query string for api
func (c *client) generateQueryString(query string, mode uint) string {
	var u *url.URL
	u, _ = url.Parse(baseUrl)
	params := url.Values{}
	params.Add("apikey", c.apikey)
	switch mode {
	case movieByTitle:
		params.Add("t", strings.ReplaceAll(query, " ", "."))
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
func (c *client) GetDataByTitle(title string) (*OmdbTitle, error) {
	return c.extractJsonData(title, movieByTitle)
}

func (c *client) GetDataById(id string) (*OmdbTitle, error) {
	return c.extractJsonData(id, movieByID)
}

func (c *client) extractJsonData(id string, mode uint) (*OmdbTitle, error) {
	var movie OmdbTitle
	resp, err := http.Get(c.generateQueryString(id, mode))
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
	postiveResponse, err := strconv.ParseBool(movie.Response)
	if err != nil {
		log.Fatal(err)
	}
	if postiveResponse == true {
		return &movie, nil
	} else {
		return &movie, errors.New("Server response:" + movie.Error)
	}
}

func (m *OmdbTitle) MovieInfo() io.Writer {
	return bytes.NewBufferString(
		fmt.Sprintf("%s (%s) %s IMDB, %s RT \n",
			m.Title,
			m.Year,
			m.Ratings[0].Value,
			m.Ratings[1].Value),
	)
}
