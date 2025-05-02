package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

type ErrorHandler struct {
	Handler http.Handler
}

func (middleware *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		fmt.Println("RECOVERED ERROR:", err)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Internal Server Error")
		}
	}()
	middleware.Handler.ServeHTTP(w, r)
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received:", r.Method, r.URL.Path)
	middleware.Handler.ServeHTTP(w, r)
	fmt.Println("Response sent")
}

func TestMiddlware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Middleware!")
	})

	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("This is a panic!")
	})
	mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
	})

	logMiddleware := LogMiddleware{
		Handler: mux,
	}

	errorHandler := ErrorHandler{
		Handler: &logMiddleware,
	}

	server := http.Server{
		Addr:    ":8080",
		Handler: &errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
