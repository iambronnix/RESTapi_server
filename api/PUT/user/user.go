package user

import (
	"database/sql"
	"net/http"
)

type Handlers struct{
	db *sql.DB
}

//update user info here
func User()http.Handler{
	_ = &Handlers{}

  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      //handle all update user logic here
      w.WriteHeader(http.StatusFound)
		w.Write([]byte("put/user"))
  })
}