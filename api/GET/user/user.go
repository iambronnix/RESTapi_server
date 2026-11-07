package user

import (
	"database/sql"
	"net/http"
)

type Handlers struct{
	db *sql.DB
}
//retrieve user info
func Users()http.Handler{
	_ = &Handlers{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handle users code here
		w.WriteHeader(http.StatusFound)
		w.Write([]byte("This is user endpoint\n")) 
	})
}