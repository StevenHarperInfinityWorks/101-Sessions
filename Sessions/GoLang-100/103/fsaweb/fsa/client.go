package fsa

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

type Client struct {
	endpoint string
}

func New() Client {
	return Client{
		endpoint: DefaultEndpoint,
	}
}

const DefaultEndpoint = "http://api.ratings.food.gov.uk/Establishments/?name="

func (c Client) CountEstablishments(s string) (count int, err error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", c.endpoint+url.QueryEscape(s), nil)
	req.Header.Add("x-api-version", "2")
	resp, err := client.Do(req)
	if err != nil {
		err = errors.Wrap(err, "failed to call ratings api")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.Wrap(err, "failed to read body of ratings api")
		return
	}
	var a autoGenerated
	err = json.Unmarshal(body, &a)
	if err != nil {
		err = errors.Wrap(err, "ratings API did not return valid JSON")
		return
	}
	count = len(a.Establishments)
	return
}

type autoGenerated struct {
	Establishments []struct {
		FHRSID                   int    `json:"FHRSID"`
		LocalAuthorityBusinessID string `json:"LocalAuthorityBusinessID"`
		BusinessName             string `json:"BusinessName"`
	} `json:"establishments"`
}