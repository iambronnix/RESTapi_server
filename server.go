package main

import (
//	dltOrder "RESTapi/api/DELETE/order"
//	dltUser "RESTapi/api/DELETE/user"
	gtOrder "RESTapi/api/GET/order"
	gtPayment "RESTapi/api/GET/payment"
//	gtUser "RESTapi/api/GET/user"
//	pstOrder "RESTapi/api/POST/order"
//	pstPayment "RESTapi/api/POST/payment"
//	pstUser "RESTapi/api/POST/user"
//	aptOrder "RESTapi/api/PUT/order"
//	aptPaymemt "RESTapi/api/PUT/payment"
//	aptUser "RESTapi/api/PUT/user"
//	mdl "RESTapi/api/middleware"
	"io"
	"sort"
	"strings"

	"log"
	"net"
	"net/http"
	"time"
)

type Methods map[string]http.Handler

func (h Methods)ServeHTTP(w http.ResponseWriter, r *http.Request){
	defer func(r io.ReadCloser){
		_,_ = io.Copy(io.Discard,r)
		_ = r.Close()
	}(r.Body)

	handler, ok := h[r.Method]
	if ok {
		if handler == nil{
			http.Error(w, "null request", http.StatusBadRequest)
		}
		//parse the handler
		handler.ServeHTTP(w,r)
	}
}

func (h Methods)allowedMethods()string{
	a := make([]string,0,len(h))

	for i := range h{
		a = append(a, i)
	}

	sort.Strings(a)
	return strings.Join(a,",")
}

func DefaultHandler()http.Handler{
	return Methods{
		http.MethodGet: gtOrder.Orders(),
		http.MethodPost: gtOrder.OrderId(),
		http.MethodPut: gtPayment.Payments(),
	}
}


func main(){
	srv := &http.Server{
		Addr: "localhost:8443",
		Handler: DefaultHandler(),
		IdleTimeout: 5 *time.Minute,//you can read the man
		ReadHeaderTimeout: time.Minute,
	}

   l, err := net.Listen("tcp",srv.Addr)
   if err != nil{
    log.Fatal(err)
   }
   
   for {
   //establishing a secure connection through tls support
   err := srv.Serve(l)
   if err != nil{
   log.Println(err)
   return
      }
   }
}