package model

import (
	"context"

	"github.com/rudyjcruz831/spoonAPI/utils/errors"
	// "github.com/rudyjcruz831/leagueAPI/model"
)

type MenuService interface {
	GetMenuItems(ctx context.Context, item string) (*errors.FoodError, *MenuItem)
	GetMenuItemInfo(ctx context.Context, id int64)
}

type MenuAPI interface {
	GetMenuItemInfo(id int64) (*MenuItemInfo, *errors.FoodError)
	GetMenuByName(itemName string) (*MenuItem, *errors.FoodError)
}
