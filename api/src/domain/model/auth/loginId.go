package auth

import (
	"errors"
	"regexp"
)

type LoginId struct {
	value string
}

func NewLoginId(value string)(*LoginId, error){
	if !validate(value){
		return nil, errors.New("Invalid id format")
	}
	return &LoginId{value: value},nil

}

func validate(value string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9]{3,20}$`)
	return regex.MatchString(value)
}