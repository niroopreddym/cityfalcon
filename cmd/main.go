package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/niroopreddym/cityfalcon/pkg/handlers"
)

func main() {
	// // Capture Ctrl-C
	// ctx := context.Background()
	// // trap Ctrl+C and call cancel on the context
	// ctx, cancel := context.WithCancel(ctx)
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)
	// defer func() {
	// 	signal.Stop(c)
	// 	cancel()
	// }()

	// //--------------logging middleware---------------------
	// logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	// stdlog.SetOutput(log.NewStdlibAdapter(logger))
	// logger = log.With(logger, "ts", log.DefaultTimestampUTC, "loc", log.DefaultCaller)
	// loggingMiddleware := middleware.LoggingMiddleware(logger)
	//---------------------------------------------------
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

	// //-------------------------prometheous endpoints--------------------
	// promMiddleware := middleware.PrometheusMiddleware
	// router.Use(promMiddleware)

	// router.PathPrefix("/").Handler(http.FileServer(http.Dir("../static/")))
	// router.Handle("/metrics", promhttp.Handler()).Methods("GET")
	// //--------------------------------------------------------------------

	// loggedRouter := loggingMiddleware(router)
	// if err := http.ListenAndServe(":9295", loggedRouter); err != nil {
	// 	logger.Log("status", "fatal", "err", err)
	// }

	// //-------------capturing the ctrl + c event----------------------
	// select {
	// case <-c:
	// 	stdlog.Println("cancel operation")
	// 	cancel()

	// case <-ctx.Done():
	// 	time.Sleep(600 * time.Millisecond)
	// }

	// stdlog.Println("done")

	// defer func() {
	// 	fmt.Println("Hello")
	// }()

	http.ListenAndServe(":9295", router)
}
