package toggl

import (
	"errors"
	"net/http"
	"net/url"
)

const (
	basicAuthPassword     string = "api_token" // Defined in Toggl Reports API v2
	SummaryReportEndPoint string = "https://toggl.com/reports/api/v2/summary"
)

type reportsClient struct {
	client            *http.Client
	basicAuthPassword string
	apiToken          string
	workSpaceId       string
	userAgent         string
	url               *url.URL
}

func NewReportsClient(apiToken, workSpaceId, userAgent, endPoint string) (*reportsClient, error) {
	if len(apiToken) == 0 {
		return nil, errors.New("Missing API token")
	}
	if len(workSpaceId) == 0 {
		return nil, errors.New("Missing workspace id")
	}
	if len(userAgent) == 0 {
		return nil, errors.New("Missing user agent")
	}
	if len(endPoint) == 0 {
		return nil, errors.New("Missing end point")
	}
	url, err := url.Parse(endPoint)
	if err != nil {
		return nil, err
	}

	newReportsClient := &reportsClient{
		client:            &http.Client{},
		basicAuthPassword: basicAuthPassword,
		apiToken:          apiToken,
		workSpaceId:       workSpaceId,
		userAgent:         userAgent,
		url:               url,
	}
	return newReportsClient, nil
}

type ReportsError struct {
	Message    string `json:"message"`
	Tip        string `json:"tip"`
	StatusCode int    `json:"code"`
}

type SummaryReport struct {
	Error *ReportsError `json:"error,omitempty`
	Data  []struct {
		Title struct {
			Project string `json:"project"`
		} `json:"title"`
		Time int `json:"time"`
	} `json:"data"`
}