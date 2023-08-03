package iam

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

func Decode(client string, token string) (map[string]interface{}, error) {
	data := map[string]interface{} {
		"client": client,
		"token": token,
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
