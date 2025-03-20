package email

import (
	"log"
	"time"

	"github.com/noneandundefined/vision-go/config"
	"gopkg.in/gomail.v2"
)

// The private email function is
// used to send the email itself.
// Initializes the html content,
// the sender's and recipient's data.
func email() {
	mail := gomail.NewMessage()
	mail.SetHeader("From", config.EMAIL_CLIENT)
	mail.SetHeader("To", config.EMAIL_CLIENT)
	mail.SetHeader("Subject", "Daily 12-hour statistics report - Vision UI")
	mail.SetBody("text/html", ``)

	d := gomail.NewDialer(config.EMAIL_SERVER, int(config.EMAIL_PORT), config.EMAIL_CLIENT, config.EMAIL_PASSWORD)
	if err := d.DialAndSend(mail); err != nil {
		log.Printf("failed to send email: %s\n", err)
		return
	}

	log.Printf("email sent: %s\n", config.EMAIL_CLIENT)
	if err := d.DialAndSend(mail); err != nil {
		log.Printf("%s\n", err)
	}
}

// The EmailStats function is used to send a new statistics stream to the mail,
// with an initial frequency of every 12 hours.
// You can change the settings and frequency in config.
func EmailStats() {
	ticker := time.NewTicker(time.Duration(config.EMAIL_PERIOD) * time.Hour)
	defer ticker.Stop()

	go email()

	for range ticker.C {
		go email()
	}
}
