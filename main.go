package main

import (
	"github.com/go-mail/mail"
)

func main() {

	m := mail.NewMessage()

	m.SetHeader("From", "rahmat.setiawan@lawencon.com")

	m.SetHeader("To", "setiawaneggy@gmail.com", "rahmat.shopee2@gmail.com", "ramadayantitiara@gmail.com")

	// m.SetAddressHeader("Cc", "oliver.doe@example.com", "Oliver")

	m.SetHeader("Subject", "Email Blast")

	m.SetBody("text/html", "Hello <b>Tiara</b> ini testing email blast ya!")

	// m.Attach("lolcat.jpg")

	d := mail.NewDialer("srv42.niagahoster.com", 465, "rahmat.setiawan@lawencon.com", "blacksettings2")

	// Send the email to Kate, Noah and Oliver.

	if err := d.DialAndSend(m); err != nil {

		panic(err)

	}

}
