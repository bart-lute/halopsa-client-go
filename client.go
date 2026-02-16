package halopsa_client_go

import (
	"fmt"
	"net/url"

	"github.com/bart-lute/halopsa-client-go/models"
)

var clientGetParameters = []string{
	"order",                 // query string The name of the field to order by
	"orderdesc",             // query bool Whether to order ascending or descending
	"search",                // query string Filter by Customers like your search string
	"toplevel_id",           // query int Filter by Customers belonging to a particular top level
	"includeinactive",       // query bool Include inactive Customers in the response
	"includeactive",         // query bool Include active Customers in the response
	"include_custom_fields", // query string Comma separated list of Custom Field IDs to include in the response
	"count",                 // query int When not using pagination, the number of results to return
	"include_website",       // query bool Will include the website field against the clients in the response
}

func (c *Client) GetClients(params *map[string]string) *map[int]models.Client {

	path := fmt.Sprintf("%s/%s", apiPrefix, "client")
	var body url.Values
	var headers map[string]string

	urlValues, err := getUrlValues(params, &assetGetParameters)
	if err != nil {
		panic(err)
	}

	pagedItems, err := c.getPaginatedItems(path, urlValues, body, headers)
	if err != nil {
		panic(err)
	}

	return &pagedItems.Clients

}
