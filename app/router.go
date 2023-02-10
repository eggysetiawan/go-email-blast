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

	emh := EmailBlastHandler{service.NewEmailBlastService(domain.NewEmailBlastRepositorySmtp())}

	router.HandleFunc("email-blast", emh.send)

	log.Fatal(http.ListenAndServe("localhost:3001", router))

}
