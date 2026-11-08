package user

import (
	"database/sql"
	"net/http"
	"time"
)

type Handlers struct{
	db *sql.DB
}

//delete user info 
func User()http.Handler{
	_ = &Handlers{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handle delete user logic here
		time.Sleep(5*time.Second)
		w.WriteHeader(http.StatusFound)
		w.Write([]byte("delete/user"))
	})
}