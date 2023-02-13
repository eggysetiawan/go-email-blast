package app

import (
	"fmt"
	"github.com/eggysetiawan/go-email-blast/domain"
	"github.com/eggysetiawan/go-email-blast/logger"
	"github.com/eggysetiawan/go-email-blast/service"
	"github.com/gocarina/gocsv"
	"os"
)

func Console() {
	ebc := EmailBlastConsole{service.NewEmailBlastService(domain.NewEmailBlastRepositorySmtp())}

	ebc.DownloadAndSend()
}

func ReadCsv(fn string) []domain.Csv {
	clientsFile, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		//return nil, errs.NewUnexpectedException(err.Error())
		logger.Error(err.Error())

	}

	defer clientsFile.Close()

	var csv []domain.Csv

	var newCsv []domain.Csv

	err = gocsv.UnmarshalFile(clientsFile, &csv)

	if err != nil {
		//return nil, errs.NewUnexpectedException(err.Error())
		logger.Error(err.Error())
	}

	for _, c := range csv {
		fn := fmt.Sprintf("./tmp/E SERTIFIKAT_%s_%s_%s.pdf", c.Training, c.ParticipantName, c.TrainingDate)

		c.Filename = fn

		newCsv = append(newCsv, c)
	}

	return newCsv

}
