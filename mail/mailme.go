package mailme

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/tkanos/gonfig"

	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var config struct {
	SendGridAPIKey string
}

func init() {
	err := gonfig.GetConf("./config.json", &config)
	if err != nil {
		fmt.Println("Error while reading configuration file.")
		config.SendGridAPIKey = os.Getenv("SENDGRID_API_KEY")
	}
}

// SendMail will send mail to specified person with specified contents
func SendMail(toName, toEmail, subject, plainContent, htmlContent string) error {
	from := mail.NewEmail("App Domain", "appdomain@gmail.com")
	to := mail.NewEmail(toName, toEmail)
	message := mail.NewSingleEmail(from, subject, to, plainContent, htmlContent)
	client := sendgrid.NewSendClient(config.SendGridAPIKey)
	_, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("Email was sent successfully to " + toName + " on " + toEmail)
	return nil
}
