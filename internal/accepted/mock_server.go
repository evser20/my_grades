package accepted

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

// mock server
func MockServer() {
	// Создаем mock HTTP-сервер
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Обрабатываем запрос и отправляем ответ
		fmt.Fprintln(w, "Hello, client!")
	}))

	// Создаем HTTP-клиент и отправляем запрос на mock сервер
	response, err := http.Get(mockServer.URL)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer response.Body.Close()

	// Читаем ответ от сервера
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	// Выводим ответ на экран
	fmt.Println("Ответ от сервера:", string(body))

	// Закрываем mock сервер
	mockServer.Close()
}
