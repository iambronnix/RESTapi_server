package order

import (
	"database/sql"
	"net/http"
	"time"
)
type size struct{
    Small bool `json:"omitempty"`
	Medium bool 
	Big bool	
}

type orderItem struct{
	product string
	Size size
	Quantity int
}

type createOrder struct{
	id string
	created time.Time 
	Orders []orderItem
}
type Orders struct{
	Order []createOrder
}

type Handlers struct{
	db *sql.DB
}

//create an order
func Order()http.Handler{
	_ = &Handlers{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handle post orders logic here
	})
}
//cancel an order 
func CancelOrder()http.Handler{
	_ = &Handlers{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handle cancel logic here
	})
}

