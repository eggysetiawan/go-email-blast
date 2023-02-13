package app

import (
	"github.com/eggysetiawan/go-email-blast/domain"
	"github.com/eggysetiawan/go-email-blast/service"
	"sync"
)

type EmailBlastConsole struct {
	service service.IEmailBlastService
}

func (ebc *EmailBlastConsole) DownloadAndSend() {
	participants := ReadCsv("./tmp/eb.csv")

	group := &sync.WaitGroup{}

	g := &sync.WaitGroup{}

	var mutex sync.Mutex

	var m sync.Mutex

	for _, p := range participants {
		group.Add(1)
		go p.Download(group, &mutex)

	}
	group.Wait()

	for _, p := range participants {
		g.Add(1)

		go func(p domain.Csv) {
			defer g.Done()
			m.Lock()

			eb := domain.EmailBlast{
				Name:     p.ParticipantName,
				Email:    p.ParticipantEmail,
				Subject:  p.Subject,
				Body:     p.Body,
				Filename: p.Filename,
			}
			ebc.service.SendEmail(eb)
			m.Unlock()
		}(p)

	}
	g.Wait()
}
