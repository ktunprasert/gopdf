package repository

import (
	"fmt"

	"github.com/ktunprasert/gopdf/db"
	"github.com/ktunprasert/gopdf/domains"
)

const (
	TENANT_INVOICE_LIST = "_partition/%s/_all_docs"
)

type InvoiceRepository struct {
	Client db.ClientInterface
}

type InvoiceRepositoryInterface interface {
	Create(invoice *domains.Invoice) (*domains.Invoice, error)
	Update(invoice *domains.Invoice) (*domains.Invoice, error)
	Get(id string) (*domains.Invoice, error)
	List(tenantId string) ([]string, error)
	Delete(id string) error
}

func NewInvoiceRepository() InvoiceRepositoryInterface {
	return &InvoiceRepository{
		db.New(""),
	}
}

func (repo *InvoiceRepository) Get(id string) (*domains.Invoice, error) {
	var invoice *domains.Invoice

	err := repo.Client.Fetch(id, &invoice)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}

func (repo *InvoiceRepository) List(tenantId string) ([]string, error) {
	var partitionResponse *db.PartitionResponse

    fmt.Println(fmt.Sprintf(TENANT_INVOICE_LIST, tenantId))
	err := repo.Client.Fetch(fmt.Sprintf(TENANT_INVOICE_LIST, tenantId), &partitionResponse)
	if err != nil {
		return nil, err
	}

	idArray := make([]string, 0)
	for _, v := range partitionResponse.Rows {
		idArray = append(idArray, v.Id)
	}

	return idArray, nil
}

func (repo *InvoiceRepository) Create(invoice *domains.Invoice) (*domains.Invoice, error) {
	var res map[string]interface{}
	err := repo.Client.Create(invoice.Id, invoice, &res)
	if err != nil {
		return nil, err
	}

	invoice.Rev = res["rev"].(string)
	return invoice, nil
}

func (repo *InvoiceRepository) Update(invoice *domains.Invoice) (*domains.Invoice, error) {
	var res map[string]interface{}
	err := repo.Client.Create(invoice.Id, invoice, &res)
	if err != nil {
		return nil, err
	}

	invoice.Rev = res["rev"].(string)
	return invoice, nil
}

func (repo *InvoiceRepository) Delete(id string) error {
	t, err := repo.Get(id)
	if err != nil {
		return err
	}

	err = repo.Client.Delete(fmt.Sprintf("%s?rev=%s", id, t.Rev), nil)
	if err != nil {
		return err
	}
	return nil
}
