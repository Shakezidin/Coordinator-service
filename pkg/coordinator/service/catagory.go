package service

import (
	"errors"
	"fmt"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	dom "github.com/Shakezidin/pkg/entities/packages"
	"gorm.io/gorm"
)

// AddCategorySVC handles the addition of a new category.
func (c *CoordinatorSVC) AddCategorySVC(p *cpb.Category) (*cpb.Response, error) {
	// Check if category already exists
	fmt.Println(p.Category_Name)
	_, err := c.Repo.FetchCategory(p.Category_Name)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return &cpb.Response{Status: "failed"}, errors.New("category already exists")
	}

	// Create new category
	category := dom.Category{Category: p.Category_Name}
	err = c.Repo.CreateCategory(category)
	if err != nil {
		return &cpb.Response{
			Status:  "fail",
			Message: "error while creating category",
		}, errors.New("error while creating category")
	}

	return &cpb.Response{
		Status:  "success",
		Message: "category created successfully",
		ID:      int64(category.ID),
	}, nil
}

// ViewCategoriesSVC retrieves a list of categories.
func (c *CoordinatorSVC) ViewCategoriesSVC(p *cpb.View) (*cpb.Categories, error) {
	// Define pagination parameters
	offset := 10 * (p.Page - 1)
	limit := 10

	// Fetch categories from the repository
	categories, err := c.Repo.FetchCategories(int(offset), limit)
	if err != nil {
		return nil, errors.New("error while fetching categories")
	}

	// Prepare response
	var pbCategories []*cpb.Category
	for _, category := range categories {
		pbCategory := &cpb.Category{
			Category_ID:   int64(category.ID),
			Category_Name: category.Category,
		}
		pbCategories = append(pbCategories, pbCategory)
	}

	return &cpb.Categories{
		Categories: pbCategories,
	}, nil
}
