//go:build !typescript
// +build !typescript

package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spacesprotocol/explorer-backend/cmd/rest/actions"
	"github.com/spacesprotocol/explorer-backend/pkg/db"

	_ "github.com/lib/pq"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	pg, err := sql.Open("postgres", os.Getenv("POSTGRES_URI"))
	if err != nil {
		log.Fatalln(err)
	}
	q := db.New(pg)

	handlers := make(map[string]http.HandlerFunc, 0)
	for path, function := range routes {
		handlers[path] = actions.NewAction(function).BuildHandlerFunc(q)
	}

	srv := &http.Server{
		Addr: os.Getenv("REST_ADDR"),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Access-Control-Allow-Origin", "*")
			if r.Method != "GET" {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			path := r.URL.Path
			if path == "/" {
				w.WriteHeader(http.StatusOK)
				return
			}
			if handler, ok := handlers[path]; ok {
				handler(w, r)
				return
			}
			w.WriteHeader(http.StatusNotFound)
		}),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalln("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}