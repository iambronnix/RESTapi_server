package user

import (
	"database/sql"
	"net/http"
)

type Handlers struct{
	db *sql.DB
}

//post new user details 
func User()http.Handler{
	_ = &Handlers{}
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handle post user logic here
		w.WriteHeader(http.StatusFound)
		w.Write([]byte("post/user")) 
	})
}