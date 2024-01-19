package service

import (
	"errors"
	"fmt"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	dom "github.com/Shakezidin/pkg/entities/packages"
)

func (c *CoordinatorSVC) AddCatagorySVC(p *cpb.Category) (*cpb.Responce, error) {
	var catagory dom.Category
	catagory.Category = p.CategoryName
	err := c.Repo.CreateCatagory(catagory)
	if err != nil {
		fmt.Println("error while creating category")
		return &cpb.Responce{
			Status:  "fail",
			Message: "error while creating category",
		}, err
	}
	return &cpb.Responce{
		Status:  "success",
		Message: "catagory created successsfully",
	}, nil
}

func (c *CoordinatorSVC) ViewCatagoriesSVC(p *cpb.View) (*cpb.Catagories, error) {
	catagories, err := c.Repo.FetchCatagories()
	if err != nil {
		return nil, errors.New("error while fetching catagories")
	}

	var ctgry cpb.Category
	var ctgries []*cpb.Category

	for _, cgry := range catagories {
		ctgry.CatagoryId = int64(cgry.ID)
		ctgry.CategoryName = cgry.Category
		ctgries = append(ctgries, &ctgry)
	}

	return &cpb.Catagories{
		Catagories: ctgries,
	}, nil

}
