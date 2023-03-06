package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/ktunprasert/gopdf/domains"
)

const MAX_UPLOAD_SIZE = 1024 * 1024 * 10 // 10 MB

func HandleUpload(filename string, w http.ResponseWriter, r *http.Request) {
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

	if filename != "" {
		filenameSlice := strings.Split(fileHeader.Filename, ".")
        fileExtension := filenameSlice[len(filenameSlice)-1]

        tempFile = fmt.Sprintf("./uploads/%s.%s", filename, fileExtension)
	}
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
