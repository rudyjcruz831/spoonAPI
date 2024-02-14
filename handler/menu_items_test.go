package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rudyjcruz831/spoonAPI/model/mocks"
	"github.com/rudyjcruz831/spoonAPI/utils/errors"

	"github.com/stretchr/testify/mock"
)

func TestDetails(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	// router.Use(func(c *gin.Context))

	// mockUserService := new(mocks.MockUserService)

	mockMenuService := new(mocks.MockMenuService)

	NewHandler(&Config{
		R:           router,
		MenuService: mockMenuService,
	})
	// TODO: I need to create better testing
	// This function just checks that the GetMenuItem service function gets called
	t.Run("Getting data from API", func(t *testing.T) {
		rr := httptest.NewRecorder()

		reqBody, _ := json.Marshal(gin.H{
			"item": "burger",
		})
		request, _ := http.NewRequest(http.MethodPost, "/menu/menuItem", bytes.NewBuffer(reqBody))
		request.Header.Set("Content-Type", "application/json")

		// // mockErr :=
		mockMenuService.On("GetMenuItem", mock.AnythingOfType("*context.emptyCtx"), "burger").Return(errors.NewBadRequestError("one"), nil)
		router.ServeHTTP(rr, request)
		// log.Println(rr)
		// assert.Equal(t, http.StatusOK, rr.Code)
		mockMenuService.AssertNotCalled(t, "GetMenuItem")
	})

	// t.Run("Invalid request with missing item", func(t *testing.T) {
	// 	// router := gin.Default()
	// 	mockMenuService := new(mocks.MockMenuService)

	// 	NewHandler(&Config{
	// 		R:           router,
	// 		MenuService: mockMenuService,
	// 	})

	// 	reqBody, _ := json.Marshal(gin.H{})
	// 	request, _ := http.NewRequest(http.MethodPost, "/menu/menuItem", bytes.NewBuffer(reqBody))
	// 	request.Header.Set("Content-Type", "application/json")

	// 	rr := httptest.NewRecorder()
	// 	router.ServeHTTP(rr, request)

	// 	assert.Equal(t, http.StatusBadRequest, rr.Code)

	// 	// mockMenuService.AssertNotCalled(t, "GetMenuItems")
	// })

}
