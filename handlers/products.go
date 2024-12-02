package handlers

import (
	"log"
	"microservice/data"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
	}

	// catch all
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	ps := data.GetProducts()
	err := ps.ToJSON(w)
	if err != nil {
		p.l.Println("Unable to marshal json", err)
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
	p.l.Println("Writing Products to Response")
	w.WriteHeader(http.StatusOK)
}
