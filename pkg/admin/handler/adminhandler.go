package handler

import (
	"context"

	adminpb "github.com/Shakezidin/pkg/admin/pb"
	inter "github.com/Shakezidin/pkg/admin/service/interface"
)

type AdminHandler struct {
	svc inter.ServiceInterface
	adminpb.AdminServer
}

func NewAdminHandler(svc inter.ServiceInterface) *AdminHandler {
	return &AdminHandler{
		svc: svc,
	}
}

func (a *AdminHandler) AdminLoginRequest(ctx context.Context, p *adminpb.Login) (*adminpb.LoginResponce, error) {
	admin, err := a.svc.LoginService(p)
	if err != nil {
		return &adminpb.LoginResponce{
			Status: "False",
			Email:  p.Email,
			Token:  "",
		}, err
	}
	return admin, nil
}
