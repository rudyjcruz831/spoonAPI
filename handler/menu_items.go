package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type menuItemReq struct {
	Item string `json:"item" binding:"required"`
}

func (h *Handler) MenuItems(c *gin.Context) {
	var req menuItemReq

	if ok := bindData(c, &req); !ok {
		fmt.Println("binding data unsuccessful")
		return
	}

	ctx := c.Request.Context()
	log.Println("Item: ", req.Item)
	// get menu items return by name
	foodErr, mItem := h.MenuService.GetMenuItems(ctx, req.Item)
	if foodErr != nil {
		log.Println("Error in MenuService")
		c.JSON(foodErr.Status, foodErr)
	}

	c.JSON(http.StatusOK, mItem)

}
