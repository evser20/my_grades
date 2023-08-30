package accepted

import (
	"log"
	"net/http"
	"path/filepath"
)

// сервер с помощью net/http
func NetHttpServer() {
	// Указываем директорию, в которой находится файл JSON
	dir := "/Users/evgeniy.ershov/Documents/data_storage/"

	// Указываем путь к файлу JSON относительно директории
	filePath := "response.json"

	// Обрабатываем запрос на получение файла JSON
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fullPath := filepath.Join(dir, filePath)
		http.ServeFile(w, r, fullPath)
	})

	// Запускаем сервер на порту 8080
	log.Fatal(http.ListenAndServe(":8091", nil))
}
