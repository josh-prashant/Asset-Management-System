package main

import (
	"fmt"
	"net/http"

	"com/josh/asset/db"
	"com/josh/asset/handler"

	"github.com/gorilla/mux"
)

func main() {
	db.InitDatabase()
	StartHttp()
}

var (
	admin = "ADMIN"
	emp   = "EMP"
)

func StartHttp() {
	fmt.Println("Server started on 8081")
	// router := mux.NewRouter()
	router := mux.NewRouter().StrictSlash(true)
	// r.Use(CommonMiddleware)

	//shared
	router.HandleFunc("/login", handler.Login).Methods("POST")
	router.HandleFunc("/asset/list", handler.ReadAllAsset).Methods(http.MethodGet)
	router.HandleFunc("/employee/assets/{empId}", handler.EmployeeAssetAllocation).Methods(http.MethodGet)

	router.Handle("/employee/create", handler.JwtVerify(http.HandlerFunc(handler.Create), admin)).Methods(http.MethodPost)
	router.Handle("/employee/list", handler.JwtVerify(http.HandlerFunc(handler.ReadAll), admin)).Methods(http.MethodGet)
	router.Handle("/employee/delete/{empId}", handler.JwtVerify(http.HandlerFunc(handler.Delete), admin)).Methods(http.MethodDelete)
	router.Handle("/employee/update", handler.JwtVerify(http.HandlerFunc(handler.Update), emp)).Methods(http.MethodPost)
	// s := router.PathPrefix("/auth").Subrouter()
	router.Handle("/asset/create", handler.JwtVerify(http.HandlerFunc(handler.CreateAsset), admin)).Methods(http.MethodPost)
	router.Handle("/asset/delete/{assetId}", handler.JwtVerify(http.HandlerFunc(handler.DeleteAsset), admin)).Methods(http.MethodDelete)
	router.Handle("/asset/update", handler.JwtVerify(http.HandlerFunc(handler.EditAsset), admin)).Methods(http.MethodPost)
	router.Handle("/asset/assign", handler.JwtVerify(http.HandlerFunc(handler.AssignAsset), admin)).Methods(http.MethodPost)

	router.Handle("/request/create", handler.JwtVerify(http.HandlerFunc(handler.CreateRequest), emp)).Methods(http.MethodPost)
	router.Handle("/request/update", handler.JwtVerify(http.HandlerFunc(handler.UpdateRequestStatus), admin)).Methods(http.MethodPost)
	router.Handle("/employee/request/{empId}", handler.JwtVerify(http.HandlerFunc(handler.GetAllRequestsByEmpId), admin)).Methods(http.MethodGet)
	// router.Handle("/employee/request/{empId}", handler.JwtVerify(http.HandlerFunc(handler.FetchAllRequests), admin)).Methods(http.MethodGet)

	http.Handle("/", router)
	http.ListenAndServe(":8081", nil)
}

func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
