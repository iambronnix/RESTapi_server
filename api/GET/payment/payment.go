package payment

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

//should provide {status} inthe Url
type status struct{
	paid,
	pending,
	cancelled string
}

type Order struct{
	Id string
	Created string
	Status status
}

type Handlers struct{
	db *sql.DB
}

//retrive a payment made
func Payments()http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handle payments here
		// persist the database here
		param := r.URL.RawQuery
		fmt.Println(param)

		response := &Order{}
		err := json.NewEncoder(w).Encode(response)
		if err != nil{
			http.Error(w,"internal server error",http.StatusInternalServerError)
			return 
		}
	})
}