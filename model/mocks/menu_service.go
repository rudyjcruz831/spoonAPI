package mocks

import (
	"context"

	"github.com/rudyjcruz831/spoonAPI/model"
	"github.com/rudyjcruz831/spoonAPI/utils/errors"
	"github.com/stretchr/testify/mock"
)

// MockMenuService is a mock type for model.MenuService
type MockMenuService struct {
	mock.Mock
}

// Get is mock of UserService Get
func (m *MockMenuService) GetMenuItems(ctx context.Context, item string) (*errors.FoodError, *model.MenuItem) {
	// args that will be passed to "Return" in the tests, when function
	// is called with a uid. Hence the name "ret"
	ret := m.Called(ctx, item)

	// first value passed to "Return"
	var r0 *model.MenuItem
	if ret.Get(0) != nil {
		// we can just return this if we know we won't be passing function to "Return"
		r0 = ret.Get(0).(*model.MenuItem)
	}

	var r1 *errors.FoodError
	if ret.Get(1) != nil {

		r1 = ret.Get(1).(*errors.FoodError)

	}

	return r1, r0
}

func (m *MockMenuService) GetMenuItemInfo(ctx context.Context, id int64) {

}
