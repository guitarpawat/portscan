// Package web serves the RESTful API service server for port scanner
package web

import (
	"github.com/gorilla/mux"
	"github.com/guitarpawat/portscan/api"
	"github.com/guitarpawat/portscan/api/model"
	"io/ioutil"
	"net/http"
)

// ListenAndServe serves HTTP server with specified host:port
func ListenAndServe(serve string) error {
	r := mux.NewRouter()
	r.HandleFunc("/api/scan", handlePostScanRequest).Methods("POST")
	r.HandleFunc("/api/token/{token}", handleGetTokenUpdate).Methods("GET")
	r.HandleFunc("/api/token/{token}", handleDeleteToken).Methods("DELETE")

	return http.ListenAndServe(serve, r)
}

func handlePostScanRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(w, err)
		return
	}

	result := api.PutNewScanRequest(body)
	returnResult(w, result)
	return
}

func handleGetTokenUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	in := model.GetInput{
		Token: vars["token"],
	}
	b, err := in.Marshal()
	if err != nil {
		handleError(w, err)
		return
	}
	result := api.GetUpdateScanResult(b)
	returnResult(w, result)
	return
}

func handleDeleteToken(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	in := model.GetInput{
		Token: vars["token"],
	}
	b, err := in.Marshal()
	if err != nil {
		handleError(w, err)
		return
	}
	result := api.KillScanRequest(b)
	if result == nil {
		w.WriteHeader(200)
		return
	}

	returnResult(w, result)
	return
}

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(400)
	jsonErr, _ := model.MakeError(err).Marshal()
	w.Write(jsonErr)
	return
}

func returnResult(w http.ResponseWriter, result model.Json) {
	jsonResult, err := result.Marshal()
	if err != nil {
		handleError(w, err)
		return
	}

	_, ok := result.(error)
	if ok {
		w.WriteHeader(400)
		b, _ := result.Marshal()
		w.Write(b)
		return
	}

	w.WriteHeader(200)
	w.Write(jsonResult)
}
