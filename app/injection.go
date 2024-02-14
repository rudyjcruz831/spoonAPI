package app

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rudyjcruz831/spoonAPI/food"
	"github.com/rudyjcruz831/spoonAPI/handler"
	"github.com/rudyjcruz831/spoonAPI/services"
)

func inject() (*gin.Engine, error) {
	log.Println("Injecting data sources...")

	/*
	 * repository layer
	 */

	/*
	 * api layer
	 */

	// get API key from enviorment variable
	apiKey := os.Getenv("API_KEY")
	menuApi := food.NewMenuAPI(apiKey)

	/*
	 * service layer
	 */

	menuService := services.NewMenuService(&services.MSConfig{
		MenuAPI: menuApi,
	})

	// initialize gin.Engine
	router := gin.Default()

	// read in project baseURL from environment variable
	// TODO: add to env variable
	baseURL := os.Getenv("BASE_URL")
	fmt.Println("baseurl: ", baseURL)

	// for timing out the call if it takes to long
	// TODO: add to env variables
	handlerTimeout := os.Getenv("HANDLER_TIMEOUT")
	ht, err := strconv.ParseInt(handlerTimeout, 0, 64)
	// fmt.Println(ht)
	if err != nil {
		return nil, fmt.Errorf("could not parse HANDLER_TIMEOUT as int: %w", err)
	}

	handler.NewHandler(&handler.Config{
		R:               router,
		MenuService:     menuService,
		BaseURL:         baseURL,
		TimeoutDuration: time.Duration(time.Duration(ht) * time.Second),
		MaxBodyBytes:    1024 * 1024 * 1024,
	})

	return router, nil

}
