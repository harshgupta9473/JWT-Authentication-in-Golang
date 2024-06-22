package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) error {
	userType := c.GetString("user-type")

	if userType != role {
		err := errors.New("unauthorized to access this resource")
		return err
	}
	return nil
}

func MatchUserTypeToUid(c *gin.Context, userID string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil
	if userType == "USER" && uid != userID {
		err = errors.New("unauthorized to access this resource")
		return err
	}
	return err
}
