package halopsa_client_go

import (
	"fmt"
	"net/url"

	"github.com/bart-lute/halopsa-client-go/models"
)

// userGetParameters
// Using the /User endpoint
// See: https://halo.halopsa.com/apidoc/resources/users
var userGetParameters = []string{
	"order",                 // query string The name of the field to order by
	"orderdesc",             // query bool Whether to order ascending or descending
	"search",                // query string Filter by Users like your search
	"search_phonenumbers",   // query bool Filter by Users with a phone number like your search
	"toplevel_id",           // query int Filter by Users belonging to a particular top level
	"client_id",             // query int Filter by Users belonging to a particular customer
	"site_id",               // query int Filter by Users belonging to a particular site
	"organisation_id",       // query int Filter by Users belonging to a particular organisation
	"department_id",         // query int Filter by Users belonging to a particular department
	"asset_id",              // query int Filter by Users assigned to a particular asset
	"includeactive",         // query bool Include active Users in the response
	"includeinactive",       // query bool Include inactive Users in the response
	"include_custom_fields", // query string Comma separated list of Custom Field IDs to include in the response
	"approversonly",         // query bool Only include Users that can approve approval processes
	"excludeagents",         // query bool Exclude Users that are linked to active agent accounts
	"count",                 // query int When not using pagination, the number of results to return
	"advanced_search",       // query object array Returns tickets based on the filters in the array
}

/*
Filter Type	Description
0			IN
1			NOT IN
2			=
3			<>
4			LIKE
5			NOT LIKE
6			=
7			>
8			>=
9			<
10			<=

Example: Returns users with a first name containing the word 'Admin'
advanced_search=[{"filter_name":"firstname","filter_type":4,"filter_value":"Admin"}]
*/

func (c *Client) GetUsers(params *map[string]string) *map[int]models.User {

	path := fmt.Sprintf("%s/%s", apiPrefix, "users")

	var body url.Values
	var headers map[string]string

	urlValues, err := getUrlValues(params, &userGetParameters)
	if err != nil {
		panic(err)
	}

	pagedItems, err := c.getPaginatedItems(path, urlValues, body, headers)
	if err != nil {
		panic(err)
	}

	return &pagedItems.Users
	
}
