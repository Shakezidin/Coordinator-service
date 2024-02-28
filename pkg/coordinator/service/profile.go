package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Shakezidin/config"
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"github.com/Shakezidin/utils"
)

func (c *CoordinatorSVC) ForgetPassword(p *cpb.ForgetPassword) (*cpb.Response, error) {
	resp, err := c.twilio.SendTwilioOTP(p.Phone)
	if err != nil {
		return &cpb.Response{
			Status: "failure",
		}, errors.New("error while senting otp")
	} else {
		if resp.Status != nil {
			log.Println(*resp.Status)
		} else {
			log.Println(resp.Status)
		}
	}
	c.redis.Set(context.Background(), "phoneNo", p.Phone, time.Minute*2)
	return &cpb.Response{
		Status:  "success",
		Message: "otp sented to phone number",
	}, nil
}

func (c *CoordinatorSVC) ForgetPasswordVerify(p *cpb.ForgetPasswordVerify) (*cpb.Response, error) {
	redisVal := c.redis.Get(context.Background(), "phoneNo")
	if redisVal.Err() != nil {
		return &cpb.Response{
			Status:  "failure",
			Message: "phone number mis-match",
		}, redisVal.Err()
	}
	savedPhone, err := redisVal.Result()
	if err != nil {
		return &cpb.Response{
			Status:  "failure",
			Message: "phone number mis-match",
		}, err
	}
	if savedPhone != p.Phone {
		return &cpb.Response{
			Status:  "failure",
			Message: "phone number mis-match",
		}, errors.New("provided phone number does not match the saved phone number")
	}

	resp, err := c.twilio.VerifyTwilioOTP(p.Phone, p.OTP)
	if err != nil {
		return &cpb.Response{
			Status: "failure",
		}, errors.New("otp verification failed")
	} else {
		if resp.Status != nil {
			log.Println(*resp.Status)
		} else {
			log.Println(resp.Status)
		}
	}

	phoneInt, _ := strconv.Atoi(p.Phone)

	user, err := c.Repo.FindUserByPhone(phoneInt)
	if err != nil {
		fmt.Println("user not found in this number")
		return &cpb.Response{
			Status: "failure",
		}, errors.New("user not found in this number")
	}

	userid := strconv.Itoa(int(user.ID))
	token, err := utils.GenerateToken(user.Email, user.Role, userid, config.LoadConfig().SECRETKEY)
	if err != nil {
		return nil, errors.New("unable to generate token for user")
	}

	return &cpb.Response{
		Status:  "now you can create new password",
		Message: token,
	}, nil
}

func (c *CoordinatorSVC) NewPassword(p *cpb.NewPassword) (*cpb.Response, error) {
	hashPassword, err := utils.HashPassword(p.New_Password)
	if err != nil {
		return &cpb.Response{
			Status: "fail",
		}, errors.New("error while hashing password")
	}
	id, _ := strconv.Atoi(p.ID)
	err = c.Repo.UpdatePassword(uint(id), string(hashPassword))
	if err != nil {
		return &cpb.Response{
			Status: "fail",
		}, errors.New("password updating error")
	}

	return &cpb.Response{
		Status:  "success",
		Message: "password updated",
		ID:      int64(id),
	}, nil
}

func (c *CoordinatorSVC) ViewDashBordSVC(p *cpb.View) (*cpb.Dashboard, error) {
	if p.ID == 0 {
		todayStart := time.Now().Truncate(24 * time.Hour)
		todayEnd := time.Now()
		dailyIncome := c.Repo.AdminCalculateDailyIncome(todayStart, todayEnd)

		currentMonthStart := time.Date(2024, time.Now().Month(), 1, 0, 0, 0, 0, time.UTC)
		currentMonthEnd := time.Now()
		monthlyIncome := c.Repo.AdminCalculateMonthlyIncome(currentMonthStart, currentMonthEnd)

		coordinatorCount := c.Repo.CoordinatorCount()

		return &cpb.Dashboard{
			Today:             int64(float64(dailyIncome) * 0.30),
			Monthly:           int64(float64(monthlyIncome) * 0.30),
			Coordinator_Count: int64(coordinatorCount),
		}, nil
	}
	todayStart := time.Now().Truncate(24 * time.Hour)
	todayEnd := time.Now()

	dailyIncome := c.Repo.CalculateDailyIncome(uint(p.ID), todayStart, todayEnd)
	currentMonthStart := time.Date(2024, time.Now().Month(), 1, 0, 0, 0, 0, time.UTC)

	currentMonthEnd := time.Now()

	monthlyIncome := c.Repo.CalculateMonthlyIncome(uint(p.ID), currentMonthStart, currentMonthEnd)

	user, _ := c.Repo.FetchUserById(uint(p.ID))

	return &cpb.Dashboard{
		Wallet:  int64(user.Wallet),
		Today:   int64(float64(dailyIncome) * 0.70),
		Monthly: int64(float64(monthlyIncome) * 0.70),
	}, nil
}
