package user

import (
	"database/sql"
	"encoding/json"
	"net/http"
)
type User struct {
	Id int
	UserName ,
	Created ,
	Password string
}

type Handlers struct{
	 *sql.DB
}
//retrieve user info
func Users()http.Handler{
	_ = &Handlers{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handle users code here
		// persist db here

		
		response := &User{}
		err := json.NewEncoder(w).Encode(response)
		if err != nil{
			http.Error(w,"Internal Server Error",http.StatusInternalServerError)
			return 
		}
	
	})
}