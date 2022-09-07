package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/ktunprasert/gopdf/db"
	"github.com/ktunprasert/gopdf/db/repository"
)

var (
	index_templates = []string{
		"server/templates/index.html",
		"server/templates/base.html",
		"server/templates/forms/tenant.html",
	}
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/ reached")
	tmpl := template.Must(template.ParseFiles(index_templates...))

	tenantRepo := repository.NewTenantRepository()

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
