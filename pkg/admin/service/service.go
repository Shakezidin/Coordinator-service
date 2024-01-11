package service

import (
	"errors"
	"log"

	jwt "github.com/Shakezidin/pkg/JWT"
	adminpb "github.com/Shakezidin/pkg/admin/pb"
	inter "github.com/Shakezidin/pkg/admin/repository/interface"
	interr "github.com/Shakezidin/pkg/admin/service/interface"
)

type AdminService struct {
	Repo inter.RepoInterface
}

func (a *AdminService) LoginService(p *adminpb.Login) (*adminpb.LoginResponce, error) {
	admin, err := a.Repo.FetchAdmin(p.Email)
	if err != nil {
		return nil, err
	}
	if admin.Password != p.Password {
		return nil, errors.New("password incorrect")
	}

	token, err := jwt.GenerateJWT(admin.Email, p.Role)
	if err != nil {
		log.Print("Generate jwt error")
		return nil, err
	}
	adminn := &adminpb.LoginResponce{
		Status: "Success",
		Email:  admin.Email,
		Token:  token,
	}

	return adminn, nil
}

func NewAdminService(repos inter.RepoInterface) interr.ServiceInterface {
	return &AdminService{
		Repo: repos,
	}
}
