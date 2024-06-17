package api

import (
	"musclemate/endpoints/account"
	"musclemate/endpoints/workouts"
	ReturnModule "musclemate/utils/return_module"
	"net/http"
)

func HandleAccounts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		account.GetRequest(w, r)
	} else {
		ReturnModule.MethodNotAllowed(w)
	}
}

func HandleWorkouts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		workouts.GetRequest(w, r)
	} else if r.Method == "POST" {
		workouts.PostRequest(w, r)
	} else if r.Method == "PUT" {
		workouts.PutRequest(w, r)
	} else if r.Method == "DELETE" {
		workouts.DeleteRequest(w, r)
	} else {
		ReturnModule.MethodNotAllowed(w)
	}
}

func HandleExercises(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		workouts.GetRequest(w, r)
	} else if r.Method == "POST" {
		workouts.PostRequest(w, r)
	} else if r.Method == "PUT" {
		workouts.PutRequest(w, r)
	} else if r.Method == "DELETE" {
		workouts.DeleteRequest(w, r)
	} else {
		ReturnModule.MethodNotAllowed(w)
	}
}
