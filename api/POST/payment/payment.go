package payment

import (
	"database/sql"
	"net/http"
)

type Handlers struct{
	db *sql.DB
}

//create a Payment
func Payments()http.Handler{
	_ = &Handlers{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handle payment POST logic here 
		w.WriteHeader(http.StatusFound)
		w.Write([]byte("post/payment"))
	})
}