package custom

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "SIMAS CONTACT"
// const CONFIG_AUTH_EMAIL = os.Getenv("email_sender")
// const CONFIG_AUTH_PASSWORD = os.Getenv("pass_sender")

func SendMail(to []string, cc []string, subject, message string) error {
	CONFIG_AUTH_EMAIL := os.Getenv("email_sender")
	CONFIG_AUTH_PASSWORD := os.Getenv("pass_sender")
	body := "From: " + CONFIG_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}