package order

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "github.com/google/uuid"
)





//retrive general lists of orders
func ListOrders()http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handle code here
		// persist the database here
		param := r.URL.RawQuery
	    fmt.Println(param) 

				
		response := &Orders{}		
		err := json.NewEncoder(w).Encode(response)
		if err != nil{
			http.Error(w, "internal server err",http.StatusInternalServerError)
			return 
		}
		
	})
}

//retrieve order by ID
func OrderId()http.Handler{
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handle retrieve order by id here
		// persist the database herepay
		param := r.URL.RawQuery
		fmt.Println(param)
		
	
		response := &Orders{}
		err := json.NewEncoder(w).Encode(response)
		if err != nil{
         http.Error(w,"internal server err",http.StatusInternalServerError)	
        return 		
		}
		
	})
}



