package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ktunprasert/gopdf/db/repository"
	"github.com/ktunprasert/gopdf/domains"
)

func GetTenant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenantId := vars["id"]
	repo := repository.NewTenantRepository()
	tenant, err := repo.Get("tenant:" + tenantId)
	if err != nil {
		handleRepoError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		domains.JsonResponse[*domains.Tenant]{
			Success: true,
			Message: "",
			Data:    tenant,
		},
	)
}

func CreateTenant(w http.ResponseWriter, r *http.Request) {
	var tenant *domains.Tenant
	json.NewDecoder(r.Body).Decode(&tenant)

	repo := repository.NewTenantRepository()
	tenant, err := repo.Create(tenant)
	if err != nil {
		handleRepoError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		domains.JsonResponse[*domains.Tenant]{
			Success: true,
			Message: "",
			Data:    tenant,
		},
	)
}

func ListTenant(w http.ResponseWriter, r *http.Request) {
	repo := repository.NewTenantRepository()
	tenants, err := repo.List()
	if err != nil {
		handleRepoError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		domains.JsonResponse[[]string]{
			Success: true,
			Message: "",
			Data:    tenants,
		},
	)
}

func DeleteTenant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenantId := vars["id"]
	repo := repository.NewTenantRepository()
	err := repo.Delete("tenant:" + tenantId)
	if err != nil {
		handleRepoError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		domains.JsonResponse[interface{}]{
			Success: true,
			Message: "",
			Data:    nil,
		},
	)
}

func UpdateTenant(w http.ResponseWriter, r *http.Request) {
	var tenant *domains.Tenant
	json.NewDecoder(r.Body).Decode(&tenant)

	repo := repository.NewTenantRepository()

	fetchedTenant, err := repo.Get(tenant.Id)
	if err != nil {
		handleRepoError(err, w)
		return
	}

	tenant.Rev = fetchedTenant.Rev

	tenant, err = repo.Update(tenant)
	if err != nil {
		handleRepoError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		domains.JsonResponse[*domains.Tenant]{
			Success: true,
			Message: "",
			Data:    tenant,
		},
	)
}
