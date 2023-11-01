package smpt

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendEmail(to string, subject string, token string) {
	smtpServer := "smtp.gmail.com"
	smtpPort := "587"
	smtpUsername := "nekahiwota@gmail.com"
	smtpPassword := "qmbkhyqbhgvvfinl" // Generate an App Password from your Google Account

	// to := email

	// subject := "Hello, Gopher!"
	// message := "This is the email body."

	// Create the email message
	// body := "Subject: " + subject + "\r\n" +
	// 	"\r\n" + token
	message := fmt.Sprintf(`
        <html>
        <head>
            <style>
                /* Add your CSS styles here */
                body {
                    font-family: Arial, sans-serif;
                    background-color: #f4f4f4;
                    text-align: center;
                }
                .container {
                    max-width: 600px;
                    margin: 0 auto;
                    padding: 20px;
                    background-color: #fff;
                    border-radius: 5px;
                }
                h1 {
                    color: #3498db;
                }
                p {
                    font-size: 18px;
                }
                a {
                    display: inline-block;
                    margin: 20px 0;
                    padding: 10px 20px;
                    background-color: #3498db;
                    color: #fff;
                    text-decoration: none;
                    border-radius: 5px;
                }
            </style>
        </head>
        <body>
            <div class="container">
                <h1>Email Verification</h1>
                <p>Dear user,</p>
                <p>Please click the button below to verify your email address:</p>
                <a href="http://localhost:3000/auth/verify-email/%s" target="_blank">Verify Email</a>
                <p>If you did not request this verification, please ignore this email.</p>
            </div>
        </body>
        </html>`, token)
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-type: text/html\r\n\r\n" + message)
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpServer)
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, smtpUsername, []string{to}, msg)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Email sent successfully!")
	}

}
