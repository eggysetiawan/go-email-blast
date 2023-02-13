package domain

import (
	"fmt"
	"github.com/eggysetiawan/go-email-blast/logger"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

type Csv struct {
	No               string `csv:"NO"`
	Instantion       string `csv:"Asal Instansi"`
	ParticipantName  string `csv:"NAMA PESERTA"`
	ParticipantEmail string `csv:"Email RS"`
	Link             string `csv:"Link Sertifikat"`
	Training         string `csv:"Nama Pelatihan"`
	TrainingDate     string `csv:"Tanggal Pelatihan"`
	Subject          string `csv:"Subject"`
	Body             string `csv:"Body"`
	Filename         string
}

func (csv *Csv) GetBody() string {
	body := csv.Body
	body = strings.ReplaceAll(body, "<<Asal Instansi>>", csv.Instantion)
	body = strings.ReplaceAll(body, "<<NAMA PESERTA>>", csv.ParticipantName)
	body = strings.ReplaceAll(body, "<<Nama Pelatihan>>", csv.Training)
	body = strings.ReplaceAll(body, "<<Tanggal Pelatihan>>", csv.TrainingDate)

	return body
}

func (csv Csv) Download(group *sync.WaitGroup, mutex *sync.Mutex) {
	defer group.Done()

	mutex.Lock()

	defer mutex.Unlock()

	if _, err := os.Stat(csv.Filename); err == nil {
		fmt.Printf("File exists\n")
		return
	}

	//url from struct
	urlId := strings.Split(csv.Link, "/")[5]

	url := fmt.Sprintf("https://drive.google.com/uc?id=%s&export=download", urlId)

	fmt.Println("Downloading file...")

	output, err := os.Create(csv.Filename)

	defer output.Close()

	response, err := http.Get(url)

	if err != nil {
		logger.Error("Error while downloading " + url + " - " + err.Error())
		return
	}
	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)

	logger.Info("Success Downloaded " + csv.Filename)

}

func (csv *Csv) split() string {
	s := strings.Split(csv.Link, "/")
	return s[5]

}
