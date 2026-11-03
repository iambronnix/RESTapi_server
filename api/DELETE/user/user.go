package user

import (
	"database/sql"
	"net/http"
)

type Handlers struct{
	db *sql.DB
}

//delete user info 
func User()http.Handler{
	_ = &Handlers{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handle delete user logic here
	})
}