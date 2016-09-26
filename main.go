package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"

	"github.com/30x/k8s-webhook/pkg/helper"
)

func main() {
	flag.Parse()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root"))
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Post("/authz", parseAuthzRequest)

	err := http.ListenAndServe(":8081", r)
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}

func parseAuthzRequest(w http.ResponseWriter, request *http.Request) {
	//Parse the json
	//Decode passed JSON body
	var requestJSON helper.AuthzRequestBody
	err := json.NewDecoder(request.Body).Decode(&requestJSON)
	if err != nil {
		fmt.Printf("Error decoding JSON Body: %s\n", err)
	}

	//Print out the user
	fmt.Printf("User: %v\n", requestJSON.Spec.User)

	//Do a check against permissions

	//Send the relevant response
	responseJSON := &helper.AuthzResponseBody{
		APIVersion: "authorization.k8s.io/v1beta1",
		Kind:       "SubjectAccessReview",
		Status: helper.Status{
			Allowed: true,
			Reason:  "testing",
		},
	}

	js, err := json.Marshal(responseJSON)
	if err != nil {
		fmt.Printf("Error Marshalling JSON response\n")
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(js)
}
