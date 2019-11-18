package GoOmdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	BaseUrl = "http://www.omdbapi.com/"
)

type Client struct {
	apikey string
}

func NewClient(ak string) *Client {
	return &Client{apikey: ak}
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

func getByTitle(query string, c *Client) string {
	return BaseUrl + "?t=" + query + "&apikey=" + c.apikey
}
func getDataById(query string, c *Client) string {
	return BaseUrl + "?i=" + query + "&apikey=" + c.apikey
}
func GetDataByTitle(title string, client *Client) {
	resp, err := http.Get(getByTitle(title, client))
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var movie OmdbTitle
		json.Unmarshal(data, &movie)
		ShowMovieInfo(&movie)
	}
}

func ShowMovieInfo(movie *OmdbTitle) {
	fmt.Printf("%s (%s): %s\n", movie.Title, movie.Year, movie.Ratings[0].Value)
}
