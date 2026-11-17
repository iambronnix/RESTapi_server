package payment

import (
	p "RESTapi/api/GET/payment"
	"database/sql"
	"encoding/json"
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
		// persist db operations here

		response := &p.Order{}
		err := json.NewEncoder(w).Encode(response)
		if err != nil{
			http.Error(w,"Internal Server Error",http.StatusInternalServerError)
			return 
		}
	})
}