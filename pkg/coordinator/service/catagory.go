package service

import (
	"errors"
	"fmt"
	"log"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	dom "github.com/Shakezidin/pkg/entities/packages"
)

func (c *CoordinatorSVC) AddCatagorySVC(p *cpb.Category) (*cpb.Responce, error) {
	var catagory dom.Category
	_, err := c.Repo.FetchCatagory(p.CategoryName)
	if err == nil {
		log.Printf("Existing category found: %v", p.CategoryName)
		return nil, errors.New("category already exists")
	}

	catagory.Category = p.CategoryName
	err = c.Repo.CreateCatagory(catagory)
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
		Id:      int64(catagory.ID),
	}, nil
}

func (c *CoordinatorSVC) ViewCatagoriesSVC(p *cpb.View) (*cpb.Catagories, error) {
	offset := 10 * (p.Page - 1)
	limit := 10
	catagories, err := c.Repo.FetchCatagories(int(offset), limit)
	if err != nil {
		return nil, errors.New("error while fetching catagories")
	}

	var ctgries []*cpb.Category

	for _, cgry := range catagories {
		var ctgry cpb.Category
		ctgry.CatagoryId = int64(cgry.ID)
		ctgry.CategoryName = cgry.Category
		ctgries = append(ctgries, &ctgry)
	}

	return &cpb.Catagories{
		Catagories: ctgries,
	}, nil

}
