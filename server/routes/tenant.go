package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/ktunprasert/gopdf/db"
	"github.com/ktunprasert/gopdf/db/repository"
	"github.com/ktunprasert/gopdf/domains"
)

var (
	tenant_templates = []string{
		"server/templates/tenant.html",
		"server/templates/base.html",
		"server/templates/forms/tenant-edit.html",
	}
)

func TenantView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenantId := vars["tenantId"]

	tmpl := template.Must(template.ParseFiles(tenant_templates...))

	tenantRepo := repository.NewTenantRepository()
    tenant, _ := tenantRepo.Get("tenant:"+tenantId)
	invoiceRepo := repository.NewInvoiceRepository()
	invoices, err := invoiceRepo.List(tenantId)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(invoices)

	extractedKeys := []db.EntityKeyObject{}
	for _, invoiceCompositeKey := range invoices {
		tokens := strings.Split(invoiceCompositeKey, ":")
		extractedKeys = append(extractedKeys, db.EntityKeyObject{
			CompositeKey: invoiceCompositeKey,
			Entity:       tokens[0],
			Id:           tokens[1],
		})
	}

	tmpl.ExecuteTemplate(w, "base", struct {
		Invoices []db.EntityKeyObject
		TenantId string
		Tenant   *domains.Tenant
		Title    string
	}{
		Invoices: extractedKeys,
		TenantId: tenantId,
		Tenant:   tenant,
		Title:    tenantId,
	})
}
