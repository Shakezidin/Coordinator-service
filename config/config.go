package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host                string
	User                string
	Password            string
	Database            string
	Port                string
	Sslmode             string
	GRPCCOORDINATORPORT string
	SID                 string
	TOKEN               string
	SERVICETOKEN        string
	SECRETKEY           string
	REDISHOST           string
	RAZORPAYKEYID       string
	RAZORPAYSECRETKEY   string
	GRPCADMINPORT       string
	REDISPassword       string
}

func LoadConfig() *Config {
	godotenv.Load("../../.env")

	var config Config

	// Use os.Getenv to retrieve environment variables
	config.Host = os.Getenv("HOST")
	config.User = os.Getenv("USER")
	config.Password = os.Getenv("PASSWORD")
	config.Database = os.Getenv("DATABASE")
	config.Port = os.Getenv("PORT")
	config.Sslmode = os.Getenv("SSLMODE")
	config.GRPCCOORDINATORPORT = os.Getenv("GRPCCOORDINATORPORT")
	config.SID = os.Getenv("SID")
	config.TOKEN = os.Getenv("TOKEN")
	config.SERVICETOKEN = os.Getenv("SERVICETOKEN")
	config.SECRETKEY = os.Getenv("SECRETKEY")
	config.REDISHOST = os.Getenv("REDISHOST")
	config.RAZORPAYKEYID = os.Getenv("RAZORPAYKEYID")
	config.RAZORPAYSECRETKEY = os.Getenv("RAZORPAYSECRETKEY")
	config.GRPCADMINPORT = os.Getenv("GRPCADMINPORT")
	config.REDISPassword = os.Getenv("REDISPASSWORD")

	return &config
}
