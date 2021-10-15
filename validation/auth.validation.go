package validation

import (
	"regexp"
	"savi8sant8s/gotasks/entity"
	"savi8sant8s/gotasks/utils"
)

func ValidUser(user entity.User) (bool, string) {
	matchUserName, _ := regexp.MatchString(REGEX_NAME, user.Name)
	if !matchUserName {
		return false, utils.ERROR_NAME
	}
	matchUserEmail, _ := regexp.MatchString(REGEX_EMAIL, user.Email)
	if !matchUserEmail {
		return false, utils.ERROR_EMAIL
	}
	matchUserPassword, _ := regexp.MatchString(REGEX_PASSWORD, user.Password)
	if !matchUserPassword {
		return false, utils.ERROR_PASSWORD
	}
	return true, utils.BLANK
}

func ValidLogin(user entity.User) (bool, string) {
	matchEmail, _ := regexp.MatchString(REGEX_EMAIL, user.Email)
	if !matchEmail {
		return false, utils.ERROR_EMAIL
	}
	matchPassword, _ := regexp.MatchString(REGEX_PASSWORD, user.Password)
	if !matchPassword {
		return false, utils.ERROR_PASSWORD
	}
	return true, utils.BLANK
}
