package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/hello", nil)
	responseRecorder := httptest.NewRecorder()

	helloHandler(responseRecorder, request)

	response := responseRecorder.Result()
	defer response.Body.Close()

	t.Run("status code is 200 OK", func(t *testing.T) {
		if response.StatusCode != http.StatusOK {
			t.Errorf(
				"ожидался статус %d, получен %d",
				http.StatusOK,
				response.StatusCode,
			)
		}
	})

	t.Run("content type is application/json", func(t *testing.T) {
		contentType := response.Header.Get("Content-Type")

		if contentType != "application/json" {
			t.Errorf(
				"ожидался Content-Type application/json, получен %q",
				contentType,
			)
		}
	})

	t.Run("response body contains correct message", func(t *testing.T) {
		var body struct {
			Message string `json:"message"`
		}

		err := json.NewDecoder(response.Body).Decode(&body)
		if err != nil {
			t.Fatalf("не удалось прочитать JSON-ответ: %v", err)
		}

		expectedMessage := "hello, world!"

		if body.Message != expectedMessage {
			t.Errorf(
				"ожидалось сообщение %q, получено %q",
				expectedMessage,
				body.Message,
			)
		}
	})
}
