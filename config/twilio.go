package config

import (
	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

// TwilioVerify encapsulates Twilio client and configuration.
type TwilioVerify struct {
	Client *twilio.RestClient
	Cfg    *Config
}

// SetupTwilio initializes and returns a new TwilioVerify instance.
func SetupTwilio(cfg *Config) *TwilioVerify {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: cfg.SID,
		Password: cfg.TOKEN,
	})
	return &TwilioVerify{
		Client: client,
		Cfg:    cfg,
	}
}

// SendTwilioOTP sends an OTP to the provided phone number via Twilio.
func (t *TwilioVerify) SendTwilioOTP(phone string) (*verify.VerifyV2Verification, error) {
	params := verify.CreateVerificationParams{}
	params.SetTo("+91" + phone)
	params.SetChannel("sms")

	resp, err := t.Client.VerifyV2.CreateVerification(t.Cfg.SERVICETOKEN, &params)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// VerifyTwilioOTP verifies the provided OTP code for the given phone number via Twilio.
func (t *TwilioVerify) VerifyTwilioOTP(phone, code string) (*verify.VerifyV2VerificationCheck, error) {
	params := verify.CreateVerificationCheckParams{}
	params.SetTo("+91" + phone)
	params.SetCode(code)

	resp, err := t.Client.VerifyV2.CreateVerificationCheck(t.Cfg.SERVICETOKEN, &params)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
