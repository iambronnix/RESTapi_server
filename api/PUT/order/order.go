package order

import (
	o "RESTapi/api/GET/order"
	"encoding/json"
	"fmt"
	"net/http"
)

//update order details
func Order()http.Handler{


	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handler all update order logic here
		//persist the db here
		param := r.URL.RawQuery
		fmt.Println(param)

		response := &o.Order{}

		err := json.NewEncoder(w).Encode(response)
		if err != nil{
			http.Error(w,"internal server err",http.StatusInternalServerError)
			return 
		}
	})
}