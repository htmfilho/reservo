package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *server) handleIndex() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Reservo!\n")
	}
}
