package server

import (
	"context"
	"flag"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"log"
	"main/internal/server/routes"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"
)

func Serve() {
	color.New(color.FgHiWhite).Add(color.Underline).Println("Server Tuning... ")
	var wait time.Duration
	var dir string
	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	router := mux.NewRouter()
	filePrefix, _ := filepath.Abs("/var/www/investments-auto-registration-rambler/captcha/")
	dir = filePrefix
	router.PathPrefix("/golang/captcha/").Handler(http.StripPrefix("/golang/captcha/", http.FileServer(http.Dir(dir))))

	routes.Router(router)

	color.New(color.FgHiWhite).Add(color.Underline).Println("Start server. Port:3002 ")

	srv := &http.Server{
		Handler: router,
		Addr:    ":3002",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)

	log.Println("RIP Server Shutdown")
	os.Exit(0)

}
