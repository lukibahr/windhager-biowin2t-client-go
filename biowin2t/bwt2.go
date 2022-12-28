package biowin2t

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

// WindhagerClient represents the WindhagerClient struct .
type WindhagerClient struct {
	MesEndpoint string
	MesUsername string
	MesPassword string
	HTTPClient  *http.Client
}

// NewWindhagerClient creates new client with given credentials
func NewWindhagerClient(mesurl, username, password string) *WindhagerClient {
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
	StatusCode int    `json:"code"`
	MetricBody Metric `json:"body"` // e.g. {"result: success"}
}

func (c *WindhagerClient) sendRequest(req *http.Request) (*successResponse, error) {
	var metric Metric
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
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

	response := successResponse{
		StatusCode: resp.StatusCode,
		MetricBody: Metric{Oid: metric.Oid, GroupNr: metric.GroupNr, MaxValue: metric.MaxValue, MemberNr: metric.MemberNr, MinValue: metric.MinValue, Name: metric.Name, Step: metric.Step, StepID: metric.StepID, SubtypeID: metric.SubtypeID, Timestamp: metric.Timestamp, TypeID: metric.TypeID, Unit: metric.Unit, UnitID: metric.UnitID, Value: metric.Value, WriteProt: metric.WriteProt},
	}
	return &response, nil
}

// GetTimeUntilNextMajorMaintenanceInHours returns the time until the next major maintenance in hours
func (c *WindhagerClient) GetTimeUntilNextMajorMaintenanceInHours(ctx context.Context) (*successResponse, error) {
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

// GetTimeUntilNextMaintenanceInHours returns the time until the next maintenance in hours
func (c *WindhagerClient) GetTimeUntilNextMaintenanceInHours(ctx context.Context) (*successResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1/60/0/98/8", c.MesEndpoint), nil)
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

// GetCountOfBurningUnitStarts returns the count of burning unit starts
func (c *WindhagerClient) GetCountOfBurningUnitStarts(ctx context.Context) (*successResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1/60/0/98/3", c.MesEndpoint), nil)
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

// GetExhaustGasesInCelsius returns the temperature of exhaust gases in celsius
func (c *WindhagerClient) GetExhaustGasesInCelsius(ctx context.Context) (*successResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1/60/0/98/1", c.MesEndpoint), nil)
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

// GetCurrentBoilerPerformanceInPercent returns the current boiler performance in percent
func (c *WindhagerClient) GetCurrentBoilerPerformanceInPercent(ctx context.Context) (*successResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1/60/0/98/0", c.MesEndpoint), nil)
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

// GetCurrentBoilerTemperatureInCelsius returns the current boiler temperature in celsius
func (c *WindhagerClient) GetCurrentBoilerTemperatureInCelsius(ctx context.Context) (*successResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1/60/0/100/1", c.MesEndpoint), nil)
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

// GetCombustionChamberTemperatureInCelsius returns the combustion chamber temperature in celsius
func (c *WindhagerClient) GetCombustionChamberTemperatureInCelsius(ctx context.Context) (*successResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1/60/0/100/2", c.MesEndpoint), nil)
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

// GetOperationalPhase returns the operational phase as an integer
func (c *WindhagerClient) GetOperationalPhase(ctx context.Context) (*successResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1/60/0/100/3", c.MesEndpoint), nil)
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

// GetPelletAmountOfScrewConveyor returns the pellet amount of screw conveyor in kilograms
func (c *WindhagerClient) GetPelletAmountOfScrewConveyor(ctx context.Context) (*successResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1/60/0/100/9", c.MesEndpoint), nil)
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

// GetTotalOperationalRuntimeInHours returns the total operational runtime in hours
func (c *WindhagerClient) GetTotalOperationalRuntimeInHours(ctx context.Context) (*successResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1/60/0/98/4", c.MesEndpoint), nil)
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

// TODO: marshal eror, the response does not match the defined struct. fix this

// GetDomesticHotWaterTemperatureInCelcius returns the domestic hot water temperature in celsius
func (c *WindhagerClient) GetDomesticHotWaterTemperatureInCelcius(ctx context.Context) (*successResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1/16/0/114/0", c.MesEndpoint), nil)
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

// TODO: marshal eror, the response does not match the defined struct. fix this

// GetOutsideTemperatureInCelcius returns the outside temperature in celsius
func (c *WindhagerClient) GetOutsideTemperatureInCelcius(ctx context.Context) (*successResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1/15/0/115", c.MesEndpoint), nil)
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
