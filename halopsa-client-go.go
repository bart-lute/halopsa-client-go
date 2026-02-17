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

type requestObject struct {
	method     string
	path       string
	parameters *url.Values
	body       *url.Values
	headers    map[string]string
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

func getResponseObj(response *http.Response, obj any) {

	// make sure to close the body on exit or err
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, obj)
	if err != nil {
		panic(err)
	}
}

// doRequest used to make a request to the HaloPSA API
// Note: a map is always passed as a reference
func (c *Client) doRequest(req *requestObject) *http.Response {

	if req.headers == nil {
		req.headers = make(map[string]string)
	}

	// Add the API Key Header
	req.headers["X-Halo-Api-Key"] = c.ApiKey

	u, err := url.Parse(fmt.Sprintf("%s/%s", c.BaseURL, req.path))
	if err != nil {
		panic(err)
	}
	if req.parameters != nil {
		q := u.Query()
		for key, values := range *req.parameters {
			for _, value := range values {
				q.Add(key, value)
			}
		}
		u.RawQuery = q.Encode()
	}

	// The body to send with the request (can be nil)
	var bodyReader io.Reader
	if req.body != nil {
		bodyReader = strings.NewReader(req.body.Encode())
	}

	slog.Debug(fmt.Sprintf("Sending %s request to %s", req.method, u.String()))
	request, err := http.NewRequest(req.method, u.String(), bodyReader)
	if err != nil {
		panic(err)
	}

	if req.headers != nil {
		for key, value := range req.headers {
			request.Header.Set(key, value)
		}
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		panic(err)
	}

	slog.Debug(fmt.Sprintf("Got response: %s", response.Status))
	if response.StatusCode != 200 {
		panic(fmt.Errorf("unexpected status code: %d", response.StatusCode))
	}

	return response
}

func newRequest(path string) requestObject {
	return requestObject{
		path:       fmt.Sprintf("%s/%s", apiPrefix, path),
		method:     http.MethodGet,
		parameters: nil,
		body:       nil,
		headers:    nil,
	}
}

func getUrlValues(params *map[string]string, validParams *[]string) *url.Values {
	for key := range *params {
		if !slices.Contains(*validParams, key) {
			panic(fmt.Errorf("invalid parameter: %s", key))
		}
	}

	urlValues := url.Values{}
	for key, value := range *params {
		urlValues.Set(key, value)
	}
	return &urlValues
}

func injectPaginationValues(values *url.Values) {
	// Default values
	values.Set("pageinate", "true")
	values.Set("page_size", "50")
	values.Set("page_no", "1")
}

func (c *Client) getPaginatedItems(req *requestObject) *genericPagedItems {

	injectPaginationValues(req.parameters)

	pagedItems := &genericPagedItems{
		Assets:  make(map[int]models.Asset),
		Clients: make(map[int]models.Client),
		Users:   make(map[int]models.User),
	}

	for {
		response := c.doRequest(req)

		var gp genericPage
		getResponseObj(response, &gp)

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
		req.parameters.Set("page_no", fmt.Sprintf("%d", gp.PageNo+1))
	}

	return pagedItems

}
