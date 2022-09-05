package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ktunprasert/gopdf/db/repository"
	"github.com/ktunprasert/gopdf/domains"
)

func GetInvoice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	invoiceId := vars["id"]
	repo := repository.NewInvoiceRepository()
	invoice, err := repo.Get(invoiceId)

	if err != nil {
		handleRepoError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		domains.JsonResponse[*domains.Invoice]{
			Success: true,
			Message: "",
			Data:    invoice,
		},
	)
}

func CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var invoice *domains.Invoice
	json.NewDecoder(r.Body).Decode(&invoice)

	repo := repository.NewInvoiceRepository()
	invoice, err := repo.Create(invoice)
	if err != nil {
		handleRepoError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		domains.JsonResponse[*domains.Invoice]{
			Success: true,
			Message: "",
			Data:    invoice,
		},
	)
}

func ListInvoice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenantId := vars["tenantId"]

	repo := repository.NewInvoiceRepository()
	invoices, err := repo.List(tenantId)
	if err != nil {
		handleRepoError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		domains.JsonResponse[[]string]{
			Success: true,
			Message: "",
			Data:    invoices,
		},
	)
}

func DeleteInvoice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	invoiceId := vars["id"]
	repo := repository.NewInvoiceRepository()
	err := repo.Delete(invoiceId)
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

func UpdateInvoice(w http.ResponseWriter, r *http.Request) {
	var invoice *domains.Invoice
	json.NewDecoder(r.Body).Decode(&invoice)

	repo := repository.NewInvoiceRepository()

	fetchedInvoice, err := repo.Get(invoice.Id)
	if err != nil {
		handleRepoError(err, w)
		return
	}

	invoice.Rev = fetchedInvoice.Rev

	invoice, err = repo.Update(invoice)
	if err != nil {
		handleRepoError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		domains.JsonResponse[*domains.Invoice]{
			Success: true,
			Message: "",
			Data:    invoice,
		},
	)
}
