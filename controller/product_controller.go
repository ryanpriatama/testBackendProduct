package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductController interface {
	Create(writer http.ResponseWriter, request *http.Request, Params httprouter.Params)
	GetProductsSorted(writer http.ResponseWriter, request *http.Request, Params httprouter.Params)
}
