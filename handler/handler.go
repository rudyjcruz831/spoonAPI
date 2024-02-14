package handler

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rudyjcruz831/spoonAPI/model"
)

type Handler struct {
	MaxBodyBytes int64
	MenuService  model.MenuService
}

type Config struct {
	R               *gin.Engine
	MenuService     model.MenuService
	BaseURL         string
	TimeoutDuration time.Duration
	MaxBodyBytes    int64
}

func NewHandler(c *Config) {
	// Create a handler (which will later have injected services)
	h := &Handler{
		MenuService:  c.MenuService,
		MaxBodyBytes: c.MaxBodyBytes,
	}

	// router for cors to be able to access from react
	c.R.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:80", "*"},
		AllowMethods: []string{"POST", "GET"},
		AllowHeaders: []string{"Content-Type", "Authorization", "No-Auth-Interceptor"},
	}))
	// Create an account group
	g := c.R.Group(c.BaseURL)

	// if gin.Mode() != gin.TestMode {

	// } else {

	// }

	// g.POST("/refresh", h.Token)
	g.GET("/", h.Home)
	g.POST("menu/menuItem", h.MenuItems)
	g.POST("menu/menuItemInfo")

}

// func (h *Handler) Login(c *gin.Context) {}

func (h *Handler) Home(c *gin.Context) {
	// time.Sleep(6 * time.Second)
	c.JSON(http.StatusOK, map[string]string{"Its working": "kind of"})
}
