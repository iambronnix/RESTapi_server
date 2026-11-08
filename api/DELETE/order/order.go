package order

import (
	"database/sql"
	"net/http"
	"time"
)

type Handlers struct{
	db *sql.DB
}

//cancel an order
func Order()http.Handler{
	_ = &Handlers{}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handle delete order logic here
		time.Sleep(5*time.Second)
		w.WriteHeader(http.StatusFound)
		w.Write([]byte("delete/order"))
	})
}