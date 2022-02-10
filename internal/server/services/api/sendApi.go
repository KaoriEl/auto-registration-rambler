package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/internal/structures"
	"net/http"
	"strconv"
)

func SendCaptcha(text string, i structures.AccInfo, prefix string) {

	CaptchaPath := prefix + "-captcha.jpg"

	reqBody, err := json.Marshal(map[string]string{
		"text":           text,
		"need_send_file": "true",
		"path_file":      CaptchaPath,
		"cid":            strconv.FormatInt(i.ClientId, 10),
		"type_message":   "default_message",
	})
	if err != nil {
		print(err)
	}

	resp, err := http.Post("http://investments-go-bot:3000/api/v1/send-captcha",
		"application/x-www-form-urlencoded", bytes.NewBuffer(reqBody))
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Println(string(body))

}

func CaptchaStatus(text string, i structures.AccInfo) {
	reqBody, err := json.Marshal(map[string]string{
		"text":           text,
		"cid":            strconv.FormatInt(i.ClientId, 10),
		"need_send_file": "false",
		"type_message":   "default_message",
	})

	if err != nil {
		print(err)
	}

	resp, err := http.Post("http://investments-go-bot:3000/api/v1/status-captcha",
		"application/x-www-form-urlencoded", bytes.NewBuffer(reqBody))
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Println(string(body))

}

func ChangeStepUser(i structures.AccInfo) {
	reqBody, err := json.Marshal(map[string]string{
		"cid":          strconv.FormatInt(i.ClientId, 10),
		"current_step": "{\"step_name\":\"AUTOREG-RegCoinlist\",\"params\":[]}",
		"last_step":    "{\"step_name\":\"AUTOREG_wait_captcha\",\"params\":[]}",
	})

	if err != nil {
		print(err)
	}

	resp, err := http.Post("http://investments-api-ms-nginx/api/v1/rambler/change/step",
		"application/x-www-form-urlencoded", bytes.NewBuffer(reqBody))
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Println(string(body))

}

func RegisterPassword(i structures.AccInfo) {
	reqBody, err := json.Marshal(map[string]string{
		"id":       strconv.FormatInt(i.Id, 10),
		"status":   "true",
		"password": i.Password,
		"email":    i.Email,
	})
	if err != nil {
		print(err)
	}

	resp, err := http.Post("http://investments-api-ms-nginx/api/v1/rambler/register/password",
		"application/x-www-form-urlencoded", bytes.NewBuffer(reqBody))
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Println(string(body))

}
