package app

import (
	"github.com/eggysetiawan/go-email-blast/domain"
	"github.com/eggysetiawan/go-email-blast/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Router() {
	router := mux.NewRouter()

	ebh := EmailBlastHandler{service.NewEmailBlastService(domain.NewEmailBlastRepositorySmtp())}

	router.HandleFunc("/email-blast", ebh.send).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:3001", router))

}
