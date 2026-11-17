package order

import (
	o "RESTapi/api/GET/order"
	"encoding/json"
	"net/http"
	
)

//update order details
func Order()http.Handler{


	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handler all update order logic here
		//persist the db here

		response := &o.Order{}

		err := json.NewEncoder(w).Encode(response)
		if err != nil{
			http.Error(w,"internal server err",http.StatusInternalServerError)
			return 
		}
	})
}