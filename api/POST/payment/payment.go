package payment

import (
	p "RESTapi/api/GET/payment"
	"encoding/json"
	"fmt"
	"net/http"
)

//create a Payment
func Payments()http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handle payment POST logic here
		// persist db operations here
		param := r.URL.RawQuery
		fmt.Println(param)

		response := &p.Order{}
		err := json.NewEncoder(w).Encode(response)
		if err != nil{
			http.Error(w,"Internal Server Error",http.StatusInternalServerError)
			return 
		}
	})
}