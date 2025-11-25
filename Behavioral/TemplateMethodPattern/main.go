package main

import "fmt"

type EmailTemplate interface {
	getSubject() string
	getBody() string
}

type EmailSender struct {
	Recipient string
	Template  EmailTemplate
}

func (e *EmailSender) SendEmail() {
	subject := e.Template.getSubject()
	body := e.Template.getBody()

	fmt.Printf("Sending Email to %s\nSubject: %s\nBody: %s\n\n", e.Recipient, subject, body)
}

// ============
// Concrete implementations
// ============

type WelcomeEmail struct {}

func (w *WelcomeEmail) getSubject() string {
	return "Welcome to our website!"
}

func (w *WelcomeEmail) getBody() string {
	return "Thanks for registering. We're excited to have you on board."
}

type UpdateEmail struct {}

func (u *UpdateEmail) getSubject() string {
	return  "Update on your order"
}

func (u *UpdateEmail) getBody() string {
	return "We've received your order and will process it shortly."
}

// ============
// demo app
// ============

func main() {
	welcome := &EmailSender{
		Recipient: "user@example.com",
		Template: &WelcomeEmail{},
	}

	welcome.SendEmail()

	update := &EmailSender{
		Recipient: "user@example.com",
		Template: &UpdateEmail{},
	}

	update.SendEmail()
}