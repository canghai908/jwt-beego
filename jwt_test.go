package jwtbeego_test

import (
	"testing"
	"time"

	jwtbeego "github.com/canghai908/jwt-beego"
)

var (
	tokenStringGlobal string
)

// TestGetToken is the testing function por jwtbeego
func TestGetToken(t *testing.T) {
	et := jwtbeego.EasyToken{
		Username: "username",
		Expires:  time.Now().Unix() + 3600, //Segundos
	}
	tokenString, err := et.GetToken()
	tokenStringGlobal = tokenString
	if tokenString == "" {
		t.Errorf("Token String empty.")
	}
	if err != nil {
		t.Errorf("Error while verifying GetToken: %v", err)
	}
}

// TestValidateToken is the testing function por jwtbeego
func TestValidateToken(t *testing.T) {
	tokenString := tokenStringGlobal //c.Ctx.Input.Query("username")

	et := jwtbeego.EasyToken{}
	valido, err := et.ValidateToken(tokenString)

	if !valido {
		t.Errorf(err.Error())
	}

	return
}
