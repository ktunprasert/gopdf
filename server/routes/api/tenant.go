package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

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
	fmt.Printf("%+v", tenant)
	if tenant.Id == "" {
		tenant.Id = "tenant:" + tenant.Name
	}

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

func UploadLogo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenantId := vars["id"]

	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)

	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(w, "Uploaded file is too big. Please choose a file that falls under 10 MB", http.StatusBadRequest)
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tempFile := fmt.Sprintf("./uploads/%s", fileHeader.Filename)
	dst, err := os.Create(tempFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Overwrite the entity with new logo path
	tenantRepo := repository.NewTenantRepository()
	tenant, _ := tenantRepo.Get("tenant:" + tenantId)
	tenant.Logo = fmt.Sprintf("/uploads/%s", fileHeader.Filename)

	updated, err := tenantRepo.Update(tenant)
	fmt.Println("post logo", updated, err)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		domains.JsonResponse[map[string]string]{
			Success: true,
			Message: "",
			Data: map[string]string{
				"path": tempFile,
			},
		},
	)
}
