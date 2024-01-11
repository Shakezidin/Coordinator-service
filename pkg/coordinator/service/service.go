package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Shakezidin/config"
	cDOM "github.com/Shakezidin/pkg/DOM/coordinator"
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	inter "github.com/Shakezidin/pkg/coordinator/repository/interface"
	SVCinter "github.com/Shakezidin/pkg/coordinator/service/interface"
	"github.com/Shakezidin/utils"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type CoordinatorSVC struct {
	Repo   inter.CoordinatorRepoInter
	twilio *config.TwilioVerify
	redis  *redis.Client
}

func (c *CoordinatorSVC) SignupSVC(p *cpb.Signup) (*cpb.SignupResponce, error) {
	hashPassword, err := utils.HashPassword(p.Password)
	if err != nil {
		log.Printf("unable to hash password in CoordinatorSvc() - service, err: %v", err.Error())
		return nil, err
	}
	user := &cDOM.User{
		Phone:    int(p.Phone),
		Email:    p.Email,
		Password: string(hashPassword),
		Name:     p.Name,
	}
	// send otp to phone number
	resp, err := c.twilio.SentTwilioOTP(string(p.Phone))
	if err != nil {
		return nil, err
	} else {
		if resp.Status != nil {
			log.Println(*resp.Status)
		} else {
			log.Println(resp.Status)
		}
	}
	userData, err := json.Marshal(&user)
	if err != nil {
		log.Printf("error parsing JSON, err: %v", err.Error())
		return nil, err
	}

	registerUser := fmt.Sprintf("register_user_%v", p.Email)
	c.redis.Set(context.Background(), registerUser, userData, time.Minute*2)
	return &cpb.SignupResponce{
		Status:  "success",
		Message: "user Creation initiated, check mail for OTP",
	}, nil
}

func (c *CoordinatorSVC) VerifySVC(p *cpb.Verify) (*cpb.VerifyResponce, error) {
	registerUser := fmt.Sprintf("register_user_%v", p.Email)
	redisVal := c.redis.Get(context.Background(), registerUser)

	if redisVal.Err() != nil {
		log.Printf("unable to get value from redis err: %v", redisVal.Err().Error())
		return nil, redisVal.Err()
	}

	var userData cDOM.User
	err := json.Unmarshal([]byte(redisVal.Val()), &userData)
	if err != nil {
		log.Println("unable to unmarshal json")
		return nil, err
	}

	code := fmt.Sprintf("%v", p.OTP)
	resp, err := c.twilio.VerifyTwilioOTP(string(userData.Phone), code)
	if err != nil {
		return nil, err
	} else {
		if resp.Status != nil {
			log.Println(*resp.Status)
		} else {
			log.Println(resp.Status)
		}
	}

	_, err = c.Repo.FindUserByEmail(userData.Email)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Existing email found  of a user %v", p.Email)
		return nil, errors.New("user already exists")
	}
	_, err = c.Repo.FindUserByPhone(userData.Phone)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Existing phone found  of a user %v", p.Email)
		return nil, errors.New("number already exists")
	}

	err = c.Repo.CreateUser(&userData)
	if err != nil {
		return nil, err
	}
	return &cpb.VerifyResponce{
		Status:  "Success",
		Message: "User creation done",
	}, nil

}

func NewCoordinatorSVC(repo inter.CoordinatorRepoInter) SVCinter.CoordinatorSVCInter {
	return &CoordinatorSVC{
		Repo: repo,
	}
}
