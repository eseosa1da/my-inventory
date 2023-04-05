package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (app *App) Initialize(DbUser string, DbPassword string, DBHost string, DbName string) error {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v)/%v", DbUser, DbPassword, DBHost, DBName)

	var err error
	app.DB, err = sql.Open("mysql", connectionString)

	if err != nil {
		return err
	}

	app.Router = mux.NewRouter().StrictSlash(true)
	app.handleRoutes()
	return nil
}

func (app *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, app.Router))
}

func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func sendError(w http.ResponseWriter, statusCode int, err string) {
	error_message := map[string]string{"error": err}
	sendResponse(w, statusCode, error_message)
}

func (app *App) getProducts(w http.ResponseWriter, r *http.Request) {

	products, err := getProducts(app.DB)

	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendResponse(w, http.StatusOK, products)

}

func (app *App) getProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])

	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid product ID")
		return
	}

	p := product{ID: key}

	err = p.getProduct(app.DB)

	if err != nil {

		switch err {
		case sql.ErrNoRows:
			sendError(w, http.StatusNotFound, "Product not found")
		default:
			sendError(w, http.StatusInternalServerError, err.Error())

		}

		return

	}

	sendResponse(w, http.StatusOK, p)

}

func (app *App) createProduct(w http.ResponseWriter, r *http.Request) {

	var p product

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		sendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := p.createProduct(app.DB); err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendResponse(w, http.StatusCreated, p)
}

func (app *App) updateProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])

	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid product ID")
		return
	}

	var p product

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		sendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	p.ID = key

	if err := p.updateProduct(app.DB); err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendResponse(w, http.StatusOK, p)
}

func (app *App) deleteProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])

	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid product ID")
		return
	}

	p := product{ID: key}

	msg, err := p.deleteProduct(app.DB)

	if err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	sendResponse(w, http.StatusOK, msg)
}

func (app *App) handleRoutes() {

	app.Router.HandleFunc("/products", app.getProducts).Methods("GET")
	app.Router.HandleFunc("/product/{id}", app.getProduct).Methods("GET")
	app.Router.HandleFunc("/product", app.createProduct).Methods("POST")
	app.Router.HandleFunc("/product/{id}", app.updateProduct).Methods("PUT")
	app.Router.HandleFunc("/product/{id}", app.deleteProduct).Methods("DELETE")

}
