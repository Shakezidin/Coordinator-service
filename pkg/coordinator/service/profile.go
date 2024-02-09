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

func (c *CoordinatorSVC) ForgetPassword(p *cpb.ForgetPassword) (*cpb.Responce, error) {
	resp, err := c.twilio.SentTwilioOTP(p.Phone)
	if err != nil {
		return &cpb.Responce{
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
	return &cpb.Responce{
		Status:  "success",
		Message: "otp sented to phone number",
	}, nil
}

func (c *CoordinatorSVC) ForgetPasswordVerify(p *cpb.ForgetPasswordVerify) (*cpb.Responce, error) {
	redisVal := c.redis.Get(context.Background(), "phoneNo")
	if redisVal.Err() != nil {
		log.Printf("unable to get value from redis err: %v", redisVal.Err().Error())
		return &cpb.Responce{
			Status:  "failure",
			Message: "phone number mis-match",
		}, redisVal.Err()
	}
	savedPhone, err := redisVal.Result()
	if err != nil {
		log.Printf("Unable to get saved phone number from Redis err: %v", err.Error())
		return &cpb.Responce{
			Status:  "failure",
			Message: "phone number mis-match",
		}, err
	}
	if savedPhone != p.Phone {
		log.Println("Provided phone number does not match the saved phone number.")
		return &cpb.Responce{
			Status:  "failure",
			Message: "phone number mis-match",
		}, errors.New("Provided phone number does not match the saved phone number.")
	}

	resp, err := c.twilio.VerifyTwilioOTP(p.Phone, p.Otp)
	if err != nil {
		return &cpb.Responce{
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
		return &cpb.Responce{
			Status: "failure",
		}, errors.New("user not found in this number")
	}

	userid := strconv.Itoa(int(user.ID))
	token, err := utils.GenerateToken(user.Email, user.Role, userid, config.LoadConfig().SECRETKEY)
	if err != nil {
		log.Printf("unable to generate token for user %v, err: %v", user.Email, err.Error())
		return nil, errors.New("unable to generate token for user")
	}

	return &cpb.Responce{
		Status:  "Now you can create new password",
		Message: token,
	}, nil
}

func (c *CoordinatorSVC) NewPassword(p *cpb.Newpassword) (*cpb.Responce, error) {
	hashPassword, err := utils.HashPassword(p.Newpassword)
	if err != nil {
		log.Printf("unable to hash password in CoordinatorSvc() - service, err: %v", err.Error())
		return &cpb.Responce{
			Status: "fail",
		}, errors.New("error while hashing password")
	}
	id, _ := strconv.Atoi(p.Id)
	err = c.Repo.UpdatePassword(uint(id), string(hashPassword))
	if err != nil {
		fmt.Println("password updating error")
		return &cpb.Responce{
			Status: "fail",
		}, errors.New("password updating error")
	}

	return &cpb.Responce{
		Status:  "success",
		Message: "password updated",
		Id:      int64(id),
	}, nil
}

func (c *CoordinatorSVC) ViewDashBordSVC(p *cpb.View) (*cpb.DashBord, error) {
	if p.Id == 0 {
		todayStart := time.Now().Truncate(24 * time.Hour)
		todayEnd := time.Now()
		dailyIncome := c.Repo.AdminCalculateDailyIncome(todayStart, todayEnd)

		currentMonthStart := time.Date(2024, time.Now().Month(), 1, 0, 0, 0, 0, time.UTC)
		currentMonthEnd := time.Now()
		monthlyIncome := c.Repo.AdminCalculateMonthlyIncome(currentMonthStart, currentMonthEnd)

		coordinatorCount := c.Repo.CoordinatorCount()

		return &cpb.DashBord{
			Today:            int64(float64(dailyIncome) * 0.70),
			Monthly:          int64(float64(monthlyIncome) * 0.70),
			CoordinatorCount: int64(coordinatorCount),
		}, nil
	}
	todayStart := time.Now().Truncate(24 * time.Hour)
	todayEnd := time.Now()

	dailyIncome := c.Repo.CalculateDailyIncome(uint(p.Id), todayStart, todayEnd)
	currentMonthStart := time.Date(2024, time.Now().Month(), 1, 0, 0, 0, 0, time.UTC)

	currentMonthEnd := time.Now()

	monthlyIncome := c.Repo.CalculateMonthlyIncome(uint(p.Id), currentMonthStart, currentMonthEnd)

	user, _ := c.Repo.FetchUserById(uint(p.Id))

	return &cpb.DashBord{
		Wallet:  int64(user.Wallet),
		Today:   int64(float64(dailyIncome) * 0.70),
		Monthly: int64(float64(monthlyIncome) * 0.70),
	}, nil
}
