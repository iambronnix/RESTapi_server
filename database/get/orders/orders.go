package orders

import (
	"bytes"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type size struct{
	Small,
	Medium, 
	Big string
}
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
	Id   uuid.UUID
	Created time.Time
	Status status
    Order OrderItem	
}

type Orders struct{
	Orders []Order
	
}

type Dependencies struct{
	*sql.DB
}

func (d *Dependencies) Orders(param string)(*bytes.Buffer,error){
	db := d
	query, err := db.Prepare("SELECT * FROM orders WHERE status = $1")
	if err != nil{
		return nil,err		
	}

	query.Exec(param)
}
