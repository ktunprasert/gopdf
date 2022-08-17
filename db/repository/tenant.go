package repository

import (
	"fmt"

	"github.com/ktunprasert/gopdf/db"
	"github.com/ktunprasert/gopdf/domains"
)

const (
	TENANT_LIST = "_partition/tenant/_all_docs"
)

type TenantRepository struct {
	Client db.ClientInterface
}

type TenantRepositoryInterface interface {
	Create(tenant *domains.Tenant) (*domains.Tenant, error)
	Update(tenant *domains.Tenant) (*domains.Tenant, error)
	Get(id string) (*domains.Tenant, error)
	List() ([]string, error)
	Delete(id string) error
}

func NewTenantRepository() TenantRepositoryInterface {
	return &TenantRepository{
		db.New(""),
	}
}

func (repo *TenantRepository) Get(id string) (*domains.Tenant, error) {
	var tenant *domains.Tenant

	err := repo.Client.Fetch(id, &tenant)
	if err != nil {
		return nil, err
	}

	return tenant, nil
}

func (repo *TenantRepository) List() ([]string, error) {
	var partitionResponse *db.PartitionResponse

	err := repo.Client.Fetch(TENANT_LIST, &partitionResponse)
	if err != nil {
		return nil, err
	}

	idArray := make([]string, 0)
	for _, v := range partitionResponse.Rows {
		idArray = append(idArray, v.Id)
	}

	return idArray, nil
}

func (repo *TenantRepository) Create(tenant *domains.Tenant) (*domains.Tenant, error) {
	var res map[string]interface{}
	err := repo.Client.Create(tenant.Id, tenant, &res)
	if err != nil {
		return nil, err
	}

	tenant.Rev = res["rev"].(string)
	return tenant, nil
}

func (repo *TenantRepository) Update(tenant *domains.Tenant) (*domains.Tenant, error) {
	var res map[string]interface{}
	err := repo.Client.Create(tenant.Id, tenant, &res)
	if err != nil {
		return nil, err
	}

	tenant.Rev = res["rev"].(string)
	return tenant, nil
}

func (repo *TenantRepository) Delete(id string) error {
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
