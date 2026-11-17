package order

import (
	
	"encoding/json"
	"net/http"

 schema "RESTapi/api/GET/order"
)


//create an order
func Order()http.Handler{

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handle post orders logic here
		order := &schema.Order{}
		json.NewDecoder(r.Body).Decode(order)
		//persist database here

		
	   response := &schema.Order{}		
	    err :=	json.NewEncoder(w).Encode(response)
		if err != nil{
			http.Error(w,"Internal Server Error",http.StatusInternalServerError)
			return 
		}
	})
}
//cancel an order 
func CancelOrder()http.Handler{

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handle db operations here 
		// 

		response := &schema.Order{}
		err := json.NewEncoder(w).Encode(response)
		if err != nil{
			http.Error(w,"Internal Server Error",http.StatusInternalServerError)
			return 
		}
		
	})
}

