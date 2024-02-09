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
		msgs := msg.Messages{
			Email:    booking.UserEmail,
			Amount:   booking.PackagePrice,
			Messages: fmt.Sprintf("Dear %s,\n\nThis is a friendly reminder that your booking for %s is scheduled for tomorrow, %s. We hope you're looking forward to your trip!\n\nIf you have any questions or need further assistance, feel free to contact us. Have a fantastic trip!\n\nBest regards,\nGlobal Package", booking.UserEmail, booking.Package.Name, tomorrow),
			Subject:  fmt.Sprintf("Reminder: Your Booking for %s Tomorrow", booking.Package.Name),
		}
		msg.PublishConfirmationMessage(msgs)
	}
}

func (c *CoordinatorSVC) UpdateExpiredPackage() {
	yesterday := time.Now().AddDate(0, 0, -1).Format("01-02-2006")

	err := c.Repo.UpdatePackageExpiration(yesterday)
	if err != nil {
		// Handle the error appropriately
		return
	}
}
