package main

import (
	"database/sql"
	"fmt"
	"net/http"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/thyagopereira/full-cycle/eda/internal/databases"
	handler "github.com/thyagopereira/full-cycle/eda/internal/events/handlers"
	"github.com/thyagopereira/full-cycle/eda/kafka"
	"github.com/thyagopereira/full-cycle/eda/pkg/events"
	"github.com/thyagopereira/full-cycle/eda/usecases"
	web_handlers "github.com/thyagopereira/full-cycle/eda/webHandlers"
)

func main() {
	// Init database connection
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql_balance", "3306", "balance"))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("NO DB connection")
		panic(err)
	}

	// Inits balanceDB and accountsDB
	balanceDB := databases.NewBalanceDB(db)
	accountsDB := databases.NewAccountsDB(db)

	// Init use cases
	balanceUseCases := usecases.NewBalanceUseCases(*accountsDB)

	// Init webserver
	r := mux.NewRouter()
	r.HandleFunc("/", web_handlers.HelloServer)
	r.HandleFunc("/balances/{account_id}", web_handlers.GetAccountBalance(balanceUseCases))
	go http.ListenAndServe(":3003", r)
	fmt.Println("Server is Up.")

	// Init kafka connection
	consumerConfigMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}
	topics := []string{"balances"}
	kafkaConsumer := kafka.NewConsumer(&consumerConfigMap, topics)

	// Inits event Dispatcher
	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("BalanceUpdated", handler.NewUpdateBalanceKafkaHandler(*balanceDB, *accountsDB))

	// Listen to events
	kafkaConsumer.Consume(*eventDispatcher)

}
