package main

import (
   
	gtOrder "RESTapi/api/GET/order"
	gtPayment "RESTapi/api/GET/payment"
	gtUser "RESTapi/api/GET/user"
	pstOrder "RESTapi/api/POST/order"
	pstPayment "RESTapi/api/POST/payment"
	pstUser "RESTapi/api/POST/user"
	aptOrder "RESTapi/api/PUT/order"
	aptPaymemt "RESTapi/api/PUT/payment"
	aptUser "RESTapi/api/PUT/user"
	dltOrder "RESTapi/api/DELETE/order"
	dltUser "RESTapi/api/DELETE/user"	
	mdl "RESTapi/api/middleware"
	"io"
	"log"
	"net"
	"net/http"
	"sort"
	"strings"
	"time"
)

//maps available methods to their handlers
// acts as a multiplexer
type Methods map[string]http.Handler


func (h Methods) ServeHTTP(w http.ResponseWriter, r *http.Request){
	 defer func(r io.ReadCloser) {
			_,_ = io.Copy(io.Discard,r)
			_ = r.Close()
		}(r.Body)
        //analyse which type of method requested
		if handler, ok := h[r.Method];ok{
          if handler == nil{
             http.Error(w, "invalid method", http.StatusBadRequest)
             return          
          }
          //if it exists lets handle the request
          // first check if the user is authenticated with the system
          mdl.Authenticate(handler.ServeHTTP)
        
		}
		//if the method isn't inthe map
		w.Header().Add("allowed methods",h.allowedMethods())
		if r.Method != http.MethodOptions{
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
}

func (h Methods)allowedMethods()string{
	a := make([]string, 0, len(h))

	for k := range h {
		a = append(a, k)
	}
	sort.Strings(a)

	return strings.Join(a, ",")
	
}

func handleRequests()http.Handler{
	//extract all dependicies from here
	return Methods{
		http.MethodGet: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			//handle all GET handlers here 
			http.Handle("/user",mdl.Authorised("admin",gtUser.Users().ServeHTTP))
			http.Handle("/order",mdl.Authorised("worker",gtOrder.Orders().ServeHTTP))
			http.Handle("/order/orderId",mdl.Authorised("worker",gtOrder.OrderId().ServeHTTP))
			http.Handle("/payment,",mdl.Authorised("admin",gtPayment.Payments().ServeHTTP))
		},
		),
		http.MethodPost: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//handle all POST handlers here 
			http.Handle("/user",pstUser.User())
			http.Handle("/order",pstOrder.Order())
            http.Handle("/payments",pstPayment.Payments())			
		},
		),
		http.MethodDelete: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//handle all DELETE handlers here
			http.Handle("/user",mdl.Authorised("admin",dltUser.User().ServeHTTP))
			http.Handle("/order",mdl.Authorised("admin",dltOrder.Order().ServeHTTP))
		},
		),
		http.MethodPut: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//handle all PUT handlers here
			http.Handle("/user",aptUser.User())
			http.Handle("/order",aptOrder.Order())
			http.Handle("/payment",aptPaymemt.Payment())
		},
		),
	
	}
}

func main(){

	srv := &http.Server{
		Addr: "localhost:8443",
		Handler: handleRequests(),
		IdleTimeout: 5 *time.Minute,//you can read the man
		ReadHeaderTimeout: time.Minute,
	}

   l, err := net.Listen("tcp",srv.Addr)
   if err != nil{
    log.Fatal(err)
   }
   
   for {
   //establishing a secure connection through tls support
   err := srv.ServeTLS(l,"cert.pem","key.perm")
   if err != nil{
   log.Println(err)
   return
      }
   }
}