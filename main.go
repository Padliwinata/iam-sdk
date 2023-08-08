package iam

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func Decode(client string, token string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"client": client,
		"token":  token,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("https://fastapiiam-1-i2172913.deta.app/decode", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var responseMap map[string]interface{}
	err = json.Unmarshal(responseData, &responseMap)
	if err != nil {
		return nil, err
	}

	return responseMap, nil
}

func DecodeWithSecret(accessToken string, secretKey string) (map[string]interface{}, error) {
	secret := []byte(secretKey)

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, errors.New("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Failed to parse claims")
	} else {
		return claims, nil
	}
}

func CheckAuth(accessToken string, secretKey string) bool {
	_, err := DecodeWithSecret(accessToken, secretKey)
	return err != nil
}

func CheckPermission(accessToken string, secretKey string, scope string) bool {
	claims, err := DecodeWithSecret(accessToken, secretKey)
	if err != nil {
		return false
	}

	scopes, exists := claims["role"]

	if !exists {
		return false
	}

	scopesString, ok := scopes.(string)
	if !ok {
		return false
	}
	scopesArray := strings.Split(scopesString, " ")

	for _, str := range scopesArray {
		if str == scope {
			return true
		}
	}

	return false

}
