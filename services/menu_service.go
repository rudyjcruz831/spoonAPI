package services

import (
	"context"
	"log"

	"github.com/rudyjcruz831/spoonAPI/model"
	"github.com/rudyjcruz831/spoonAPI/utils/errors"
)

// menuService acts as a struct for injecting an implementation of MenuService
// for use in service methods
type menuService struct {
	MenuAPI model.MenuAPI
}

// MenuConfig will hold repositories that will eventually be injected into this
// this service layer
type MSConfig struct {
	MenuAPI model.MenuAPI
}

// NewMenuService is a factory function for
// initializing a MenuService with its repository layer dependencies
func NewMenuService(c *MSConfig) model.MenuService {
	return &menuService{
		MenuAPI: c.MenuAPI,
	}
}

func (s *menuService) GetMenuItems(ctx context.Context, item string) (*errors.FoodError, *model.MenuItem) {
	// panic("GetMenuItems service")
	menuItems, foodErr := s.MenuAPI.GetMenuByName(item)
	if foodErr != nil {
		log.Println("Error running the GetMenuByName in API")
		return foodErr, nil
	}
	return nil, menuItems

}

func (s *menuService) GetMenuItemInfo(ctx context.Context, id int64) {
	panic("GetMenuItemInfo service")
}
