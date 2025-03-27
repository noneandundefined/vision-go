package vemail

import (
	"log"
	"time"

	"github.com/noneandundefined/vision-go"
	"github.com/noneandundefined/vision-go/helpers"
	"github.com/noneandundefined/vision-go/pkg"
	"github.com/noneandundefined/vision-go/vconfig"
	"gopkg.in/gomail.v2"
)

// The private email function is
// used to send the email itself.
// Initializes the html content,
// the sender's and recipient's data.
func email(stats *vision.Vision, logs []string) {
	mail := gomail.NewMessage()
	mail.SetHeader("From", string(helpers.RestoreBytes([]byte(vconfig.EMAIL_VISIONUI), vconfig.EMAIL_VISIONUI_INDX[:])))
	mail.SetHeader("To", vconfig.EMAIL_CLIENT)
	mail.SetHeader("Subject", "Daily 12-hour statistics report - Vision UI")
	mail.SetBody("text/html", LoadEmailTemplate(stats))

	if len(logs) > 0 {
		for _, file := range logs {
			mail.Attach(file)
		}
	}

	d := gomail.NewDialer(vconfig.EMAIL_SERVER, int(vconfig.EMAIL_PORT), string(helpers.RestoreBytes([]byte(vconfig.EMAIL_VISIONUI), vconfig.EMAIL_VISIONUI_INDX[:])), string(helpers.RestoreBytes([]byte(vconfig.EMAIL_PASSWD_VISIONUI), vconfig.EMAIL_PASSWD_VISIONUI_INDX[:])))
	if err := d.DialAndSend(mail); err != nil {
		log.Printf("failed to send email: %s\n", err)
		return
	}

	log.Printf("sending 12-hour statistics by email: %s\n", vconfig.EMAIL_CLIENT)
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

	ticker := time.NewTicker(time.Duration(vconfig.EMAIL_PERIOD) * time.Hour)
	defer ticker.Stop()

	// If config is enabled, attach log files to the email.
	if vconfig.ATTACH_LOGFILES {
		gitRoot := pkg.GitRoot()
		logFiles, err = pkg.FindLogFiles(gitRoot, vconfig.LOGFILES_BY_TIME_STYLES)
		if err != nil {
			log.Printf("Error find logs: %v\n", err)
		}
	}

	for range ticker.C {
		go email(stats, logFiles)
	}
}
