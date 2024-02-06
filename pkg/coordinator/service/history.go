package service

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/Shakezidin/pkg/coordinator/client/pb"
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
)

func (c *CoordinatorSVC) ViewhistorySVC(p *cpb.View) (*cpb.Histories, error) {
	offset := 10 * (p.Page - 1)
	limit := 10
	if p.Status == "false" {
		history, err := c.Repo.FetchHistory(int(offset), limit, uint(p.Id))
		if err != nil {
			return nil, err
		}

		var histoy []*cpb.History
		for _, hstry := range *history {
			histoy = append(histoy, &cpb.History{
				Id:              int64(hstry.ID),
				PaymentMode:     hstry.PaymentMode,
				BookingStatus:   hstry.BookingStatus,
				CancelledStatus: hstry.CancelledStatus,
				TotalPrice:      int64(hstry.PackagePrice),
				UserId:          int64(hstry.UserId),
				BookingId:       hstry.BookingId,
				BookDate:        hstry.BookDate.Format("02-01-2006"),
				StartDate:       hstry.StartDate.Format("02-01-2006"),
				PaidAmount:      int64(hstry.PaidPrice),
			})
		}

		return &cpb.Histories{
			History: histoy,
		}, nil
	}
	history, err := c.Repo.FetchBookings(int(offset), limit, uint(p.Id))
	if err != nil {
		return nil, err
	}

	var histoy []*cpb.History
	for _, hstry := range *history {
		histoy = append(histoy, &cpb.History{
			Id:              int64(hstry.ID),
			PaymentMode:     hstry.PaymentMode,
			BookingStatus:   hstry.BookingStatus,
			CancelledStatus: hstry.CancelledStatus,
			TotalPrice:      int64(hstry.PackagePrice),
			UserId:          int64(hstry.UserId),
			BookingId:       hstry.BookingId,
			BookDate:        hstry.BookDate.Format("02-01-2006"),
			StartDate:       hstry.StartDate.Format("02-01-2006"),
			PaidAmount:      int64(hstry.PaidPrice),
		})
	}

	return &cpb.Histories{
		History: histoy,
	}, nil
}

func (c *CoordinatorSVC) ViewBookingSVC(p *cpb.View) (*cpb.History, error) {
	booking, err := c.Repo.FetchBooking(uint(p.Id))
	if err != nil {
		return nil, err
	}
	var traveller []*cpb.TravellerDetails
	for _, trvler := range booking.Bookings {
		traveller = append(traveller, &cpb.TravellerDetails{
			Id:     int64(trvler.ID),
			Name:   trvler.Name,
			Age:    trvler.Age,
			Gender: trvler.Gender,
		})
	}

	return &cpb.History{
		Id:              int64(booking.ID),
		PaymentMode:     booking.PaymentMode,
		BookingStatus:   booking.BookingStatus,
		CancelledStatus: booking.CancelledStatus,
		TotalPrice:      int64(booking.PackagePrice),
		UserId:          int64(booking.UserId),
		BookingId:       booking.BookingId,
		BookDate:        booking.BookDate.Format("02-01-2006"),
		StartDate:       booking.StartDate.Format("02-01-2006"),
		Travellers:      traveller,
		PaidAmount:      int64(booking.PaidPrice),
	}, nil

}

func (c *CoordinatorSVC) CancelBookingSVC(p *cpb.View) (*cpb.Responce, error) {
	booking, err := c.Repo.FetchBooking(uint(p.Id))
	if err != nil {
		return &cpb.Responce{
			Status:  "fail",
			Message: "fetching booking error",
		}, errors.New("booking not found")
	}

	if booking.CancelledStatus == "cancelled" {
		return &cpb.Responce{
			Status: "false",
		}, errors.New("Package already cancelled")
	}

	booking.CancelledStatus = "cancelled"

	pkg, err := c.Repo.FetchPackage(booking.PackageId)
	if err != nil {
		return &cpb.Responce{
			Status:  "fail",
			Message: "fetching package error",
		}, errors.New("package not found")
	}

	pkg.MaxCapacity += len(booking.Bookings)
	coordinator, err := c.Repo.FetchUserById(pkg.CoordinatorId)
	if err != nil {
		return &cpb.Responce{
			Status: "fail",
		}, errors.New("error while fetching coordinator")
	}
	if booking.PaymentMode == "full amount" {
		coordinator.Wallet -= float64(booking.PackagePrice) * 0.70
		err = c.Repo.UpdateUser(coordinator)
	}
	err = c.Repo.UpdateBooking(*booking)
	err = c.Repo.UpdatePackage(pkg)
	var ctx = context.Background()
	_, err = c.client.AdminReduseWalletRequesr(ctx, &pb.AdminAddWallet{
		Amount: float32(booking.PackagePrice) * 0.30,
	})

	if err != nil {
		fmt.Println(err)
		return &cpb.Responce{
			Status: "fail",
		}, err
	}

	return &cpb.Responce{
		Status:  "Success",
		Message: "Package cancelled success",
	}, nil

}

func (c *CoordinatorSVC) ViewTravellerSVC(p *cpb.View) (*cpb.TravellerDetails, error) {
	traveller, err := c.Repo.FetchTraveller(uint(p.Id))
	if err != nil {
		return nil, err
	}

	activitBooking, _ := c.Repo.FetchActivityBookingofUser(uint(p.Id))

	var activity []*cpb.Activity
	for _, actvty := range activitBooking {
		activity = append(activity, &cpb.Activity{
			ActivityId:   int64(actvty.Activity.ID),
			Activityname: actvty.Activity.ActivityName,
			Description:  actvty.Activity.Description,
			Location:     actvty.Activity.Location,
			ActivityType: actvty.Activity.ActivityType,
			Amount:       int64(actvty.Activity.Amount),
			Date:         actvty.Activity.Date.Format("02-01-2006"),
			Time:         actvty.Activity.Time.Format("03:04 PM"),
		})
	}

	return &cpb.TravellerDetails{
		Name:     traveller.Name,
		Age:      traveller.Age,
		Gender:   traveller.Gender,
		Activity: activity,
	}, nil
}
