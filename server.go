package main

import (
	mdl "RESTapi/api/middleware"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"sort"
	"strings"
	"time"
)

type Methods map[string]http.Handler
var  Counter map[string]int
var id = 0

//since this part is can't be handled concurrently according to amdahl's law 
//then it limits the whole program
func (h Methods)ServeHTTP(w http.ResponseWriter, r *http.Request){
	defer func(r io.ReadCloser){
			_,_ = io.Copy(io.Discard,r)
		_ = r.Close()
	}(r.Body)
	

	
	handler, ok := h[r.Method]
	if ok {
		if handler == nil{
			http.Error(w, "Internal Server Error", http.StatusBadRequest)
		}else{
		//parse the handler
		
		id += 1
		fmt.Println("request:",id)
		handler.ServeHTTP(w,r)
		}
		return
	}

	w.Header().Add("Allow",h.allowedMethods())
	if r.Method != http.MethodOptions{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
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
		http.MethodGet:mdl.Authenticate(true,mdl.Authorised("admin",mdl.GetUrlInspector())),
		http.MethodPost:mdl.Authenticate(true,mdl.Authorised("admin",mdl.PostUrlInspector())),
	    http.MethodPut:mdl.Authenticate(true,mdl.Authorised("admin",mdl.PutUrlInspector())),
	    http.MethodDelete:mdl.Authenticate(true,mdl.Authorised("admin",mdl.DeleteUrlInspector())),					
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