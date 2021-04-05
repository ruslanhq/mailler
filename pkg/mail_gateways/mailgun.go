package mail_gateways

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v4"
	"gitlab.com/lawchad/mailler"
	"log"
	"time"
)

func MgSendEmail(name, email, htmlText string) {
	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(mailler.MgDomain, mailler.MgPrivateAPIKey)

	sender := "example@gmail.com"
	subject := "Fancy Test!"
	body := "Hello from Mailgun Go!"
	recipient := email

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject, body, recipient)

	message.SetHtml(htmlText)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}
