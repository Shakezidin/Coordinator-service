package service

import (
	"fmt"
	"time"

	msg "github.com/Shakezidin/pkg/rabbitmq"
)

func (c *CoordinatorSVC) FetchNextDayTrip() {
	tomorrow := time.Now().Add(24 * time.Hour).Format("01-02-2006")
	bookings, err := c.Repo.FetchNextDayTrip(tomorrow)
	if err != nil {
		// Handle the error appropriately
		return
	}
	for _, booking := range *bookings {
		msg := fmt.Sprintf("Dear %s,\n\nThis is a friendly reminder that your booking for %s is scheduled for tomorrow, %s. We hope you're looking forward to your trip!\n\nIf you have any questions or need further assistance, feel free to contact us. Have a fantastic trip!\n\nBest regards,\nGlobal Package", booking.UserEmail, booking.Package.Name, tomorrow)
		sbjct := fmt.Sprintf("Reminder: Your Booking for %s Tomorrow", booking.Package.Name)
		go CreateMessage(booking.PackagePrice, booking.UserEmail, msg, sbjct)
	}
}

func CreateMessage(amount int, email, message, subject string) {
	msgs := msg.Messages{
		Email:    email,
		Amount:   amount,
		Messages: message,
		Subject:  subject,
	}
	msg.PublishConfirmationMessage(msgs)
}

func (c *CoordinatorSVC) UpdateExpiredPackage() {
	yesterday := time.Now().AddDate(0, 0, -1).Format("01-02-2006")

	err := c.Repo.UpdatePackageExpiration(yesterday)
	if err != nil {
		return
	}
}
