package halopsa_client_go

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"time"

	"github.com/bart-lute/halopsa-client-go/models"
)

const apiPrefix = "api"

// A page for paginated items
type genericPage struct {
	PageNo      int             `json:"page_no"`
	PageSize    int             `json:"page_size"`
	RecordCount int             `json:"record_count"`
	Assets      []models.Asset  `json:"assets"`
	Clients     []models.Client `json:"clients"`
	Users       []models.User   `json:"users"`
}

// The return value struct
type genericPagedItems struct {
	Assets  map[int]models.Asset
	Clients map[int]models.Client
	Users   map[int]models.User
}

type Client struct {
	BaseURL    string
	ApiKey     string
	httpClient *http.Client
}

// Init used to create a new client
func Init(baseUrl string, apiKey string) *Client {
	client := &Client{
		BaseURL: baseUrl,
		ApiKey:  apiKey,
		//ClientId:     clientId,
		//ClientSecret: clientSecret,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	return client
}

func getResponseObj(response *http.Response, obj any) error {

	// make sure to close the body on exit or err
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, obj)
	if err != nil {
		return err
	}
	return nil
}

// doRequest used to make a request to the Halopsa API
// Note: a map is always passed as a reference
func (c *Client) doRequest(method string, path string, parameters *url.Values, body *url.Values, headers map[string]string) (*http.Response, error) {

	if headers == nil {
		headers = map[string]string{}
	}

	// Add the API Key Header
	headers["X-Halo-Api-Key"] = c.ApiKey

	u, err := url.Parse(fmt.Sprintf("%s/%s", c.BaseURL, path))
	if err != nil {
		return nil, err
	}
	if parameters != nil {
		q := u.Query()
		for key, values := range *parameters {
			for _, value := range values {
				q.Add(key, value)
			}
		}
		u.RawQuery = q.Encode()
	}

	// The body to send with the request (can be nil)
	var bodyReader io.Reader
	if body != nil {
		bodyReader = strings.NewReader(body.Encode())
	}

	slog.Debug(fmt.Sprintf("Sending %s request to %s", method, u.String()))
	request, err := http.NewRequest(method, u.String(), bodyReader)
	if err != nil {
		return nil, err
	}

	if headers != nil {
		for key, value := range headers {
			request.Header.Set(key, value)
		}
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	slog.Debug(fmt.Sprintf("Got response: %s", response.Status))
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	return response, nil
}

func getApiPath(path string) string {
	return fmt.Sprintf("%s/%s", apiPrefix, path)
}

func getUrlValues(params *map[string]string, validParams *[]string) (*url.Values, error) {
	for key := range *params {
		if !slices.Contains(*validParams, key) {
			return nil, fmt.Errorf("invalid parameter: %s", key)
		}
	}

	urlValues := url.Values{}
	for key, value := range *params {
		urlValues.Set(key, value)
	}
	return &urlValues, nil
}

func injectPaginationValues(values *url.Values) {
	// Default values
	values.Set("pageinate", "true")
	values.Set("page_size", "50")
	values.Set("page_no", "1")
}

func (c *Client) getPaginatedItems(path string, urlValues *url.Values, body url.Values, headers map[string]string) (*genericPagedItems, error) {

	injectPaginationValues(urlValues)

	pagedItems := &genericPagedItems{
		Assets:  make(map[int]models.Asset),
		Clients: make(map[int]models.Client),
		Users:   make(map[int]models.User),
	}

	for {
		response, err := c.doRequest(http.MethodGet, path, urlValues, &body, headers)
		if err != nil {
			return nil, err
		}

		var gp genericPage
		err = getResponseObj(response, &gp)
		if err != nil {
			return nil, err
		}

		// Assets
		for _, asset := range gp.Assets {
			pagedItems.Assets[asset.ID] = asset
		}

		// Clients
		for _, client := range gp.Clients {
			pagedItems.Clients[client.ID] = client
		}

		// Users
		for _, user := range gp.Users {
			pagedItems.Users[user.ID] = user
		}

		if gp.PageNo*gp.PageSize >= gp.RecordCount {
			break
		}
		urlValues.Set("page_no", fmt.Sprintf("%d", gp.PageNo+1))
	}

	return pagedItems, nil

}
