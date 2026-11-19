package order

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	schema "RESTapi/api/GET/order"
)

var buf bytes.Buffer

//create an order
func Order()http.Handler{

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handle post orders logic here
		json.NewDecoder(r.Body).Decode(&buf)
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
		param := r.URL.RawQuery
	    fmt.Println(param) 

		response := &schema.Order{}
		err := json.NewEncoder(w).Encode(response)
		if err != nil{
			http.Error(w,"Internal Server Error",http.StatusInternalServerError)
			return 
		}
		
	})
}

