package main

import "main/internal/server"

func main() {
	//initEnv()
	//registration.Registration()
	server.Serve()

	//reqBody, err := json.Marshal(map[string]string{
	//	"text":           "12313",
	//	"cid":            "213",
	//	"need_send_file": "false",
	//	"type_message":   "default_message",
	//})
	//if err != nil {
	//	print(err)
	//}
	//
	//resp, err := http.Post("http://investments-api-ms-nginx/api/v1/rambler/register/password",
	//	"application/x-www-form-urlencoded", bytes.NewBuffer(reqBody))
	//if err != nil {
	//	print(err)
	//}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	print(err)
	//}
	//fmt.Println(string(body))
}
