package halopsa_client_go

import (
	"github.com/bart-lute/halopsa-client-go/models"
)

var lookupGetParameters = []string{
	"access_control_level",  // integer
	"assettype_id",          // integer
	"client_id",             // integer
	"clientname",            // string
	"contract_id",           // integer
	"country_code_id",       // integer
	"dbc_company_id",        // string
	"domain",                // string
	"exclude_nocharge",      // boolean
	"exclude_nolinkedtypes", // boolean
	"exclude_zero",          // boolean
	"iscustomfield",         // boolean
	"istree",                // boolean
	"lookupid",              // integer
	"ordervaluetype",        // integer
	"outcome_id",            // integer
	"showallcodes",          // boolean
	"ticket_id",             // integer
	"unameaprestriction",    // boolean
	"use",                   // integer
	"use2",                  // integer
}

func (c *Client) GetLookupParameters() *[]string {
	return &lookupGetParameters
}

func (c *Client) GetAllLookups() *[]models.Lookup {
	req := newRequest("Lookup")

	response := c.doRequest(&req)

	var lookups []models.Lookup
	getResponseObj(response, &lookups)

	return &lookups
}

func (c *Client) GetLookups(parameters *map[string]string) *[]models.Lookup {
	req := newRequest("Lookup")
	req.parameters = getUrlValues(parameters, &lookupGetParameters)

	response := c.doRequest(&req)
	var lookups []models.Lookup
	getResponseObj(response, &lookups)

	return &lookups
}
