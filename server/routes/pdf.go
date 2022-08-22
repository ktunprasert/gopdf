package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ktunprasert/gopdf/gotenberg"
)

func GeneratePdf(w http.ResponseWriter, req *http.Request) {
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
