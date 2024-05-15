package internal

import (
	"net/http"

	"github.com/adotomov/sf-sequence-api/internal/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(sh *handlers.SequenceHandler, sth *handlers.StepHandler) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(handlers.EndpointSequences, sh.GetSequences).Methods(http.MethodGet)
	router.HandleFunc(handlers.EndpointSequence, sh.GetSequenceByID).Methods(http.MethodGet)
	router.HandleFunc(handlers.EndpointSequence, sh.CreateSequence).Methods(http.MethodPost)
	router.HandleFunc(handlers.EndpointSequence, sh.UpdateSequence).Methods(http.MethodPut)
	router.HandleFunc(handlers.EndpointStep, sth.CreateStep).Methods(http.MethodPost)
	router.HandleFunc(handlers.EndpointStep, sth.UpdateStep).Methods(http.MethodPut)
	router.HandleFunc(handlers.EndpointStep, sth.DeleteStep).Methods(http.MethodDelete)
	return router
}
