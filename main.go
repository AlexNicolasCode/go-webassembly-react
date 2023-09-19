package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const dir = "./src/presentation/wasm/static/"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{file}", fileHandle).Methods(http.MethodGet)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)
	fmt.Printf("Server running")
	http.ListenAndServe(":8000", handler)
}

func fileHandle(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Cache-Control", "no-cache")
	if strings.HasSuffix(req.URL.Path, ".wasm") {
		file := mux.Vars(req)["file"]
		res.Header().Set("content-type", "application/wasm")
		data, _ := os.ReadFile(dir + file)
		res.WriteHeader(http.StatusOK)
		res.Write(data)
	}
}
