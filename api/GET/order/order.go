package order

import (
	"database/sql"
	"net/http"
	"time"
)
type size struct{
	small,
	medium,
	big bool
}
type OrderItem struct{
	product string
	Size size
}
type status struct{
	created ,
	progress,
	cancelled,
	dispatched,
	delivered bool
}
type CreateOrder struct{
	id    ,
	Created string
	Status status
    Order OrderItem	
}

type CreateOrders struct{
	Orders []CreateOrder
	
}

type Handlers struct{
	db *sql.DB
}

//retrive general lists of orders
func Orders()http.Handler{
	_ = &Handlers{}	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handle code here
		w.WriteHeader(http.StatusFound)
		w.Write([]byte("get/order"))
	})
}

//retrieve order by ID
func OrderId()http.Handler{
	_ = &Handlers{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handle retrieve order by id here
		time.Sleep(5*time.Second)
		w.WriteHeader(http.StatusFound)
		w.Write([]byte("get/order/orderId"))
	})
}