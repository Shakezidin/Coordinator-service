package service

import (
	"errors"
	"strconv"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
)

func (c *CoordinatorSVC) ViewCoordinatorsSVC(p *cpb.View) (*cpb.Users, error) {
	offset := 10 * (p.Page - 1)
	limit := 10

	coordinators, err := c.Repo.FetchAllCoordinators(int(offset), limit)
	if err != nil {
		return &cpb.Users{
			Users: nil,
		}, errors.New("error while fetching all coordinators")
	}

	var coordinator []*cpb.User
	for _, cdntrs := range *coordinators {
		phone := strconv.Itoa(cdntrs.Phone)
		coordinator = append(coordinator, &cpb.User{
			Id:    int64(cdntrs.ID),
			Name:  cdntrs.Name,
			Email: cdntrs.Email,
			Phone: phone,
			Role:  cdntrs.Role,
		})
	}

	return &cpb.Users{
		Users: coordinator,
	}, nil
}
