package handler

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rudyjcruz831/spoonAPI/utils/errors"
	// "github.com/rudyjcruz831/doorAPI/util/errors"
)

// used to help extract validation errors
type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

func bindData(c *gin.Context, req interface{}) bool {
	log.Println("binding data...")
	if c.ContentType() != "application/json" {
		log.Println("Failed content type is not application/json")
		msg := fmt.Sprintf("%s only accepts Content-Type application/json", c.FullPath())

		doorErr := errors.NewUnsupportedMediaType(msg)

		c.JSON(doorErr.Status, doorErr)
		return false
	}

	// Bind incoming json to struct and check for validation errors
	if err := c.ShouldBind(req); err != nil {
		log.Printf("Error binding data: %+v\n", err)

		// I need validator [github.com/go-playground/validator/v10] to be able to catch error
		if errs, ok := err.(validator.ValidationErrors); ok {
			log.Println("Failed to validator")
			// could probably extract this, it is also in middleware_auth_user
			var invalidArgs []invalidArgument

			for _, err := range errs {
				invalidArgs = append(invalidArgs, invalidArgument{
					err.Field(),
					err.Value().(string),
					err.Tag(),
					err.Param(),
				})
			}

			doorErr := errors.NewBadRequestError("Invalid request parameters. See invalidArgs")

			c.JSON(doorErr.Status, gin.H{
				"error":       doorErr,
				"invalidArgs": invalidArgs,
			})
			return false
		}
		// fmt.Println("req length: ", len(req.([]byte)))
		// later we'll add code for validating max body size here!

		// if we aren't able to properly extract validation errors,
		// we'll fallback and return an internal server error
		fallBack := errors.NewInternalServerError("did not properly extract validation errors")

		c.JSON(fallBack.Status, gin.H{"error": fallBack})
		return false
	}

	return true
}
