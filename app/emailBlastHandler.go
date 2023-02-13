package app

import (
	"encoding/base64"
	"encoding/json"
	"github.com/eggysetiawan/go-email-blast/config"
	"github.com/eggysetiawan/go-email-blast/domain"
	"github.com/eggysetiawan/go-email-blast/errs"
	"github.com/eggysetiawan/go-email-blast/service"
	"net/http"
	"os"
)

type EmailBlastHandler struct {
	service service.IEmailBlastService
}

func (h EmailBlastHandler) send(w http.ResponseWriter, r *http.Request) {
	var eb domain.EmailBlast

	err := json.NewDecoder(r.Body).Decode(&eb)

	if err != nil {
		response := config.NewUnexpectedResponse(err.Error())

		config.JsonResponse(w, response.Code, response)

		return
	}

	writeToFile(eb.Attachment, eb.Filename)

	appErr := h.service.SendEmail(eb)

	if err != nil {
		response := config.NewUnexpectedResponse(appErr.Message)

		config.JsonResponse(w, response.Code, response)

		return
	}

	response := config.NewDefaultResponse()

	config.JsonResponse(w, response.Code, response)

	return
}

func writeToFile(b64 string, fileName string) *errs.Exception {

	dec, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return errs.NewUnexpectedException(err.Error())
	}

	f, err := os.Create("./tmp/" + fileName)

	if err != nil {
		return errs.NewUnexpectedException(err.Error())
	}

	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		return errs.NewUnexpectedException(err.Error())
	}

	if err := f.Sync(); err != nil {
		return errs.NewUnexpectedException(err.Error())
	}

	return nil
}
