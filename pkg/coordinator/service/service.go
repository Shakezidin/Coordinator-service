package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
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
	cfg    *config.Config
}

func (c *CoordinatorSVC) SignupSVC(p *cpb.Signup) (*cpb.SignupResponce, error) {
	hashPassword, err := utils.HashPassword(p.Password)
	if err != nil {
		log.Printf("unable to hash password in CoordinatorSvc() - service, err: %v", err.Error())
		return nil, err
	}
	phone, _ := strconv.Atoi(p.Phone)
	user := &cDOM.User{
		Phone:    phone,
		Email:    p.Email,
		Password: string(hashPassword),
		Name:     p.Name,
	}
	// send otp to phone number

	resp, err := c.twilio.SentTwilioOTP(p.Phone)
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
	phone := strconv.Itoa(userData.Phone)
	fmt.Println(phone, "hhhhhhhhhhhhhhhhh")
	resp, err := c.twilio.VerifyTwilioOTP(phone, code)
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

func (c *CoordinatorSVC) UserLogin(p *cpb.CoorinatorLogin) (*cpb.CordinatorLoginResponce, error) {
	user, err := c.Repo.FindUserByEmail(p.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found og %v", p.Email)
			return nil, err
		} else {
			log.Printf("unable to login %v, err: %v", p.Email, err.Error())
			return nil, err
		}
	}

	check := utils.CheckPasswordMatch([]byte(user.Password), p.Password)
	if !check {
		log.Printf("password mismatch for user %v", p.Email)
		return nil, fmt.Errorf("password mismatch for user %v", p.Email)
	}

	token, err := utils.GenerateToken(p.Email, p.Role, config.LoadConfig().SECRETKEY)
	if err != nil {
		log.Printf("unable to generate token for user %v, err: %v", p.Email, err.Error())
		return nil, err
	}

	packages, _ := c.Repo.FindCoordinatorPackages(user.ID)

	var cdpackages []*cpb.Package
	for _, packagess := range *packages {
		// Check if packagess.Images is not nil before marshaling
		var imgs []byte
		if packagess.Images != nil {
			imgs, _ = json.Marshal(packagess.Images)
		}

		pkgs := &cpb.Package{
			DestinationCount: int32(packagess.NumOfDestination),
			Name:             packagess.Name,
			Destination:      packagess.Destination,
			Enddatetime:      packagess.EndDate,
			Endlocation:      packagess.EndLoaction,
			Image:            string(imgs),
			Price:            int32(packagess.Price),
			Startdatetime:    packagess.StartDate,
			Startlocation:    packagess.StartLocation,
		}

		cdpackages = append(cdpackages, pkgs)
	}

	return &cpb.CordinatorLoginResponce{
		Packages: cdpackages,
		Token:    token,
	}, nil
}

func NewCoordinatorSVC(repo inter.CoordinatorRepoInter, twilio *config.TwilioVerify, redis *redis.Client, cfg *config.Config) SVCinter.CoordinatorSVCInter {
	return &CoordinatorSVC{
		Repo:   repo,
		twilio: twilio,
		redis:  redis,
		cfg:    cfg,
	}
}
