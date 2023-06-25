package main

import (
	"fmt"
	stdlog "log"
	"net/http"
	"os"

	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/niroopreddym/cityfalcon/middleware"
	"github.com/niroopreddym/cityfalcon/pkg/handlers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	//--------------logging middleware---------------------
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	stdlog.SetOutput(log.NewStdlibAdapter(logger))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "loc", log.DefaultCaller)
	loggingMiddleware := middleware.LoggingMiddleware(logger)
	// ---------------------------------------------------

	router := mux.NewRouter()
	handler := handlers.NewBankAndAccountsHandlerInstance()
	fmt.Println("started listening on port : ", 9295)
	router.Handle("/bank", http.HandlerFunc(handler.CreateBank)).Methods("POST")
	router.Handle("/bank", http.HandlerFunc(handler.GetAllBanks)).Methods("GET")
	router.Handle("/bank/{uuid}", http.HandlerFunc(handler.GetBankDetails)).Methods("GET")
	router.Handle("/bank/{uuid}", http.HandlerFunc(handler.UpdateBankDetails)).Methods("PUT")
	router.Handle("/bank/{uuid}", http.HandlerFunc(handler.RemoveBank)).Methods("DELETE")

	router.Handle("/account", http.HandlerFunc(handler.CreateAccount)).Methods("POST")
	//the below endpoint is long running process
	router.Handle("/account/{uuid}", http.HandlerFunc(handler.GetAccountDetails)).Methods("GET")
	router.Handle("/account/getaccountdetails/asyncresponse/{uuid}", http.HandlerFunc(handler.GetAccountDetailsResponse)).Methods("GET")
	router.Handle("/account/{uuid}", http.HandlerFunc(handler.UpdateAccountDetails)).Methods("PUT")

	//-------------------------prometheous endpoints--------------------
	promMiddleware := middleware.PrometheusMiddleware
	router.Use(promMiddleware)
	router.Handle("/metrics", promhttp.Handler()).Methods("GET")
	//--------------------------------------------------------------------

	loggedRouter := loggingMiddleware(router)
	if err := http.ListenAndServe(":9295", loggedRouter); err != nil {
		logger.Log("status", "fatal", "err", err)
	}
}
