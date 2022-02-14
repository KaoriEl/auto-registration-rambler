package services

import (
	"encoding/json"
	"io/ioutil"
	"main/internal/structures"
	"net/http"
)

func AccInfo(r *http.Request) []structures.AccInfo {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var t []structures.AccInfo
	err = json.Unmarshal(body, &t)
	if err != nil {
		return nil
	}
	return t
}
