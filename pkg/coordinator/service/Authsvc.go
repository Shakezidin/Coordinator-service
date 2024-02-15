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

// SignupSVC handles the signup process.
func (c *CoordinatorSVC) SignupSVC(p *cpb.Signup) (*cpb.Response, error) {
	hashPassword, err := utils.HashPassword(p.Password)
	if err != nil {
		return nil, errors.New("unable to hash password")
	}

	phone, _ := strconv.Atoi(p.Phone)
	user := &cDOM.User{
		Phone:    phone,
		Email:    p.Email,
		Password: string(hashPassword),
		Name:     p.Name,
		Role:     "coordinator",
	}

	// Send OTP
	resp, err := c.twilio.SendTwilioOTP(p.Phone)
	if err != nil {
		return nil, errors.New("error while sending OTP")
	} else {
		if resp.Status != nil {
			log.Println(*resp.Status)
		} else {
			log.Println(resp.Status)
		}
	}

	// Store user data in Redis for OTP verification
	userData, err := json.Marshal(&user)
	if err != nil {
		return nil, errors.New("error while marshalling user data")
	}

	registerUser := fmt.Sprintf("register_user_%v", p.Email)
	c.redis.Set(context.Background(), registerUser, userData, time.Minute*2)

	return &cpb.Response{
		Status:  "success",
		Message: "user creation initiated, check message for OTP",
	}, nil
}

// VerifySVC verifies the OTP and creates the user.
func (c *CoordinatorSVC) VerifySVC(p *cpb.Verify) (*cpb.Response, error) {
	registerUser := fmt.Sprintf("register_user_%v", p.Email)
	redisVal := c.redis.Get(context.Background(), registerUser)
	if redisVal.Err() != nil {
		return &cpb.Response{Status: "failed"}, errors.New("unable to get value from redis")
	}

	// Unmarshal user data from Redis
	var userData cDOM.User
	err := json.Unmarshal([]byte(redisVal.Val()), &userData)
	if err != nil {
		return &cpb.Response{Status: "failed"}, errors.New("error while unmarshalling data")
	}

	code := fmt.Sprintf("%v", p.OTP)
	phone := strconv.Itoa(userData.Phone)
	resp, err := c.twilio.VerifyTwilioOTP(phone, code)
	if err != nil {
		return &cpb.Response{Status: "failed"}, errors.New("OTP verification failed")
	} else {
		if resp.Status != nil {
			log.Println(*resp.Status)
		} else {
			log.Println(resp.Status)
		}
	}

	// Check if email or phone already exists
	_, err = c.Repo.FindUserByEmail(userData.Email)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return &cpb.Response{Status: "failed"}, errors.New("email already exists")
	}

	_, err = c.Repo.FindUserByPhone(userData.Phone)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return &cpb.Response{Status: "failed"}, errors.New("phone number already exists")
	}

	// Create the user
	err = c.Repo.CreateUser(&userData)
	if err != nil {
		return &cpb.Response{Status: "failed"}, errors.New("error while creating user")
	}

	return &cpb.Response{
		Status:  "success",
		Message: "user creation done",
	}, nil
}

// UserLogin handles user login.
func (c *CoordinatorSVC) UserLogin(p *cpb.Login) (*cpb.LoginResponse, error) {
	// Find user by email
	user, err := c.Repo.FindUserByEmail(p.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("error while logging in")
	}

	// Check password
	check := utils.CheckPasswordMatch([]byte(user.Password), p.Password)
	if !check {
		return nil, fmt.Errorf("password mismatch for user %v", p.Email)
	}

	// Generate JWT token
	userID := strconv.Itoa(int(user.ID))
	token, err := utils.GenerateToken(p.Email, p.Role, userID, config.LoadConfig().SECRETKEY)
	if err != nil {
		return nil, errors.New("error while generating JWT")
	}

	return &cpb.LoginResponse{
		Token: token,
	}, nil
}
