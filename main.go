package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"chiapitest/routehandler"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		RequestTraceMiddleware,
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		//middleware.RequestLogger()
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		//middleware.RequestID,
		middleware.RealIP,
		middleware.Recoverer,
	)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/people", routehandler.Routes())
	})

	return router
}

func main() {
	log.SetPrefix("Hola ")
	log.Print("Starting API")
	port := ":8080"
	router := Routes()

	walkFunc := func(
		method string,
		route string,
		handler http.Handler,
		middlewares ...func(http.Handler) http.Handler) error {

		log.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}

	log.Print("API available at localhost" + port)
	log.Fatal(http.ListenAndServe(port, router))
}

const TRACE_HEADER = "X-Correlation-ID"

func RequestTraceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if w.Header().Get(TRACE_HEADER) == "" {
			if r.Header.Get(TRACE_HEADER) != "" {
				w.Header().Set(TRACE_HEADER, "traceid")
			} else {
				w.Header().Set(TRACE_HEADER, "traceid")
			}
		}

		next.ServeHTTP(w, r)
	})
}
