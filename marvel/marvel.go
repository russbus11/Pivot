package marvel

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const BaseURL = "https://gateway.marvel.com:443/v1/public/"

type Website struct {
	baseURL    string
	publicKey  string
	privateKey string
	httpClient *http.Client
}

type CharHTTPResponse struct {
	Data struct {
		Offset  int    `json:"offset"`
		Limit   int    `json:"limit"`
		Total   int    `json:"total"`
		Count   int    `json:"count"`
		Results []Cast `json:"results"`
	} `json:"data"`
}

type Cast struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func marvelKeys() (string, string) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env file")
	}
	pub := os.Getenv("MARVEL_PUBLIC_KEY")
	priv := os.Getenv("MARVEL_PRIVATE_KEY")
	return pub, priv
}

func NewWebsite(url string) Website {
	var publicKey, privateKey = marvelKeys()
	var httpClient = &http.Client{
		Timeout: 10 * time.Second,
	}
	return Website{url, publicKey, privateKey, httpClient}
}

func (c *Website) getHash(t int64) string {
	ts := strconv.FormatInt(t, 10)
	hash := md5.Sum([]byte(ts + c.privateKey + c.publicKey))
	return hex.EncodeToString(hash[:])
}

func (c *Website) signURL(url string) string {
	t := time.Now().Unix()
	hash := c.getHash(t)
	return fmt.Sprintf("%s&ts=%d&apikey=%s&hash=%s", url, t, c.publicKey, hash)
}

func (c *Website) GetCharacters(l int) ([]Cast, error) {
	url := c.baseURL + fmt.Sprintf("characters?limit=%d", l)
	url = c.signURL(url)

	res, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var charResponse CharHTTPResponse
	if err = json.NewDecoder(res.Body).Decode(&charResponse); err != nil {
		return nil, err
	}

	return charResponse.Data.Results, nil
}
