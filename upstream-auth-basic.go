package main

import (
	"encoding/base64"
	"fmt"

	"github.com/Kong/go-pdk"
)

type Config struct {
	Username string
	Password string
}

func New() interface{} {
	return &Config{}
}

func (conf Config) Access(kong *pdk.PDK) {
	authorizationValue := fmt.Sprintf("%s:%s", conf.Username, conf.Password)
	encodedAuthorizationValue := base64.StdEncoding.EncodeToString([]byte(authorizationValue))
	kong.ServiceRequest.SetHeader("Authorization",
		fmt.Sprintf("Basic %s", encodedAuthorizationValue))
}
