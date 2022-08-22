package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ktunprasert/gopdf/server/routes"
	"github.com/ktunprasert/gopdf/server/routes/api"
)

type Server struct{}

func (s *Server) setupRoutes(router *mux.Router) {
	router.HandleFunc("/gotenberg/", routes.GeneratePdf)

	router.HandleFunc("/", routes.Index)
}

func (s *Server) setupApiRoutes(router *mux.Router) {
	apiRouter := router.PathPrefix("/api/").Subrouter()

	apiRouter.HandleFunc("/tenant/{id}/", api.GetTenant).Methods("GET")
	apiRouter.HandleFunc("/tenant/{id}/", api.DeleteTenant).Methods("DELETE")
	apiRouter.HandleFunc("/tenant/{id}/", api.UpdateTenant).Methods("PUT")
	apiRouter.HandleFunc("/tenant/", api.ListTenant).Methods("GET")
	apiRouter.HandleFunc("/tenant/", api.CreateTenant).Methods("POST")
	apiRouter.HandleFunc("/tenant/{tenantId}/invoices/", api.ListInvoice).Methods("GET")

	apiRouter.HandleFunc("/invoice/{id}/", api.GetInvoice).Methods("GET")
	apiRouter.HandleFunc("/invoice/{id}/", api.DeleteInvoice).Methods("DELETE")
	apiRouter.HandleFunc("/invoice/{id}/", api.UpdateInvoice).Methods("PUT")
	apiRouter.HandleFunc("/invoice/", api.CreateInvoice).Methods("POST")

	apiRouter.HandleFunc("/upload/", api.HandleUpload).Methods("POST")
}

func (s *Server) setupFileServers(router *mux.Router) {
	static := http.FileServer(http.Dir("server/static/"))
	uploads := http.FileServer(http.Dir("uploads"))

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", static))
	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", uploads))
}

func (s *Server) Start() {
	router := mux.NewRouter()
	s.setupApiRoutes(router)
	s.setupRoutes(router)
	s.setupFileServers(router)

	fmt.Println("Listening at localhost:8090...")

	http.ListenAndServe(":8090", router)
}
