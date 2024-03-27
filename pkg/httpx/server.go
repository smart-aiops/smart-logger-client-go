package httpx

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func InitServer(port int, handler http.Handler) func() {
	addr := fmt.Sprintf(":%d", port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  50 * time.Second,
		WriteTimeout: 50 * time.Second,
		IdleTimeout:  50 * time.Second,
	}

	go func() {
		fmt.Printf("http server listening on: %s\n", addr)

		var err = srv.ListenAndServe()

		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(50))
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			fmt.Printf("cannot shutdown http server: %s", err)
		}

		select {
		case <-ctx.Done():
			fmt.Printf("http exiting")
		default:
			fmt.Printf("http server stopped")
		}
	}
}
