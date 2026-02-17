package halopsa_client_go

import (
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

	req := newRequest("Asset")
	req.parameters = getUrlValues(params, &assetGetParameters)

	pagedItems := c.getPaginatedItems(&req)

	return &pagedItems.Assets

}

func (c *Client) GetAllAssetTypes() *map[int]models.AssetType {

	req := newRequest("AssetType")

	response := c.doRequest(&req)
	var assetTypes []models.AssetType
	getResponseObj(response, &assetTypes)

	// Convert to Map
	assetTypesMap := make(map[int]models.AssetType)
	for _, assetType := range assetTypes {
		assetTypesMap[assetType.ID] = assetType
	}

	return &assetTypesMap
}
