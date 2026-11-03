package order

import (
	"database/sql"
	"net/http"
)

type Handlers struct{
	db *sql.DB
}

//update order details
func Order()http.Handler{
	_ = &Handlers{}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handler all update order logic here
	})
}