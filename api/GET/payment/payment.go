package payment

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Status int
const(
	Paid Status = iota
	Pending
	Cancelled
)

type Order struct{
	Id string
	Created string
	Status Status
}

type Handlers struct{
	db *sql.DB
}

//retrive a payment made
func Payments()http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handle payments here
		// persist the database here

		response := &Order{}
		err := json.NewEncoder(w).Encode(response)
		if err != nil{
			http.Error(w,"internal server error",http.StatusInternalServerError)
			return 
		}
	})
}