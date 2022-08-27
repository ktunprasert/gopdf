package domains

type Item struct {
	Code        string `json:"code,required"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	Cost        int    `json:"cost"`
}

type Customer struct {
	Name      string `json:"name"`
	Address1  string `json:"address1"`
	Address2  string `json:"address2"`
	Telephone string `json:"telephone"`
}

type Invoice struct {
	Id  string `json:"_id,required"`
	Rev string `json:"_rev,required,omitempty"`

	TenantId   string   `json:"tenant_id,required"`
	InvoiceId  string   `json:"invoice_id,required"`
	Items      []Item   `json:"items,required"`
	Type       string   `json:"type"`
	IssuedDate string   `json:"issued_date,required"`
	DueDate    string   `json:"due_date"`
	PONumber   string   `json:"po_number"`
	Customer   Customer `json:"customer"`
	Notes      string   `json:"notes"`

	ContactName      string `json:"contact_name"`
	ContactTelephone string `json:"contact_telephone"`
}
