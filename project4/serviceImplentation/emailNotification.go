package serviceimplentation

import "fmt"

//  "net/htp"
//  "net/http
// "net/smtp"
//  services "poject6/Notification/services"

type EmailNotification struct {
}

func (email *EmailNotification) Notify() {
	// Sender data.
	// from := "from@gmail.com"
	// password := <Email Password>"

	// // Receiver emal address.
	// to := []string{
	// 	sender@example.com",
	// }

	// / smtp server configuration.

	// smtpHost := "smtpgmail.com"
	// smtpPort := "587"

	// // Message.
	// message := []byte("This is a test email message.")

	// // Authentication.
	// auth := smtp.PlainAuth("", from, password, smtpHost)

	// // Sending email.
	// err := smtp.SenMail(smtpHost+":"+smtpPort, auth, from, to, message)
	// if err != nil {
	// 	fmt.Prntln(err)
	// 	eturn
	// }
	fmt.Println("Email Sent Successfully!")

}
