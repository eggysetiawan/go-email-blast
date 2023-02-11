package app

import (
	"fmt"
	"github.com/eggysetiawan/go-email-blast/domain"
	"github.com/eggysetiawan/go-email-blast/logger"
	"github.com/gocarina/gocsv"
	"os"
	"sync"
)

func Console() {
	ReadCsv("./tmp/test_csv")
}

func ReadCsv(fn string) {
	clientsFile, err := os.OpenFile("./tmp/test_csv.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		//return nil, errs.NewUnexpectedException(err.Error())
		logger.Error(err.Error())

	}

	defer clientsFile.Close()

	var csv []domain.Csv

	err = gocsv.UnmarshalFile(clientsFile, &csv)

	if err != nil {
		//return nil, errs.NewUnexpectedException(err.Error())
		logger.Error(err.Error())
	}

	group := &sync.WaitGroup{}

	mutex := &sync.Mutex{}

	var t = 0

	for _, client := range csv {
		group.Add(1)
		go func(client domain.Csv) {
			defer group.Done()
			mutex.Lock()
			fmt.Println("Hello", client)
			t++
			mutex.Unlock()
		}(client)

	}
	group.Wait()

	fmt.Println("Total ", t)

}
