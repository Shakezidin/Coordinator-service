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
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	cDOM "github.com/Shakezidin/pkg/entities/packages"

	"github.com/Shakezidin/utils"
	"gorm.io/gorm"
)

func (c *CoordinatorSVC) SignupSVC(p *cpb.Signup) (*cpb.Responce, error) {
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
		Role:     "coordinator",
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
	return &cpb.Responce{
		Status:  "success",
		Message: "user creation initiated, check message for OTP",
	}, nil
}

func (c *CoordinatorSVC) VerifySVC(p *cpb.Verify) (*cpb.Responce, error) {
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
	return &cpb.Responce{
		Status:  "Success",
		Message: "User creation done",
	}, nil

}

func (c *CoordinatorSVC) UserLogin(p *cpb.Login) (*cpb.LoginResponce, error) {
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

	userid := strconv.Itoa(int(user.ID))
	token, err := utils.GenerateToken(p.Email, p.Role, userid, config.LoadConfig().SECRETKEY)
	if err != nil {
		log.Printf("unable to generate token for user %v, err: %v", p.Email, err.Error())
		return nil, err
	}

	packages, _ := c.Repo.FindCoordinatorPackages(user.ID)

	var cdpackages []*cpb.Package
	for _, packagess := range *packages {

		pkgs := &cpb.Package{
			PackageId:        int64(packagess.ID),
			DestinationCount: int64(packagess.NumOfDestination),
			Packagename:      packagess.Name,
			Destination:      packagess.Destination,
			Enddate:          packagess.EndDate.Format("2006-01-02"),
			Image:            packagess.Images,
			Price:            int64(packagess.MinPrice),
			Startdate:        packagess.EndDate.Format("2006-01-02"),
			Startlocation:    packagess.StartLocation,
			Description:      packagess.Description,
			MaxCapacity:      int64(packagess.MaxCapacity),
		}

		cdpackages = append(cdpackages, pkgs)
	}

	return &cpb.LoginResponce{
		Packages: cdpackages,
		Token:    token,
	}, nil
}
