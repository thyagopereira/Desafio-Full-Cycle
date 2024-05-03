package web_handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thyagopereira/full-cycle/eda/usecases"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello from webServer.")
}

func GetAccountBalance(uc usecases.BalanceUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		accountId, ok := vars["account_id"]
		if !ok {
			w.Write([]byte(accountId))
		}
		dto, err := uc.GetBalanceFromAccoundId(accountId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(dto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}

}
