package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/ktunprasert/gopdf/db"
	"github.com/ktunprasert/gopdf/db/repository"
	"github.com/ktunprasert/gopdf/domains"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/ reached")
	tmpl := template.Must(template.ParseFiles("server/templates/index.html", "server/templates/base.html"))

	tenantRepo := repository.NewTenantRepository()
	if r.Method == http.MethodPost {
		tenant := &domains.Tenant{
			Id:          "tenant:" + strings.ToLower(r.FormValue("name")),
			Name:        r.FormValue("name"),
			Address1:    r.FormValue("address1"),
			Address2:    r.FormValue("address2"),
			Telephone:   r.FormValue("telephone"),
			Taxcode:     r.FormValue("taxcode"),
			BankAddress: r.FormValue("bankaddress"),
		}

		_, err := tenantRepo.Create(tenant)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	tenants, err := tenantRepo.List()
	if err != nil {
		fmt.Println(err)
		return
	}

	extractedKeys := []db.EntityKeyObject{}
	for _, tenantCompositeKey := range tenants {
		tokens := strings.Split(tenantCompositeKey, ":")
		extractedKeys = append(extractedKeys, db.EntityKeyObject{
			CompositeKey: tenantCompositeKey,
			Entity:       tokens[0],
			Id:           tokens[1],
		})
	}

	tmpl.ExecuteTemplate(w, "base", struct {
		Tenants       []string
		ExtractedKeys []db.EntityKeyObject
	}{
		Tenants:       tenants,
		ExtractedKeys: extractedKeys,
	})
}
