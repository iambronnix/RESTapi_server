package order

import (
	"database/sql"
	"net/http"
)

type Handlers struct{
	db *sql.DB
}

//cancel an order
func Order()http.Handler{
	_ = &Handlers{}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handle delete order logic here
	})
}