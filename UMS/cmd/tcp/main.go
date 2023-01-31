package main

import (
	"UMS/api/rpc"
	"UMS/config"
	"UMS/internal/tcp/service"
	"log"
	"net/http"
)

func main() {
	go func() {
		err := http.ListenAndServe(config.PprofAddr, nil)
		if err != nil {
			return
		}
	}()

	var services service.UserServices
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	s := rpc.NewServer()
	s.Register(&services)
	err := s.ListenAndServe(config.TCPServerAddr)
	if err != nil {
		log.Fatalf("ListenAndServe failed! err: %v\n", err)
		return
	}
}
