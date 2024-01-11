package interfaces

import (
	adminpb "github.com/Shakezidin/pkg/admin/pb"
)

type ServiceInterface interface {
	LoginService(p *adminpb.Login) (*adminpb.LoginResponce, error)
}
