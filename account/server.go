package account

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	// Generate a new gorilla mux router
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	// POST method on user path that takes in our endpoint of CreateUser and use httptransport.NewServer to convert that endpoint into the appropriate type to go into this method
	// Need to create 2 functions to enable NewServer to work - DecodeRequestFunc and EncodeResponse - This is done in reqresp.go
	r.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		endpoints.CreateUser,
		decodeUserReq,
		encodeResponse,
	))

	r.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoints.GetUser,
		// Getting the ID from the path and not from a piece of JSON we need to create a unique decode function
		decodeEmailReq,
		encodeResponse,
	))

	return r
}

// Middleware to verify that all of our requests and responses will be of JSON type
func commonMiddleware(next http.Handler) http.Handler {
	// return a http.HandleFunc which takes in the ResponseWriter and Request
	// Then change the header to content type JSON, then call the next HTTP handler
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
