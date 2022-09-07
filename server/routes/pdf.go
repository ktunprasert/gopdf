package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ktunprasert/gopdf/gotenberg"
)

func GeneratePdf(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/gotenberg reached")
	fmt.Printf("%+v\n", r)

	tenant, invoice := r.URL.Query().Get("tenant"), r.URL.Query().Get("invoice")

	client := gotenberg.Client{}
	fileBytes, err := client.GetPdfStream(fmt.Sprintf("http://server:8090/tenant/%s/%s/?raw=true", tenant, invoice))
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
}
