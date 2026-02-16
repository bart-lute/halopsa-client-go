package halopsa_client_go

import (
	"net/http"
	"net/url"

	"github.com/bart-lute/halopsa-client-go/models"
)

// assetGetParameters valid parameters for Listing Assets
// Using the /Asset endpoint
// See: https://halo.halopsa.com/apidoc/resources/assets
var assetGetParameters = []string{
	"assetgroup_id",      // int
	"assettype_id",       // int
	"client_id",          // int
	"contract_id",        // int
	"includeactive",      // bool
	"includeassetfields", // bool
	"includechildren",    // bool
	"includeinactive",    // bool
	"linkedto_id",        // int
	"order",              // string
	"orderdesc",          // bool
	"search",             // string
	"site_id",            // int
	"ticket_id",          // int
	"username",           // string
}

// assetGetIdParameters valid parameters for Getting Asset by ID
// Using the /Asset/{id} endpoint
// See: https://halo.halopsa.com/apidoc/resources/assets
var assetGetIdParameters = []string{
	"includedetails",        // bool
	"includediagramdetails", // bool
}

func (c *Client) GetAssets(params *map[string]string) *map[int]models.Asset {

	path := getApiPath("asset")
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

	return &pagedItems.Assets

}

func (c *Client) GetAllAssetTypes() *map[int]models.AssetType {

	assetTypesMap := make(map[int]models.AssetType)

	path := getApiPath("AssetType")
	var body url.Values
	var headers map[string]string

	response, err := c.doRequest(http.MethodGet, path, nil, &body, headers)
	if err != nil {
		panic(err)
	}

	var assetTypes *[]models.AssetType
	err = getResponseObj(response, &assetTypes)
	if err != nil {
		panic(err)
	}

	for _, assetType := range *assetTypes {
		assetTypesMap[assetType.ID] = assetType
	}

	return &assetTypesMap
}
