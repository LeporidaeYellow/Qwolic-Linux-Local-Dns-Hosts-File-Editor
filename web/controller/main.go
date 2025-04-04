package controller

import (
	"hostsEditor/web"
	"log"
	"net/http"
)

func HostService() {
	http.HandleFunc("/", web.Home)
	http.HandleFunc("/adding", web.FileEditionByAdding)
	http.HandleFunc("/removing", web.FileEditionByRemoving)
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
