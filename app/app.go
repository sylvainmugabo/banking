package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sylvainmugabo/microservices-lib/logger"
	"github.com/sylvainmugabo/microservices/banking/config"
	"github.com/sylvainmugabo/microservices/banking/domain"
	"github.com/sylvainmugabo/microservices/banking/services"
)

func Start(conf *config.EnvConfigs) {
	validationCheck(conf)

	dbclient := getDbClient(conf)
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbclient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbclient)

	ch := CustomerHandlers{service: services.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service: services.NewAccountService(accountRepositoryDb)}

	router := mux.NewRouter()
	router.
		HandleFunc("/customers", ch.getAllCustomers).
		Methods(http.MethodGet).
		Name("GetAllCustomers")
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).
		Methods(http.MethodGet).
		Name("GetCustomer")
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).
		Methods(http.MethodPost).
		Name("NewAccount")
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).
		Methods(http.MethodPost).
		Name("NewTransaction")

	am := AuthMiddleware{domain.NewAuthRepository()}
	router.Use(am.authorizationHandler())
	log.Fatal(http.ListenAndServe(conf.ServerPort, router))
}

func validationCheck(conf *config.EnvConfigs) {
	if conf.DatabaseAddress == "" || conf.DatabaseName == "" || conf.DatabasePort == "" || conf.DatabasePwd == "" || conf.DatabaseUser == "" {
		logger.Fatal("Environment variable(s) not defined ...")
	}
}

func getDbClient(conf *config.EnvConfigs) *sqlx.DB {

	dbUser := conf.DatabaseUser
	dbPasswd := conf.DatabasePwd
	dbAddr := conf.DatabaseAddress
	dbPort := conf.DatabasePort
	dbName := conf.DatabaseName
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	fmt.Println(dataSource)
	client, err := sqlx.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}
	client.SetConnMaxIdleTime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
