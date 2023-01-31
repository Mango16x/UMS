package main

import (
	"UMS/config"
	httpService "UMS/internal/http"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func init() {
	httpService.InitHttp()
}

func main() {

	// log and header
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// RPC
	var err error
	if err != nil {
		log.Fatal("Failed to connect with RPC service!")
		return
	}

	// static file service
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/images"))))

	http.HandleFunc("/index", httpService.GetIndex)
	http.HandleFunc("/profile", httpService.GetProfile)
	http.HandleFunc("/login", httpService.Login)
	http.HandleFunc("/signup", httpService.SignUp)
	http.HandleFunc("/nickname", httpService.UpdateNickName)
	http.HandleFunc("/avatar", httpService.UploadProfilePicture)

	if err := http.ListenAndServe(config.HTTPServerAddr, nil); err != nil {
		fmt.Println(err)
	}
}
