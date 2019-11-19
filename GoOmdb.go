package GoOmdb

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

func adjustTitle(title string) string {
	return strings.Replace(title, " ", ".", -1)
}

func getByTitle(query string, c *client) string {
	return baseUrl + "?t=" + adjustTitle(query) + "&apikey=" + c.apikey
}

func getDataById(query string, c *client) string {
	return baseUrl + "?i=" + query + "&apikey=" + c.apikey
}

// Getting title for argument as client argument create with NewClient Function
func GetDataByTitle(title string, c *client) *OmdbTitle {
	var movie OmdbTitle
	resp, err := http.Get(getByTitle(title, c))
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

func ShowMovieInfo(movie *OmdbTitle) {
	fmt.Printf("%s (%s): %s\n", movie.Title, movie.Year, movie.Ratings[0].Value)
}
