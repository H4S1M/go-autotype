package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/go-vgo/robotgo"
	"runtime"
)

type barcode struct {
	Result string `json:"Result"`
 }

type postHandler struct {
	ch *chan string
}


func main() {
	// TODO: what the hell is this? 
	runtime.GOMAXPROCS(2)

	buff := make(chan string, 30)
	mux := http.NewServeMux()

	// serve static file di folder app
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./app"))))

	PostHandler := postHandler{ch: &buff}
	mux.HandleFunc("/api", PostHandler.sendTobuff)

	// goroutine untuk autoType
	go func() {
		for {
			if len(buff) != 0 {
				robotgo.TypeStr(<-buff)
				robotgo.KeySleep = 25
				robotgo.KeyTap("enter")
			}
		}
	}()

	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *postHandler) sendTobuff(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	var Barcode barcode

	err := json.NewDecoder(r.Body).Decode(&Barcode)
	if err != nil {
		log.Print(err)
	}

	if method != "POST" {
		fmt.Fprintf(w, "!POST\n")
	}else {
		*h.ch <- Barcode.Result
	}
}
