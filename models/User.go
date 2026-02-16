package models

type UserList struct {
	PageNo      int    `json:"page_no"`
	PageSize    int    `json:"page_size"`
	RecordCount int    `json:"record_count"`
	Users       []User `json:"users"`
}
type User struct {
	ID                         int     `json:"id"`
	Autotaskid                 int     `json:"autotaskid"`
	Azureoid                   string  `json:"azureoid,omitempty"`
	ClientAccountManagerID     int     `json:"client_account_manager_id"`
	ClientID                   int     `json:"client_id"`
	ClientName                 string  `json:"client_name"`
	Colour                     string  `json:"colour"`
	Connectwiseid              int     `json:"connectwiseid"`
	Email2                     string  `json:"email2"`
	Email3                     string  `json:"email3"`
	Emailaddress               string  `json:"emailaddress"`
	Firstname                  string  `json:"firstname"`
	Ignoreautomatedbilling     bool    `json:"ignoreautomatedbilling,omitempty"`
	Inactive                   bool    `json:"inactive"`
	Initials                   string  `json:"initials"`
	IsProspect                 bool    `json:"is_prospect"`
	Isimportantcontact         bool    `json:"isimportantcontact"`
	Isimportantcontact2        bool    `json:"isimportantcontact2"`
	Isserviceaccount           bool    `json:"isserviceaccount"`
	Key                        int     `json:"key"`
	LinkedAgentID              int     `json:"linked_agent_id"`
	MessagegroupID             int     `json:"messagegroup_id,omitempty"`
	Name                       string  `json:"name"`
	Neversendemails            bool    `json:"neversendemails"`
	OnpremiseActivedirectoryDn string  `json:"onpremise_activedirectory_dn,omitempty"`
	Other5                     string  `json:"other5,omitempty"`
	Overridepdftemplatequote   int     `json:"overridepdftemplatequote"`
	Phonenumber                string  `json:"phonenumber,omitempty"`
	PhonenumberPreferred       string  `json:"phonenumber_preferred,omitempty"`
	PriorityID                 int     `json:"priority_id"`
	SiteID                     float64 `json:"site_id"`
	SiteIDInt                  int     `json:"site_id_int"`
	SiteName                   string  `json:"site_name"`
	Sitephonenumber            string  `json:"sitephonenumber"`
	Sitetimezone               string  `json:"sitetimezone"`
	Surname                    string  `json:"surname"`
	Table                      int     `json:"table"`
	Telpref                    int     `json:"telpref"`
	Use                        string  `json:"use"`
}
