package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"
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

func TestSimpleHTTPServer(t *testing.T) {
	srv := &http.Server{
		Addr: "127.0.0.1:8443",
		Handler: http.TimeoutHandler(DefaultHandler(),2 *time.Minute,""),
		IdleTimeout: 5 * time.Minute,
		ReadHeaderTimeout: time.Minute,
	}

	l, err := net.Listen("tcp", srv.Addr)
	if err != nil{
		t.Fatal(err)
	}

	go func() {
		err := srv.Serve(l)
		if err != http.ErrServerClosed{
			t.Error(err)
		}
	}()

	order := &OrderItem{
		product: "coffee",
		Size: size{
			small: true,
			medium: false,
			big: false,
		},
	}
	buf := new(bytes.Buffer)
	err = json.NewEncoder(buf).Encode(order)
	if err != nil{
		t.Error(err)
	}

	
   testCases := []struct{
     method string
     body io.Reader
     code  int
     response string
   }{

   {http.MethodGet,nil,http.StatusFound,"This is orders endpoint"},
   {http.MethodPost,nil,http.StatusFound,"This is order by ID endpoint"},
   {http.MethodPut,nil,http.StatusFound,"This is payment endpoint"},
   
   }
   
fmt.Print(buf)
   client := new(http.Client)
  // path1 := fmt.Sprintf("http://%s/orders", srv.Addr)
  // path2 := fmt.Sprintf("http://%s/orders/orderId",srv.Addr)
  // path3 := fmt.Sprintf("http://%s/payment",srv.Addr)
  // path4 := fmt.Sprintf("http://%s/user",srv.Addr)
  // path := []string{path1,path2,path3,path4}
  path := fmt.Sprintf("http://%s", srv.Addr)
   
   for i, c := range testCases{
    // testPath := path[i]    
   r, err := http.NewRequest(c.method,path,c.body)
   if err != nil{
   t.Errorf("%d: %v",i, err)
   continue
   }

   resp, err := client.Do(r)
   if err != nil{
   t.Errorf("%d: %v",i, err)
   continue
   }

   if resp.StatusCode !=  c.code{
    t.Errorf("%d unexpected status code: %q", i, resp.Status)
   }
  
  data , err := io.ReadAll(resp.Body)
     if err != nil{
     t.Errorf("%d: %v",i, err)
     continue
     }
     _ = resp.Body.Close()

     if c.response != string(data){
     t.Errorf("%d: expected %q; actual %q",i, c.response,data)
     }
   }
   
	
}