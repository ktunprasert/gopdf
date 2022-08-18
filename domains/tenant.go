package domains

type Tenant struct {
	Id  string `json:"_id,required"`
	Rev string `json:"_rev,required,omitempty"`

	Name        string `json:"name,required,omitempty"`
	Address1    string `json:"address1,omitempty"`
	Address2    string `json:"address2,omitempty"`
	Telephone   string `json:"telephone,omitempty"`
	Taxcode     string `json:"taxcode,omitempty"`
	Logo        string `json:"logo,omitempty"`
	BankAddress string `json:"bankaddress,omitempty"`
}
