package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ktunprasert/gopdf/gotenberg"
	"github.com/ktunprasert/gopdf/server/routes"
)

type Server struct{}

func (s *Server) hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("/hello reached")
	fmt.Fprintf(w, "hello world\n")
}

func (s *Server) template(w http.ResponseWriter, req *http.Request) {
	fmt.Println("/template reached")
	tmpl := template.Must(template.ParseFiles("server/templates/template.html"))
	tmpl.Execute(w, struct {
		Title string
		Todos []string
	}{
		Title: "My Title",
		Todos: []string{
			"abc",
			"def",
			"ghk",
		},
	})
}

func (s *Server) index(w http.ResponseWriter, req *http.Request) {
	fmt.Println("/ reached")
	tmpl := template.Must(template.ParseFiles("server/templates/index.html"))
	tmpl.Execute(w, struct {
		PageTitle   string
		BodyContent string
	}{
		PageTitle:   "I am a Page Title",
		BodyContent: "I am a body content that came from Golang",
	})
}

func (s *Server) generatePdf(w http.ResponseWriter, req *http.Request) {
    // Params are as follows:
    // PDF_URL
    // FILE_NAME
	fmt.Println("/gotenberg reached")
    fmt.Printf("%+v\n", req)

    client := gotenberg.Client{}
    fileBytes, err := client.GetPdfStream("https://example.com")
    if err != nil {
        log.Fatal(err)
    }

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", `attachment; filename="example.pdf"`)
	w.Write(fileBytes)
}

func (s *Server) setupRoutes(router *mux.Router) {
	router.HandleFunc("/gotenberg/", s.generatePdf)

    router.HandleFunc("/tenant/{id}/", routes.GetTenant).Methods("GET")
    router.HandleFunc("/tenant/{id}/", routes.DeleteTenant).Methods("DELETE")
    router.HandleFunc("/tenant/{id}/", routes.UpdateTenant).Methods("PUT")
    router.HandleFunc("/tenant/", routes.ListTenant).Methods("GET")
    router.HandleFunc("/tenant/", routes.CreateTenant).Methods("POST")
    router.HandleFunc("/tenant/{tenantId}/invoices/", routes.ListInvoice).Methods("GET")

    router.HandleFunc("/invoice/{id}/", routes.GetInvoice).Methods("GET")
    router.HandleFunc("/invoice/{id}/", routes.DeleteInvoice).Methods("DELETE")
    router.HandleFunc("/invoice/{id}/", routes.UpdateInvoice).Methods("PUT")
    router.HandleFunc("/invoice/", routes.CreateInvoice).Methods("POST")

    router.HandleFunc("/upload/", routes.HandleUpload).Methods("POST")
    router.HandleFunc("/", s.index)
}

func (s *Server) setupStatic() {
	fs := http.FileServer(http.Dir("server/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

    uploads := http.FileServer(http.Dir("server/uploads"))
    http.Handle("/uploads/", http.StripPrefix("/uploads/", uploads))
}

func (s *Server) Start() {
    router := mux.NewRouter()
	s.setupRoutes(router)
	s.setupStatic()
	fmt.Println("Listening at localhost:8090...")

	http.ListenAndServe(":8090", router)
}
