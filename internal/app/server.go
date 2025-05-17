package server

import (
	"io"
	"log"
	"math/rand"
	"net/http"
)

const (
	host            = "localhost"
	port            = "8080"
	letters         = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	shortenerLength = 8
)

var urlDatabase = make(map[string]string)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s Server) Run(addr string) error {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", shortURL)
	mux.HandleFunc("/{id}", getURL)

	log.Printf("Starting server at %s...\r\n", addr)
	return http.ListenAndServe(addr, mux)
}

// shortURL() принимает метод POST и возвращает короткую ссылку
func shortURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	URL := string(body)
	shortenedURLId := shortenerGenerator(shortenerLength)
	urlDatabase[shortenedURLId] = URL
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("http://" + host + ":" + port + "/" + shortenedURLId))
}

// shortenerGenerator() возвращает сгенерированный ключ для короткой ссылки
func shortenerGenerator(shortenerLength int) (shortURL string) {
	buffer := []byte{}
	var cursor int
	for range shortenerLength {
		cursor = (rand.Intn(len(letters)))
		buffer = append(buffer, letters[cursor])
	}
	shortURL = string(buffer)
	return
}

func getURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Path[1:]
	URL := urlDatabase[id]
	w.Header().Add("Location", URL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
