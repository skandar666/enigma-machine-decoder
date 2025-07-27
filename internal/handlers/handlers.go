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

func DataReceive(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("index.html")

	if err != nil {
		http.Error(w, "File not found", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(data)
}

func DataLoad(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Parse error", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "File from form error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	dataFile, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "File reading error", http.StatusInternalServerError)
		return
	}

	resultStr := service.Service(string(dataFile))

	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("%s%s", time.Now().Format("01022006_150405"), ext)

	fileOut, err := os.Create(filename)
	if err != nil {
		http.Error(w, "File create error", http.StatusInternalServerError)
		return
	}
	defer fileOut.Close()

	_, err = fileOut.WriteString(resultStr)
	if err != nil {
		http.Error(w, "File recording error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(resultStr))
}
