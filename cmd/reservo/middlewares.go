package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *server) spit(h httprouter.Handle) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		fmt.Fprint(res, "Reservo ")
		h(res, req, params)
	}
}
