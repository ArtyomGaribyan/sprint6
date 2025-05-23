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

func HandleMain(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "ошибка парсинга формы", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "ошибка получения файла", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "ошибка чтения файла", http.StatusInternalServerError)
		return
	}

	res := service.Reverse(string(data))

	timeNow := time.Now().UTC().Format("02-01-2006_15-04-05")
	fileName := fmt.Sprintf("files/localFile_%s%s", timeNow, filepath.Ext(header.Filename))

	newFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		http.Error(w, "ошибка создания файла", http.StatusInternalServerError)
	}
	defer newFile.Close()

	_, err = fmt.Fprintln(newFile, res)
	if err != nil {
		http.Error(w, "ошибка записи в файл", http.StatusInternalServerError)
	}
}
