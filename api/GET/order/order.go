package order

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	// "github.com/google/uuid"
)
type size int
const(
	Small size = iota
	Medium
	Big
)
type OrderItem struct{
	Product string
	Size size
}
type status struct{

	Created  ,
	Progress,
	Cancelled,
	Dispatched,
	Delivered string

}
type Order struct{
	Id   string
	Created time.Time
	Status status
    Order OrderItem	
}

type Orders struct{
	Orders []Order
	
}




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
		// persist the database here
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



