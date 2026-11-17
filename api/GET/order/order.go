package order

import (
	"database/sql"
	"encoding/json"
	"net/http"
)
type Size int
const(
	Small Size = iota
	Medium
	Big
)
type OrderItem struct{
	Product string
	Size Size
}
type Status int
const (
	Created  Status = iota
	Progress
	Cancelled
	Dispatched
	Delivered
)
type Order struct{
	Id   string
	Created string
	Status Status
    Order OrderItem	
}

type Orders struct{
	Orders []Order
	
}

type Handlers struct{
    *sql.DB
}

//retrive general lists of orders
func ListOrders()http.Handler{
	_ = &Handlers{}	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handle code here
		// persist the database here
		
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

		response := &Order{}
		err := json.NewEncoder(w).Encode(response)
		if err != nil{
         http.Error(w,"internal server err",http.StatusInternalServerError)	
        return 		
		}
		
	})
}