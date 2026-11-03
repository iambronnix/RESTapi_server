package middleware

import (
	"io"
	"log"
	"net/http"
	"os"
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
func Monitor(next http.HandlerFunc)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	incoming := &Incoming{Logger: log.New(os.Stdout,"monitor:",0)}
	outgoing := &Incoming{Logger: log.New(os.Stdout,"monitor:",0)}

	done := make(chan struct{})

	go func(w http.ResponseWriter, r *http.Request) {
		defer close(done)
		//write client side requests
		b := make([]byte,uint64(1<<30))
		j := io.TeeReader(r,incoming)

		n, err := j.Read(b)
		if err != nil && err != io.EOF{
			incoming.Println(err)
			return
		}
		//logs server side response 
		w := io.MultiWriter(w,outgoing)
		_, err = w.Write(b[:n])
		if err != nil && err != io.EOF{
			OutGoing.Println(err)
			return
			
		}
		
	}(w,r)

	<-done
	next.ServeHTTP(w,r)
	})
}
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
