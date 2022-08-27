package domains

type Tenant struct {
	Id  string `json:"_id,required"`
	Rev string `json:"_rev,required,omitempty"`

	Name            string `json:"name,required"`
	Address1        string `json:"address1"`
	Address2        string `json:"address2"`
	Telephone       string `json:"telephone"`
	Taxcode         string `json:"taxcode"`
	Logo            string `json:"logo"`
	BankAddress     string `json:"bankaddress"`
	BackgroundColor string `json:"backgroundcolor"`
	MultiplePages   bool   `json:"multiplepages"`
}
