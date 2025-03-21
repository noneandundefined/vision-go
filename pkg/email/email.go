package email

import (
	"log"
	"time"

	"github.com/noneandundefined/vision-go"
	"github.com/noneandundefined/vision-go/config"
	"github.com/noneandundefined/vision-go/pkg"
	"gopkg.in/gomail.v2"
)

// The private email function is
// used to send the email itself.
// Initializes the html content,
// the sender's and recipient's data.
func email(stats *vision.Vision, logs []string) {
	mail := gomail.NewMessage()
	mail.SetHeader("From", config.EMAIL_CLIENT)
	mail.SetHeader("To", config.EMAIL_CLIENT)
	mail.SetHeader("Subject", "Daily 12-hour statistics report - Vision UI")
	mail.SetBody("text/html", LoadEmailTemplate(stats))

	if len(logs) > 0 {
		for _, file := range logs {
			mail.Attach(file)
		}
	}

	d := gomail.NewDialer(config.EMAIL_SERVER, int(config.EMAIL_PORT), config.EMAIL_CLIENT, config.EMAIL_PASSWORD)
	if err := d.DialAndSend(mail); err != nil {
		log.Printf("failed to send email: %s\n", err)
		return
	}

	log.Printf("sending 12-hour statistics by email: %s\n", config.EMAIL_CLIENT)
	if err := d.DialAndSend(mail); err != nil {
		log.Printf("%s\n", err)
	}
}

// The EmailStats function is used to send a new statistics stream to the mail,
// with an initial frequency of every 12 hours.
// You can change the settings and frequency in config.
func EmailStats(stats *vision.Vision) {
	var logFiles []string
	var err error

	ticker := time.NewTicker(time.Duration(config.EMAIL_PERIOD) * time.Hour)
	defer ticker.Stop()

	gitRoot := pkg.GitRoot()
	logFiles, err = pkg.FindLogFiles(gitRoot)
	if err != nil {
		log.Println("Error find logs: %v", err)
	}

	go email(stats, logFiles)

	for range ticker.C {
		go email(stats, logFiles)
	}
}
