package service

import (
	"errors"
	"strconv"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
)

// ViewCoordinatorsSVC retrieves a list of coordinators.
func (c *CoordinatorSVC) ViewCoordinatorsSVC(p *cpb.View) (*cpb.Users, error) {
	// Define pagination parameters
	offset := 10 * (p.Page - 1)
	limit := 10

	// Fetch coordinators from the repository
	coordinators, err := c.Repo.FetchAllCoordinators(int(offset), limit)
	if err != nil {
		return nil, errors.New("error while fetching coordinators")
	}

	// Prepare response
	var pbCoordinators []*cpb.User
	for _, coordinator := range *coordinators {
		phone := strconv.Itoa(coordinator.Phone)
		pbCoordinator := &cpb.User{
			Id:    int64(coordinator.ID),
			Name:  coordinator.Name,
			Email: coordinator.Email,
			Phone: phone,
			Role:  coordinator.Role,
		}
		pbCoordinators = append(pbCoordinators, pbCoordinator)
	}

	return &cpb.Users{
		Users: pbCoordinators,
	}, nil
}
