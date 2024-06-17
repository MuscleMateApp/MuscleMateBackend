package api

import (
	"musclemate/endpoints/account"
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