package order

import (
	"fmt"
	"net/http"

)



//cancel an order
func Order()http.Handler{

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handle delete order logic here
		param := r.URL.RawQuery
		fmt.Println(param)
		
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("delete/order"))
	})
}