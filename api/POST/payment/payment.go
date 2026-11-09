package payment

import (
	"database/sql"
	"net/http"
	"time"
)

type Handlers struct{
	db *sql.DB
}

//create a Payment
func Payments()http.Handler{
	_ = &Handlers{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handle payment POST logic here
		time.Sleep(5*time.Second) 
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("post/payment"))
	})
}