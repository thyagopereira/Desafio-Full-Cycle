package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	web_handlers "github.com/thyagopereira/full-cycle/eda/webHandlers"
)

func main() {

	// Init database connection
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3308", "balance"))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer db.Close()

	// Init webserver
	r := mux.NewRouter()
	r.HandleFunc("/", web_handlers.HelloServer)
	r.HandleFunc("/balances/{account_id}", web_handlers.GetAccountBalance)
	http.ListenAndServe(":3003", r)

	// Init kafka connection
	// consumerConfigMap := ckafka.ConfigMap{
	// 	"bootstrap.servers": "kafka:29092",
	// 	"group.id":          "wallet",
	// }

	// topics := []string{"transactions", "balances"}
	// kafkaConsumer := kafka.NewConsumer(&consumerConfigMap, topics)
}
