package main

import (
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/pandakn/go-devops-lab/pkg/logger"
)

func main() {
	logger := logger.New()

	logger.Info("starting server on :8080")

	http.HandleFunc("/", getHello)
	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		logger.Info("server closed")
	} else if err != nil {
		logger.Error("error starting server: %s", err)
		os.Exit(1)
	}
}

func getHello(w http.ResponseWriter, r *http.Request) {
	logger := logger.New()

	name := r.URL.Query().Get("name")
	if _, err := io.WriteString(w, "Hello "+name+"\n"); err != nil {
		logger.Error("error writing response: %s", err)
		http.Error(w, "error writing response", http.StatusInternalServerError)
		return
	}

	logger.Debug("response sent successfully to %s", r.RemoteAddr)
}
