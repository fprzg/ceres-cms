package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Router() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", pageIndex)
	router.GET("/user/signup", pageSignup)
	router.GET("/user/login", pageLogin)
	router.GET("/user/logout", pageLogout)

	return router
}

func pageIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "index page")
}

func pageSignup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "signup page")
}

func pageLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "login page")
}

func pageLogout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "logout page")
}
