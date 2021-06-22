package main

// - Laufzeit bis Hauptreinigung: http://192.168.2.121/api/1.0/lookup/1/60/0/98/9
// - Laufzeit bis Reinigung: http://192.168.2.121/api/1.0/lookup/1/60/0/98/8
// - Betriebsstunden: http://192.168.2.121/api/1.0/lookup/1/60/0/98/4
// - Anzahl der Brennerstarts: http://192.168.2.121/api/1.0/lookup/1/60/0/98/3
// - Temperatur Abgas: http://192.168.2.121/api/1.0/lookup/1/60/0/98/1
// - Aktuelle Kesselleistung: http://192.168.2.121/api/1.0/lookup/1/60/0/98/0

// - Kesseltemperatur Istwert: http://192.168.2.121/api/1.0/lookup/1/60/0/100/1
// - Brennkammertemperatur: http://192.168.2.121/api/1.0/lookup/1/60/0/100/2
// - Aktuelle Betriebsphase: http://192.168.2.121/api/1.0/lookup/1/60/0/100/3
// - Brennstoffmenge Förderschnecke Istwert: http://192.168.2.121/api/1.0/lookup/1/60/0/100/9

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	 dac "github.com/xinsnake/go-http-digest-auth-client"
)

// Client .
type Client struct {
	baseURL    string
	username   string
	password   string
	HTTPClient *http.Client
}

// NewClient creates new client with given credentials API key
func NewClient(apiKey string) *Client {
	return &Client{
		baseURL:  "http://192.168.2.121/api/1.0/lookup",
		username: "Service",
		password: "Pmg|@03T1M{+",
		HTTPClient: &http.Client{
			Timeout: 5 * time.Minute,
		},
	}
}

// Rectangle .
type Rectangle struct {
	Top    int `json:"top"`
	Left   int `json:"left"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// Content-type and body should be already added to req
func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	// Try to unmarshall into errorResponse
	if res.StatusCode != http.StatusOK {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	// Unmarshall and populate v
	fullResponse := successResponse{
		Data: v,
	}
	if err = json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
		return err
	}

	return nil
}
