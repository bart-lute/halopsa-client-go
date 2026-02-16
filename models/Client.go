package models

type ClientList struct {
	PageNo      int      `json:"page_no"`
	PageSize    int      `json:"page_size"`
	RecordCount int      `json:"record_count"`
	Clients     []Client `json:"clients"`
}

type Client struct {
	ID                          int                    `json:"id"`
	Name                        string                 `json:"name"`
	ToplevelId                  int                    `json:"toplevel_id"`
	ToplevelName                string                 `json:"toplevel_name"`
	Inactive                    bool                   `json:"inactive"`
	Colour                      string                 `json:"colour"`
	Confirmemail                int                    `json:"confirmemail"`
	Actionemail                 int                    `json:"actionemail"`
	Clearemail                  int                    `json:"clearemail"`
	MessagegroupId              int                    `json:"messagegroup_id"`
	OverrideOrgLogo             bool                   `json:"override_org_logo,omitempty"`
	OverrideOrgName             string                 `json:"override_org_name,omitempty"`
	OverrideOrgPhone            string                 `json:"override_org_phone,omitempty"`
	OverrideOrgEmail            string                 `json:"override_org_email,omitempty"`
	OverrideOrgWebsite          string                 `json:"override_org_website,omitempty"`
	OverrideOrgPortalurl        string                 `json:"override_org_portalurl,omitempty"`
	MailboxOverride             int                    `json:"mailbox_override"`
	DefaultMailboxId            int                    `json:"default_mailbox_id"`
	Calldate                    string                 `json:"calldate,omitempty"`
	ItemTaxCode                 int                    `json:"item_tax_code"`
	ServiceTaxCode              int                    `json:"service_tax_code"`
	PrepayTaxCode               int                    `json:"prepay_tax_code"`
	ContractTaxCode             int                    `json:"contract_tax_code"`
	Pritech                     int                    `json:"pritech"`
	Sectech                     int                    `json:"sectech"`
	Accountmanagertech          int                    `json:"accountmanagertech"`
	Thirdpartynhdapiurl         string                 `json:"thirdpartynhdapiurl,omitempty"`
	Xeroid                      string                 `json:"xeroid,omitempty"`
	Use                         string                 `json:"use"`
	Key                         int                    `json:"key"`
	Table                       int                    `json:"table"`
	Logo                        string                 `json:"logo,omitempty"`
	XeroTenantId                string                 `json:"xero_tenant_id"`
	Accountsid                  string                 `json:"accountsid,omitempty"`
	Excludefrominvoicesync      bool                   `json:"excludefrominvoicesync"`
	Overridepdftemplateinvoice  int                    `json:"overridepdftemplateinvoice"`
	KashflowTenantId            int                    `json:"kashflow_tenant_id,omitempty"`
	ClientToInvoice             int                    `json:"client_to_invoice"`
	Invoiceduedaysextraclient   int                    `json:"invoiceduedaysextraclient,omitempty"`
	ItglueId                    string                 `json:"itglue_id"`
	Clientcurrency              string                 `json:"clientcurrency,omitempty"`
	SentinelSubscriptionId      string                 `json:"sentinel_subscription_id"`
	SentinelWorkspaceName       string                 `json:"sentinel_workspace_name"`
	SentinelResourceGroupName   string                 `json:"sentinel_resource_group_name"`
	DefaultCurrencyCode         int                    `json:"default_currency_code"`
	ClientToInvoiceRecurring    int                    `json:"client_to_invoice_recurring"`
	Cautomateid                 int                    `json:"cautomateid,omitempty"`
	DbcCompanyId                string                 `json:"dbc_company_id"`
	Stopped                     int                    `json:"stopped"`
	Customertype                int                    `json:"customertype"`
	CustomerRelationship        []CustomerRelationship `json:"customer_relationship"`
	SentinelDefaultUserOverride int                    `json:"sentinel_default_user_override,omitempty"`
	Ref                         string                 `json:"ref,omitempty"`
	TicketInvoicesForEachSite   bool                   `json:"ticket_invoices_for_each_site"`
	IsVip                       bool                   `json:"is_vip"`
	Taxable                     bool                   `json:"taxable"`
	PercentageToSurvey          int                    `json:"percentage_to_survey"`
	Overridepdftemplatequote    int                    `json:"overridepdftemplatequote"`
	IsAccount                   bool                   `json:"is_account"`
	QboCompanyId                string                 `json:"qbo_company_id,omitempty"`
	CustomerRelationshipList    string                 `json:"customer_relationship_list,omitempty"`
	ServicenowValidated         bool                   `json:"servicenow_validated,omitempty"`
	JiraValidated               bool                   `json:"jira_validated,omitempty"`
	Notes                       string                 `json:"notes,omitempty"`
}

type CustomerRelationship struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
