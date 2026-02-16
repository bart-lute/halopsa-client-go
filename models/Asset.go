package models

type AssetList struct {
	PageNo      int     `json:"page_no"`
	PageSize    int     `json:"page_size"`
	RecordCount int     `json:"record_count"`
	Assets      []Asset `json:"assets"`
}
type Field struct {
	ID             int    `json:"id"`
	AccessLevel    int    `json:"access_level"`
	Display        string `json:"display"`
	Lookup         int    `json:"lookup"`
	Mandatory      bool   `json:"mandatory"`
	Name           string `json:"name"`
	ParenttypeID   int    `json:"parenttype_id"`
	Showonactivity bool   `json:"showonactivity"`
	Systemuse      int    `json:"systemuse"`
	TypeinfoID     int    `json:"typeinfo_id"`
	Validate       string `json:"validate"`
	Value          string `json:"value"`
}
type Asset struct {
	ID                       int     `json:"id"`
	AddigyID                 string  `json:"addigy_id"`
	AssetTypePriority        int     `json:"asset_type_priority,omitempty"`
	AssettypeID              int     `json:"assettype_id"`
	AssettypeName            string  `json:"assettype_name"`
	AutomateID               int     `json:"automate_id"`
	AuvikDeviceID            string  `json:"auvik_device_id"`
	BusinessOwnerCabID       int     `json:"business_owner_cab_id"`
	BusinessOwnerID          int     `json:"business_owner_id"`
	BusinessOwnerName        string  `json:"business_owner_name"`
	ClientID                 int     `json:"client_id"`
	ClientName               string  `json:"client_name"`
	Colour                   string  `json:"colour"`
	ContractRef              string  `json:"contract_ref"`
	Criticality              int     `json:"criticality"`
	DattoAlternateID         int     `json:"datto_alternate_id"`
	DattoID                  string  `json:"datto_id"`
	DattoURL                 string  `json:"datto_url"`
	Defaultsequence          int     `json:"defaultsequence,omitempty"`
	Device42ID               int     `json:"device42_id"`
	DeviceNumber             int     `json:"device_number"`
	Fields                   []Field `json:"fields"`
	Icon                     string  `json:"icon,omitempty"`
	Inactive                 bool    `json:"inactive"`
	InventoryNumber          string  `json:"inventory_number"`
	IssueConsignmentLineID   int     `json:"issue_consignment_line_id"`
	ItemID                   int     `json:"item_id"`
	ItemName                 string  `json:"item_name"`
	ItemstockID              int     `json:"itemstock_id"`
	ItglueURL                string  `json:"itglue_url"`
	Key                      int     `json:"key"`
	KeyField                 string  `json:"key_field"`
	KeyField2                string  `json:"key_field2"`
	KeyField3                string  `json:"key_field3"`
	ManufacturerID           int     `json:"manufacturer_id"`
	NcentralDetailsID        int     `json:"ncentral_details_id"`
	NinjarmmID               int     `json:"ninjarmm_id"`
	NonConsignable           bool    `json:"non_consignable"`
	PassportalID             int     `json:"passportal_id"`
	PriorityID               int     `json:"priority_id"`
	ReservedSalesorderID     int     `json:"reserved_salesorder_id"`
	ReservedSalesorderLineID int     `json:"reserved_salesorder_line_id"`
	SLAID                    int     `json:"sla_id"`
	SiteID                   int     `json:"site_id"`
	SiteName                 string  `json:"site_name"`
	StatusID                 int     `json:"status_id"`
	SupplierContractID       int     `json:"supplier_contract_id"`
	SupplierID               int     `json:"supplier_id"`
	SupplierPriorityID       int     `json:"supplier_priority_id"`
	SupplierSLAID            int     `json:"supplier_sla_id"`
	Syncroid                 int     `json:"syncroid"`
	Table                    int     `json:"table"`
	TechnicalOwnerCabID      int     `json:"technical_owner_cab_id"`
	TechnicalOwnerID         int     `json:"technical_owner_id"`
	TechnicalOwnerName       string  `json:"technical_owner_name"`
	ThirdPartyID             int     `json:"third_party_id"`
	Use                      string  `json:"use"`
	Username                 string  `json:"username"`
}

// FieldValue Gets the value from a specific field
func (a *Asset) FieldValue(name string) string {
	for _, field := range a.Fields {
		if field.Name == name {
			return field.Display
		}
	}
	return ""
}
