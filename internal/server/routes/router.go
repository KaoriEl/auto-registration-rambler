package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"main/internal/server/controllers"
	"main/internal/server/services"
	"main/internal/server/services/api"
	"main/internal/structures"
	"net/http"
)

func Router(router *mux.Router) {

	router.HandleFunc("/autoregistration", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		//files := Controllers.GetWidgets()

		//json.NewEncoder(w).Encode(files)
		v := services.AccInfo(r)
		fmt.Println(v)
		go controllers.Index(v)

		defer r.Body.Close()

	}).Methods("POST", "OPTIONS")

	router.HandleFunc("/verify_acc", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		c := make(chan string, 1)
		v := services.AccInfo(r)

		fmt.Println(v)
		controllers.Verify(v, c)

		json.NewEncoder(w).Encode(map[string]string{"VerifyLink": <-c})

		//json.NewEncoder(w).Encode(map[string]string{"VerifyLink": n})
		defer r.Body.Close()

	}).Methods("POST", "OPTIONS")

	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		i := structures.AccInfo{
			ClientId: 708715140,
		}
		api.CaptchaStatus("Тестовый роутер для капчи", i)
	}).Methods("GET", "OPTIONS")

}
