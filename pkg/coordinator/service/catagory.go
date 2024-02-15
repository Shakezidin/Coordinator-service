package service

import (
	"errors"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	dom "github.com/Shakezidin/pkg/entities/packages"
)

// AddCategorySVC handles the addition of a new category.
func (c *CoordinatorSVC) AddCategorySVC(p *cpb.Category) (*cpb.Response, error) {
	// Check if category already exists
	_, err := c.Repo.FetchCategory(p.CategoryName)
	if err == nil {
		return &cpb.Response{
			Status:  "fail",
			Message: "category already exists",
		}, errors.New("category already exists")
	}

	// Create new category
	category := dom.Category{Category: p.CategoryName}
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
		Id:      int64(category.ID),
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
			CategoryId:   int64(category.ID),
			CategoryName: category.Category,
		}
		pbCategories = append(pbCategories, pbCategory)
	}

	return &cpb.Categories{
		Categories: pbCategories,
	}, nil
}
