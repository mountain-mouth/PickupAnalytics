package domain/model/auth

import (
	"errors"
	"regexp"
)

//constant message
var (
	ErrInvalidFormat = errors.New(Invalid id format)
)

//property
type LoginId struct {
	value string
}

//factory method
func NewLoginId(value string)(*LoginId, error){
	if !validate(value){
		return nil, ErrInvalidFormat
	}
	return &LoginId{value: value},nil

}

//get method
func (id *LoginId) getValue() string {
	return id.value
}

//validation
func validate(value string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9]{3,20}$`)
	return regex.MatchString(value)
}