package api

import (
	"github.com/gorilla/mux"
	"github.com/mohammad-hamzei/yourypto/api/responses"
	"net/http"
)
var dispatch dispatcher
func Start() error {

	r := mux.NewRouter()
	r.HandleFunc("/new", handleIncoming).Methods("POST")
	r.HandleFunc("/result", fetchResult).Methods("GET")
	if err := http.ListenAndServe(":8080", r); err != nil {
		return err
	}
	return nil
}

func handleIncoming(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	url := r.PostFormValue("url")

	if url == "" {
		w.WriteHeader(400)
		return
	}
	err := dispatch.HandleUrl(url)
	if err != nil {
		w.WriteHeader(500)
		return
	}
}

func fetchResult(w http.ResponseWriter, r *http.Request)  {
	url := r.URL.Query().Get("url")

	exists, err := dispatch.Exists(url)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	if !exists {
		w.WriteHeader(404)
		return
	}

	isPending, err := dispatch.IsPending(url)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	if isPending {
		w.WriteHeader(204)
		return
	}

	result, err := dispatch.GetResult(url)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	//w.WriteHeader(200)
	//w.Write([]byte(result))
	response := make(map[string]interface{})
	response["message"] = "success"
	response["status"] = http.StatusOK
	response["data"] = result
	responses.JSON(w, http.StatusOK, response)
}
