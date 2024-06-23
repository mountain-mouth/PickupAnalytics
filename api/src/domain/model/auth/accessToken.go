package domain/model/auth

import (
	"time"
	"github.com/golang-jwt/jwt"
)

//constant message
var (
	ErrInvalidFormat = errors.New(Invalid id format)
)

//property
type AccessToken struct {
	value string
}

//factory method
func NewLoginId(value LoginId)(*AccessToken, error){
	if !validate(value){
		return nil, ErrInvalidFormat
	}
	return &AccessToken{value: value},nil

}

//get method
func (id *AccessToken) getValue() string {
	return id.value
}

//validation
func validate(value string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9]{3,20}$`)
	return regex.MatchString(value)
}

//generate jwt
func generateJwt(loginId LoginId, Password password) (AccessToken,error){
	claims := jwt.
}