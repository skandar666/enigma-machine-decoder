package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func GetHttp(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("index.html")
	if err != nil {
		http.Error(w, "File not found", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if _, err := io.Copy(w, file); err != nil {
		http.Error(w, "File send error", http.StatusInternalServerError)
	}
}

func LoadHttp(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Parse error", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "File from form error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "File reading error", http.StatusInternalServerError)
		return
	}

	resultStr := service.Service(string(data))

	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("%s%s", time.Now().UTC().String(), ext)

	fileOut, err := os.Create(filename)
	if err != nil {
		http.Error(w, "File create error", http.StatusInternalServerError)
		return
	}
	defer fileOut.Close()

	if _, err := fileOut.WriteString(resultStr); err != nil {
		http.Error(w, "File recording error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	fmt.Fprintln(w, resultStr)
}
