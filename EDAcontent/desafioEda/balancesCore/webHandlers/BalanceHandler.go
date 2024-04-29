package web_handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {

	fmt.Println(w, "Hello from webServer.")
}

func GetAccountBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId, ok := vars["account_id"]
	if !ok {
		fmt.Println(w, "Missing account id")
	}
	fmt.Println(w, accountId)
}
