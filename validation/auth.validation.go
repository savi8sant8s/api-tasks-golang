package validation

import (
	"regexp"
	"savi8sant8s/api/data"
	"savi8sant8s/api/entity"
	"savi8sant8s/api/utils"
)

type AuthValidation struct { }

func (validation *AuthValidation) ValidRegister(user entity.User) (bool, string) {
	matchName, _ := regexp.MatchString(REGEX_NAME, user.Name)
	if !matchName {
		return false, utils.MESSAGE_ERROR_NAME
	}
	matchEmail, _ := regexp.MatchString(REGEX_EMAIL, user.Email)
	if !matchEmail {
		return false, utils.MESSAGE_ERROR_EMAIL	
	}
	matchPassword, _ := regexp.MatchString(REGEX_PASSWORD, user.Password)
	if !matchPassword {
		return false, utils.MESSAGE_ERROR_PASSWORD
	}
	return true, utils.MESSAGE_BLANK
}

func (validation *AuthValidation) ValidLogin(user data.Login) (bool, string) {
	matchEmail, _ := regexp.MatchString(REGEX_EMAIL, user.Email)
	if !matchEmail {
		return false, utils.MESSAGE_ERROR_EMAIL
	}
	matchPassword, _ := regexp.MatchString(REGEX_PASSWORD, user.Password)
	if !matchPassword {
		return false, utils.MESSAGE_ERROR_PASSWORD	
	}
	return true, utils.MESSAGE_BLANK
}