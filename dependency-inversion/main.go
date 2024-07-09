package main

import "fmt"

// Notifier interface
type Notifier interface {
	Send(message string)
}

// EmailNotifier struct
type EmailNotifier struct{}

func (e EmailNotifier) Send(message string) {
	fmt.Println("Sending email with message:", message)
}

// SMSNotifier struct
type SMSNotifier struct{}

func (s SMSNotifier) Send(message string) {
	fmt.Println("Sending SMS with message:", message)
}

// NotificationService using Notifier
type NotificationService struct {
	notifier Notifier
}

func (n NotificationService) Notify(message string) {
	n.notifier.Send(message)
}

func main() {
	emailNotifier := EmailNotifier{}
	smsNotifier := SMSNotifier{}

	notificationService := NotificationService{notifier: emailNotifier}
	notificationService.Notify("Hello via Email")

	notificationService = NotificationService{notifier: smsNotifier}
	notificationService.Notify("Hello via SMS")
}