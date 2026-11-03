package payment

import (
	"database/sql"
	"net/http"
)

type Handlers struct{
	db *sql.DB
}

//retrive a payment made
func Payments()http.Handler{
	_ = &Handlers{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handle payments here
	})
}