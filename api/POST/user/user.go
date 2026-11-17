package user

import (
	u "RESTapi/api/GET/user"
	"encoding/json"
	"net/http"
	
)

//post new user details
func User()http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handle post user logic here
		// persist database code here 
		// 

		response := &u.User{}
		err := json.NewEncoder(w).Encode(response)
		if err != nil{
			http.Error(w,"Internal Server Error",http.StatusInternalServerError)
			return
		}
	})
}