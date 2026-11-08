package payment

import (
	"database/sql"
	"net/http"
	"time"
)

type Handlers struct{
	db *sql.DB
}

//update payments here
func Payment()http.Handler{
	_ = &Handlers{}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handle all update logic here
		time.Sleep(5*time.Second)
		w.WriteHeader(http.StatusFound)
		w.Write([]byte("put/payment"))
	})
}