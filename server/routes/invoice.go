package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ktunprasert/gopdf/db/repository"
	"github.com/ktunprasert/gopdf/domains"
)

type templatePayload struct {
	Invoice     *domains.Invoice
	TenantId    string
	Tenant      *domains.Tenant
	Items       []domains.Item
	Pages       int
	CurrentPage int
	Total       int
	Offset      int
	Raw         bool
}

func InvoiceView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenantId := vars["tenantId"]
	invoiceId := vars["invoiceId"]

	tmpl := template.Must(template.New("").Funcs(funcMap).ParseFiles("server/templates/invoice.html", "server/templates/base.html"))

	tenantRepo := repository.NewTenantRepository()
	tenant, err := tenantRepo.Get("tenant:" + tenantId)
	if err != nil {
		fmt.Println(err)
	}

	invoiceRepo := repository.NewInvoiceRepository()
	invoice, err := invoiceRepo.Get(tenantId + ":" + invoiceId)
	if err != nil {
		fmt.Println(err)
	}

	pages := (len(invoice.Items) / 15) + 1
	items := invoice.Items
	if len(items) > 15 {
		items = items[:15]
    } 

	payload := templatePayload{
		Invoice:     invoice,
		TenantId:    tenantId,
		Tenant:      tenant,
		Items:       items,
		Pages:       pages,
		CurrentPage: 1,
		Total:       0,
		Offset:      1,
		Raw:         r.URL.Query().Get("raw") == "true",
	}

	if tenant.MultiplePages {
		genMultiPageTemplates(tmpl, w, payload)
	} else {
		tmpl.ExecuteTemplate(w, "base", payload)
	}
}


func genMultiPageTemplates(tmpl *template.Template, w http.ResponseWriter, payload templatePayload) {
	tmpl.ExecuteTemplate(w, "base", payload)

	for i, n := len(payload.Invoice.Items), 1; i > 15; i -= 15 {
		var total int
		if n == payload.Pages-1 {
			for _, item := range payload.Invoice.Items[:(n * 15)] {
				total += item.Cost * item.Amount
			}
		}

		items := make([]domains.Item, 15)
		for i, n := range payload.Invoice.Items[(n * 15):] {
			items[i] = n
		}

		tmpl.ExecuteTemplate(w, "body", templatePayload{
			Invoice:     payload.Invoice,
			TenantId:    payload.TenantId,
			Tenant:      payload.Tenant,
			Items:       items,
			Pages:       payload.Pages,
			CurrentPage: n + 1,
			Total:       total,
			Offset:      (n * 15) + 1,
		})

		n++
	}
}
