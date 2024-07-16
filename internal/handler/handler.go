package handler

import (
	"encoding/json"
	"fmt"
	"gourlshortener/internal/storage"
	"log"
	"net/http"
)

var rdb = storage.NewClient()

type RequestBody struct {
	URL string `json:"url"`
}

func SetupRoutes(router *http.ServeMux) {
	router.HandleFunc("/add-url", LongUrlHandler)
	router.HandleFunc("/ping-redis", PingRedis)
	router.HandleFunc("/all", GetAllKeys)
}

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Логирование метода и URL запроса
        log.Printf("[%s]: %s", r.Method, r.URL.String())
        // Передача управления следующему обработчику
        next.ServeHTTP(w, r)
    })
}

func PingRedis(w http.ResponseWriter, r *http.Request) {
	// Проверяем подключение
    pong, err := storage.Ping(rdb)
    if err != nil {
        log.Fatalf("Не удалось подключиться к Redis: %v", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintln(w, "Redis не доступен")
    }
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, pong)
}

func LongUrlHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        fmt.Fprintln(w, "Method Not Allowed")
        return
    }

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintln(w, "Accepted")
	var reqBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", http.StatusBadRequest)
		return
	}

	storage.AddUrl(rdb, reqBody.URL)
}

func GetAllKeys(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	keys, err := storage.GetAllKeys(rdb)
	if err != nil {
		log.Println("Error getting keys:", err)
	}
	for _, key := range keys {
		fmt.Fprintln(w, key)
	}
}