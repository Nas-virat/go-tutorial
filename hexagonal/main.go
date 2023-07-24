package main

import (
	"bank/handler"
	"bank/logs"
	"bank/repository"
	"bank/service"
	"fmt"
	_ "fmt"
	_ "log"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {

	initTimeZone()
	initConfig()
	db := initDatabase()

	// customer
	customerRepository := repository.NewCustomerRepositoryDB(db)
	customerRepositoryMock := repository.NewCustomerRepositoryMock()
	_ = customerRepository
	customerService := service.NewCustomerService(customerRepositoryMock)
	customerHandler := handler.NewCustomerHandler(customerService)

	//account 
	accountRepositoryDB := repository.NewAccountRepositoryDB(db)
	accountService := service.NewAccountService(accountRepositoryDB)
	accountHandler := handler.NewAccountHandler(accountService)


	_ = customerHandler

	router := mux.NewRouter()

	router.HandleFunc("/customers",customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	router.HandleFunc("/customers/{customerID:[0-9]+}/accounts",accountHandler.GetAccounts).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}/accounts",accountHandler.NewAccounts).Methods(http.MethodPost)

	//log.Printf("Banking service started at port %v",viper.GetInt("app.port"))
	logs.Info("Banking service started at port" + viper.GetString("app.port"))
	http.ListenAndServe(fmt.Sprintf(":%v",viper.GetInt("app.port")), router)
}


func initConfig(){
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".","_"))

	err := viper.ReadInConfig()
	if err != nil{
		panic(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func initDatabase() *sqlx.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)

	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}