package middleware

import (
	gtOrder "RESTapi/api/GET/order"
	gtPayment "RESTapi/api/GET/payment"
	gtUser "RESTapi/api/GET/user"
	dltOrder "RESTapi/api/DELETE/order"
	dltUser "RESTapi/api/DELETE/user"
	pstOrder "RESTapi/api/POST/order"
	pstPayment "RESTapi/api/POST/payment"
	pstUser "RESTapi/api/POST/user"
	aptOrder "RESTapi/api/PUT/order"
	aptPaymemt "RESTapi/api/PUT/payment"
	aptUser "RESTapi/api/PUT/user"
	"log"
	"net/http"
	"path"
	"strings"
)

type Incoming struct{
	*log.Logger
}
type OutGoing struct{
	*log.Logger
}


func (m Incoming) Write(p []byte)(int, error){
	return len(p),m.Output(2,string(p))
}

func (m OutGoing) Write(p []byte)(int, error){
	return len(p),m.Output(2,string(p))
}

//monitors requests
//func monitor(next http.HandlerFunc)http.Handler{
	//return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//incoming := &Incoming{Logger: log.New(os.Stdout,"monitor:",0)}
	//outgoing := &Incoming{Logger: log.New(os.Stdout,"monitor:",0)}

	//done := make(chan struct{})

	//go func(w http.ResponseWriter, r *http.Request) {
		//defer close(done)
		//write client side requests
		//b := make([]byte,uint64(1<<30))
		//j := io.TeeReader(r,incoming)

		//n, err := j.Read(b)
		//if err != nil && err != io.EOF{
			//incoming.Println(err)
			//return
		//}
		//logs server side response 
	//	w := io.MultiWriter(w,outgoing)
		//_, err = w.Write(b[:n])
		//if err != nil && err != io.EOF{
			//OutGoing.Println(err)
			//return
			
		//}
		
//	}(w,r)

	//<-done
	//next.ServeHTTP(w,r)
	//})
//}
//checks if the User has Authenticate with the system
func Authenticate(next http.HandlerFunc)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//simulate authentication
		// 
		IsAuthenticated := true

		if !IsAuthenticated{
			http.Error(w, "UnAuthorized",http.StatusUnauthorized)
			return 
		}
		next.ServeHTTP(w,r)
	})
}

//checks if the title of the user and logs into the system
func Authorised(role string, next http.HandlerFunc)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//simulate authorization 
		Authorised := "admin"

		if role!=Authorised{
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w,r)
	})
}

//examines the url for any elements that start with a given prefix
func RestrictPrefix(prefix string,next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, p := range strings.Split(path.Clean(r.URL.Path),"/"){
			if strings.HasPrefix(p,prefix){
				http.Error(w,"Not Found",http.StatusBadRequest)
			}
			
		}
		next.ServeHTTP(w,r)
	})
}

//inspects all GET requests url for routing to the right handler
func GetUrlInspector()http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlPath := r.URL.Path
		//switch acts as a secondary multiplexer 
		switch urlPath{
			case "/order":
			gtOrder.Orders().ServeHTTP(w,r)
			case "/order/orderId":
			gtOrder.OrderId().ServeHTTP(w,r)
			case "/payment":
			gtPayment.Payments().ServeHTTP(w,r)
			case "/user":
			gtUser.Users().ServeHTTP(w,r)
			default:
			w.WriteHeader(http.StatusBadRequest)
			
		}
	})
}

//inspects all POST requests url for routing to the right handler
func PostUrlInspector()http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlPath := r.URL.Path
		switch urlPath{
			case "/order":
			pstOrder.Order().ServeHTTP(w,r)
			case "/order/cancel":
			pstOrder.CancelOrder().ServeHTTP(w,r)
			case "/payment":
			pstPayment.Payments().ServeHTTP(w,r)
			case "/user":
			pstUser.User().ServeHTTP(w,r)
			default :
			w.WriteHeader(http.StatusBadRequest)
			
		}
	})
}

//inspects all PUT requests url for routing to the right handler
func PutUrlInspector()http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlPath := r.URL.Path
		switch urlPath{
			case "/order":
			aptOrder.Order().ServeHTTP(w,r)
			case "/payment":
			aptPaymemt.Payment().ServeHTTP(w,r)
			case "/user":
			aptUser.User().ServeHTTP(w,r)
			default:
			w.WriteHeader(http.StatusBadRequest)
		}
	})
}

//inspects all DELETE requests url for routing to the right handler
func DeleteUrlInspector()http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlPath := r.URL.Path
		switch urlPath{
			case "/order":
			dltOrder.Order().ServeHTTP(w,r)
			case "/user":
			dltUser.User().ServeHTTP(w,r)
			default:
			w.WriteHeader(http.StatusBadRequest)
		}
	})
}