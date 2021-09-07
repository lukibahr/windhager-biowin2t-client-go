package windhager

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	dac "github.com/xinsnake/go-http-digest-auth-client"
)

const (
	mesurl = "http://192.168.2.140/api/1.0/lookup"
)

// WindhagerClient represents the WindhagerClient struct .
type WindhagerClient struct {
	MesEndpoint string
	MesUsername string
	MesPassword string
	HTTPClient  *http.Client
}

// NewWindhagerClient creates new client with given credentials
func NewWindhagerClient(username, password string) *WindhagerClient {
	return &WindhagerClient{
		MesEndpoint: mesurl,
		MesUsername: username,
		MesPassword: password,
		HTTPClient: &http.Client{
			Timeout: 1 * time.Minute,
		},
	}
}

type Metric struct {
	Oid       string `json:"OID"`
	GroupNr   int    `json:"groupNr"`
	MaxValue  string `json:"maxValue"`
	MemberNr  int    `json:"memberNr"`
	MinValue  string `json:"minValue"`
	Name      string `json:"name"`
	Step      string `json:"step"`
	StepID    int    `json:"stepId"`
	SubtypeID int    `json:"subtypeId"`
	Timestamp string `json:"timestamp"`
	TypeID    int    `json:"typeId"`
	Unit      string `json:"unit"`
	UnitID    int    `json:"unitId"`
	Value     string `json:"value"`
	WriteProt bool   `json:"writeProt"`
}
type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func (c *WindhagerClient) sendRequest(req *http.Request) (*Metric, error) {

	var metric Metric
	req.Header.Set("Accept", "application/json; charset=utf-8")
	t := dac.NewTransport(c.MesUsername, c.MesPassword)
	resp, err := t.RoundTrip(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(resp.Body).Decode(&errRes); err == nil {
			return nil, errors.New(errRes.Message)
		}

		return nil, fmt.Errorf("unknown error, status code: %d", resp.StatusCode)
	}
	if err = json.NewDecoder(resp.Body).Decode(&metric); err != nil {
		return nil, err
	}

	return &metric, nil
}

func (c *WindhagerClient) GetTimeUntilNextMajorMaintenanceInHours(ctx context.Context) (*Metric, error) {
	//http://192.168.2.121/api/1.0/lookup/1/60/0/98/9
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1/60/0/98/9", c.MesEndpoint), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *WindhagerClient) GetTimeUntilNextMaintenanceInHours(ctx context.Context) (*Metric, error) {
	// - Laufzeit bis Reinigung: http://192.168.2.121/api/1.0/lookup/1/60/0/98/8
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1/60/0/98/8", c.MesEndpoint), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}

	fmt.Print(res.Value)
	return res, nil
}

func (c *WindhagerClient) GetCountOfBurningUnit(ctx context.Context) (*Metric, error) {
	// - Anzahl der Brennerstarts: http://192.168.2.121/api/1.0/lookup/1/60/0/98/3
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1/60/0/98/3", c.MesEndpoint), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}

	fmt.Print(res.Value)
	return res, nil
}

func (c *WindhagerClient) GetExhaustGases(ctx context.Context) (*Metric, error) {
	// - Temperatur Abgas: http://192.168.2.121/api/1.0/lookup/1/60/0/98/1
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1/60/0/98/1", c.MesEndpoint), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}

	fmt.Print(res.Value)
	return res, nil
}

func (c *WindhagerClient) GetCurrentBoilerPerformance(ctx context.Context) (*Metric, error) {
	// - Aktuelle Kesselleistung: http://192.168.2.121/api/1.0/lookup/1/60/0/98/0
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1/60/0/98/0", c.MesEndpoint), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}

	fmt.Print(res.Value)
	return res, nil
}

func (c *WindhagerClient) GetCurrentTemperature(ctx context.Context) (*Metric, error) {
	// - Kesseltemperatur Istwert: http://192.168.2.121/api/1.0/lookup/1/60/0/100/1
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1/60/0/100/1", c.MesEndpoint), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}

	fmt.Print(res.Value)
	return res, nil
}

func (c *WindhagerClient) GetCombustorTemperature(ctx context.Context) (*Metric, error) {
	// - Brennkammertemperatur: http://192.168.2.121/api/1.0/lookup/1/60/0/100/2
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1/60/0/100/2", c.MesEndpoint), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}

	fmt.Print(res.Value)
	return res, nil
}

func (c *WindhagerClient) GetOperationalPhase(ctx context.Context) (*Metric, error) {
	// - Aktuelle Betriebsphase: http://192.168.2.121/api/1.0/lookup/1/60/0/100/3
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1/60/0/100/3", c.MesEndpoint), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}

	fmt.Print(res.Value)
	return res, nil
}

func (c *WindhagerClient) GetPelletAmountOfScrewConveyor(ctx context.Context) (*Metric, error) {
	// - Brennstoffmenge Förderschnecke Istwert: http://192.168.2.121/api/1.0/lookup/1/60/0/100/9
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1/60/0/100/9", c.MesEndpoint), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}

	fmt.Print(res.Value)
	return res, nil
}

// GetTotalOperationalRuntime returns total hours
func (c *WindhagerClient) GetTotalOperationalRuntime(ctx context.Context) (*Metric, error) {
	// - Betriebsstunden: http://192.168.2.121/api/1.0/lookup/1/60/0/98/4
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1/60/0/98/4", c.MesEndpoint), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}

	fmt.Print(res.Value)
	return res, nil
}
